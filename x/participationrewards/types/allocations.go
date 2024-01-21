package types

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
)

type RewardsAllocation struct {
	ValidatorSelection sdkmath.Int
	Holdings           sdkmath.Int
	Lockup             sdkmath.Int
}

// GetRewardsAllocations returns an instance of rewardsAllocation with values
// set according to the given moduleBalance and distribution proportions.
func GetRewardsAllocations(moduleBalance sdkmath.Int, proportions DistributionProportions) (*RewardsAllocation, error) {
	if moduleBalance.IsNil() || moduleBalance.IsZero() {
		return nil, ErrNothingToAllocate
	}

	if sum := proportions.Total(); !sum.Equal(sdkmath.LegacyOneDec()) {
		return nil, fmt.Errorf("%w: got %v", ErrInvalidTotalProportions, sum)
	}

	var allocation RewardsAllocation

	// split participation rewards allocations
	allocation.ValidatorSelection = sdkmath.LegacyNewDecFromInt(moduleBalance).Mul(proportions.ValidatorSelectionAllocation).TruncateInt()
	allocation.Holdings = sdkmath.LegacyNewDecFromInt(moduleBalance).Mul(proportions.HoldingsAllocation).TruncateInt()
	allocation.Lockup = sdkmath.LegacyNewDecFromInt(moduleBalance).Mul(proportions.LockupAllocation).TruncateInt()

	// use sum to check total distribution to collect and allocate dust
	sum := allocation.Lockup.Add(allocation.ValidatorSelection).Add(allocation.Holdings)
	dust := moduleBalance.Sub(sum)

	// Add dust to validator choice allocation (favors decentralization)
	allocation.ValidatorSelection = allocation.ValidatorSelection.Add(dust)

	return &allocation, nil
}
