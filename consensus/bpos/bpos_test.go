package bpos

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"
)

func TestCalcBlockReward(t *testing.T) {

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
			if got := calcBlockSubsidy(tt.height); strings.Compare(got.String(), tt.want.String()) != 0 {
				t.Errorf("hieght= %v, CalcBlockReward() = %v, want %v", tt.height, got, tt.want)
			}
		})
	}
}
