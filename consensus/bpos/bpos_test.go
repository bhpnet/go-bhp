package bpos

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

//block explorer:  Height: 3,883,759  total: 63,214,825.02
// current block subsidy is 2.345
//
// go test -v -run ^TestCalcTotalSupply$ github.com/ethereum/go-ethereum/consensus/bpos
func TestCalcTotalSupply(t *testing.T) {
	height := uint64(3883759)
	totalSupplyAtHeight3883759 := new(big.Int).Mul(big.NewInt(6321482502), big.NewInt(1e16))
	currentBlockSubsidy := new(big.Int).Mul(big.NewInt(2345), big.NewInt(1e15))
	upgradeHeightTotalSupply := new(big.Int).Add(totalSupplyAtHeight3883759,
		new(big.Int).Mul(big.NewInt(int64(bhpv1UpgradeToV2Height-height)), currentBlockSubsidy))

	totalSupply := upgradeHeightTotalSupply
	for i := uint64(0); ; i++ {
		blockSubsidy := calcBlockSubsidy(i)
		if blockSubsidy.Cmp(big.NewInt(0)) == 0 {
			break
		}
		totalSupply = new(big.Int).Add(totalSupply, blockSubsidy)
	}
	t.Log("Total supply is: ", totalSupply.String())
	if totalSupply.Cmp(new(big.Int).Mul(big.NewInt(1_000_536_119), big.NewInt(1e17))) != 0 {
		t.Errorf("Total supply should be 100053611900000000000000000 ,but get %s\n", totalSupply.String())
	}
}

func TestCalcBlockSubsidy(t *testing.T) {

	type test struct {
		height uint64
		want   *big.Int
	}
	tests := make([]test, 9)
	tests[0] = test{0, bhpv1LastHalfBlockSubsidy}
	tests[1] = test{1, bhpv1LastHalfBlockSubsidy}

	v1UpgradeHeighDiff := bhpv1UpgradeToV2Height - bhpv1LastHalfHeight
	tests[2] = test{
		subsidyReductionInterval - 1 - v1UpgradeHeighDiff,
		bhpv1LastHalfBlockSubsidy}
	tests[3] = test{
		subsidyReductionInterval - v1UpgradeHeighDiff,
		new(big.Int).Quo(bhpv1LastHalfBlockSubsidy, big.NewInt(int64(math.Pow(2, 1))))}
	tests[4] = test{
		subsidyReductionInterval + 1 - v1UpgradeHeighDiff,
		new(big.Int).Quo(bhpv1LastHalfBlockSubsidy, big.NewInt(int64(math.Pow(2, 1))))}

	tests[5] = test{
		subsidyReductionInterval*2 - 1 - v1UpgradeHeighDiff,
		new(big.Int).Quo(bhpv1LastHalfBlockSubsidy, big.NewInt(int64(math.Pow(2, 1))))}
	tests[6] = test{
		subsidyReductionInterval*2 - v1UpgradeHeighDiff,
		new(big.Int).Quo(bhpv1LastHalfBlockSubsidy, big.NewInt(int64(math.Pow(2, 2))))}
	tests[7] = test{
		subsidyReductionInterval*2 + 1 - v1UpgradeHeighDiff,
		new(big.Int).Quo(bhpv1LastHalfBlockSubsidy, big.NewInt(int64(math.Pow(2, 2))))}

	tests[8] = test{
		18446744073709551615 - v1UpgradeHeighDiff, //uint64 max and prevent overflow
		big.NewInt(0)}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.height), func(t *testing.T) {
			if got := calcBlockSubsidy(tt.height); got.Cmp(tt.want) != 0 {
				t.Errorf("hieght= %v, CalcBlockReward() = %v, want %v", tt.height, got, tt.want)
			}
		})
	}
}
