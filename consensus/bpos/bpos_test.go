package bpos

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestCalcTotalSupply(t *testing.T) {
	// TODO: bhpnet finish unit test
	totalSupply := big.NewInt(0)
	for i := uint64(0); ; i++ {
		blockSubsidy := calcBlockSubsidy(i)
		if blockSubsidy.Cmp(big.NewInt(0)) == 0 ||
			totalSupply.Cmp(new(big.Int).Mul(big.NewInt(1e8), big.NewInt(1e18))) == 1 {
			break
		}
		totalSupply = new(big.Int).Add(totalSupply, blockSubsidy)
		fmt.Println(i, "-----", totalSupply.String())
	}
	totalSupplyStr := totalSupply.String()
	fmt.Println(totalSupplyStr)
}

func TestCalcBlockSubsidy(t *testing.T) {

	type test struct {
		height uint64
		want   *big.Int
	}
	tests := make([]test, 12)
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
		subsidyReductionInterval*62 - 1 - v1UpgradeHeighDiff,
		big.NewInt(1)}
	tests[9] = test{
		subsidyReductionInterval*62 - v1UpgradeHeighDiff,
		big.NewInt(0)}
	tests[10] = test{
		subsidyReductionInterval*62 + 1 - v1UpgradeHeighDiff,
		big.NewInt(0)}

	tests[11] = test{
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
