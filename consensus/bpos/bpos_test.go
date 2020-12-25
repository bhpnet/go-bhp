package bpos

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

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

func TestCalcFundStakingSubsidy(t *testing.T) {
	blockSubsidy := big.NewInt(0)
	fundSubsidy, stakingSubsidy := calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(0)) != 0 || stakingSubsidy.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy = big.NewInt(1)
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(0)) != 0 || stakingSubsidy.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy = big.NewInt(2)
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(1)) != 0 || stakingSubsidy.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy = big.NewInt(3)
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(2)) != 0 || stakingSubsidy.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy = big.NewInt(9)
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(6)) != 0 || stakingSubsidy.Cmp(big.NewInt(3)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy = big.NewInt(10)
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(7)) != 0 || stakingSubsidy.Cmp(big.NewInt(3)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy = big.NewInt(11)
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(7)) != 0 || stakingSubsidy.Cmp(big.NewInt(4)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy = big.NewInt(100)
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(70)) != 0 || stakingSubsidy.Cmp(big.NewInt(30)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy, _ = new(big.Int).SetString("1000000000000000000", 10)
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(7e17)) != 0 || stakingSubsidy.Cmp(big.NewInt(3e17)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
	blockSubsidy = bhpv1LastHalfBlockSubsidy
	fundSubsidy, stakingSubsidy = calcFundStakingSubsidy(blockSubsidy)
	if fundSubsidy.Cmp(big.NewInt(1.6415e18)) != 0 || stakingSubsidy.Cmp(big.NewInt(7.035e17)) != 0 {
		t.Errorf("block subsidy = %v , fund subsidy = %v, stakign subsidy = %v", blockSubsidy.String(), fundSubsidy.String(), stakingSubsidy.String())
	}
}
