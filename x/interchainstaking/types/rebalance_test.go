package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	"github.com/quicksilver-zone/quicksilver/v7/utils/addressutils"
	"github.com/quicksilver-zone/quicksilver/v7/x/interchainstaking/types"
)

func TestDetermineAllocationsForRebalancing(t *testing.T) {
	vals := addressutils.GenerateValidatorsDeterministic(5)

	type testcase struct {
		name        string
		allocations map[string]sdkmath.Int
		target      types.ValidatorIntents
		locked      map[string]bool
		expected    types.RebalanceTargets
	}

	tcs := []testcase{
		{
			name: "100% No Existing Redelegations",
			allocations: map[string]sdkmath.Int{
				vals[0]: sdkmath.NewInt(10),
				vals[1]: sdkmath.NewInt(10),
				vals[2]: sdkmath.NewInt(10),
				vals[3]: sdkmath.NewInt(10),
				vals[4]: sdkmath.NewInt(10),
			},
			target: types.ValidatorIntents{
				&types.ValidatorIntent{
					ValoperAddress: vals[0],
					Weight:         sdkmath.LegacyNewDec(1),
				},
			},
			locked: map[string]bool{},
			expected: types.RebalanceTargets{
				{
					Source: vals[1],
					Target: vals[0],
					Amount: sdkmath.NewInt(10),
				},
				{
					Source: vals[2],
					Target: vals[0],
					Amount: sdkmath.NewInt(10),
				},
				{
					Source: vals[3],
					Target: vals[0],
					Amount: sdkmath.NewInt(5),
				},
			},
		},
		{
			name: "50/50 No Existing Redelegations, Constrained by total",
			allocations: map[string]sdkmath.Int{
				vals[0]: sdkmath.NewInt(10),
				vals[1]: sdkmath.NewInt(10),
				vals[2]: sdkmath.NewInt(10),
				vals[3]: sdkmath.NewInt(10),
				vals[4]: sdkmath.NewInt(10),
			},
			target: types.ValidatorIntents{
				&types.ValidatorIntent{
					ValoperAddress: vals[0],
					Weight:         sdkmath.LegacyNewDecWithPrec(5, 1),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[1],
					Weight:         sdkmath.LegacyNewDecWithPrec(5, 1),
				},
			},
			locked: map[string]bool{},
			expected: types.RebalanceTargets{
				{
					Source: vals[2],
					Target: vals[0],
					Amount: sdkmath.NewInt(10),
				},
				{
					Source: vals[3],
					Target: vals[0],
					Amount: sdkmath.NewInt(5),
				},
				{
					Source: vals[3],
					Target: vals[1],
					Amount: sdkmath.NewInt(5),
				},
				{
					Source: vals[4],
					Target: vals[1],
					Amount: sdkmath.NewInt(5),
				},
			},
		},
		{
			name: "50/50 No Existing Redelegations, Unconstrained",
			allocations: map[string]sdkmath.Int{
				vals[0]: sdkmath.NewInt(10),
				vals[1]: sdkmath.NewInt(10),
				vals[2]: sdkmath.NewInt(10),
				vals[3]: sdkmath.NewInt(10),
			},
			target: types.ValidatorIntents{
				&types.ValidatorIntent{
					ValoperAddress: vals[0],
					Weight:         sdkmath.LegacyNewDecWithPrec(5, 1),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[1],
					Weight:         sdkmath.LegacyNewDecWithPrec(5, 1),
				},
			},
			locked: map[string]bool{},
			expected: []*types.RebalanceTarget{
				{
					Source: vals[2],
					Target: vals[0],
					Amount: sdkmath.NewInt(10),
				},
				{
					Source: vals[3],
					Target: vals[1],
					Amount: sdkmath.NewInt(10),
				},
			},
		},
		{
			name: "Drop one validator, No Existing Redelegations",
			allocations: map[string]sdkmath.Int{
				vals[0]: sdkmath.NewInt(8),
				vals[1]: sdkmath.NewInt(8),
				vals[2]: sdkmath.NewInt(8),
				vals[3]: sdkmath.NewInt(8),
				vals[4]: sdkmath.NewInt(8),
			},
			target: types.ValidatorIntents{
				&types.ValidatorIntent{
					ValoperAddress: vals[0],
					Weight:         sdkmath.LegacyNewDecWithPrec(25, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[1],
					Weight:         sdkmath.LegacyNewDecWithPrec(25, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[2],
					Weight:         sdkmath.LegacyNewDecWithPrec(25, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[3],
					Weight:         sdkmath.LegacyNewDecWithPrec(25, 2),
				},
			},
			locked: map[string]bool{},
			expected: types.RebalanceTargets{
				{
					Source: vals[4],
					Target: vals[0],
					Amount: sdkmath.NewInt(2),
				},
				{
					Source: vals[4],
					Target: vals[1],
					Amount: sdkmath.NewInt(2),
				},
				{
					Source: vals[4],
					Target: vals[2],
					Amount: sdkmath.NewInt(2),
				},
				{
					Source: vals[4],
					Target: vals[3],
					Amount: sdkmath.NewInt(2),
				},
			},
		},
		{
			name: "Add one validator, No Existing Redelegations",
			allocations: map[string]sdkmath.Int{
				vals[0]: sdkmath.NewInt(10),
				vals[1]: sdkmath.NewInt(10),
				vals[2]: sdkmath.NewInt(10),
				vals[3]: sdkmath.NewInt(10),
			},
			target: types.ValidatorIntents{
				&types.ValidatorIntent{
					ValoperAddress: vals[0],
					Weight:         sdkmath.LegacyNewDecWithPrec(20, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[1],
					Weight:         sdkmath.LegacyNewDecWithPrec(20, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[2],
					Weight:         sdkmath.LegacyNewDecWithPrec(20, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[3],
					Weight:         sdkmath.LegacyNewDecWithPrec(20, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[4],
					Weight:         sdkmath.LegacyNewDecWithPrec(20, 2),
				},
			},
			locked: map[string]bool{},
			expected: types.RebalanceTargets{
				{
					Source: vals[0],
					Target: vals[4],
					Amount: sdkmath.NewInt(2),
				},
				{
					Source: vals[1],
					Target: vals[4],
					Amount: sdkmath.NewInt(2),
				},
				{
					Source: vals[2],
					Target: vals[4],
					Amount: sdkmath.NewInt(2),
				},
				{
					Source: vals[3],
					Target: vals[4],
					Amount: sdkmath.NewInt(2),
				},
			},
		},
		{
			name: "Attempt redelegate away from locked validator; no-op",
			allocations: map[string]sdkmath.Int{
				vals[0]: sdkmath.NewInt(10),
				vals[1]: sdkmath.NewInt(10),
				vals[2]: sdkmath.NewInt(10),
				vals[3]: sdkmath.NewInt(10),
				vals[4]: sdkmath.NewInt(10),
			},
			target: types.ValidatorIntents{
				&types.ValidatorIntent{
					ValoperAddress: vals[0],
					Weight:         sdkmath.LegacyNewDecWithPrec(10, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[1],
					Weight:         sdkmath.LegacyNewDecWithPrec(225, 3),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[2],
					Weight:         sdkmath.LegacyNewDecWithPrec(225, 3),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[3],
					Weight:         sdkmath.LegacyNewDecWithPrec(225, 3),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[4],
					Weight:         sdkmath.LegacyNewDecWithPrec(225, 3),
				},
			},
			locked: map[string]bool{
				vals[0]: true,
			},
			expected: types.RebalanceTargets{},
		},
		{
			name: "Delegate away from 2; 1 locked validator; v1 -15; v2 + 10; v3 +5",
			allocations: map[string]sdkmath.Int{
				vals[0]: sdkmath.NewInt(20),
				vals[1]: sdkmath.NewInt(20),
				vals[2]: sdkmath.NewInt(20),
				vals[3]: sdkmath.NewInt(20),
				vals[4]: sdkmath.NewInt(20),
			},
			target: types.ValidatorIntents{
				&types.ValidatorIntent{
					ValoperAddress: vals[0],
					Weight:         sdkmath.LegacyNewDecWithPrec(5, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[1],
					Weight:         sdkmath.LegacyNewDecWithPrec(5, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[2],
					Weight:         sdkmath.LegacyNewDecWithPrec(30, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[3],
					Weight:         sdkmath.LegacyNewDecWithPrec(30, 2),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[4],
					Weight:         sdkmath.LegacyNewDecWithPrec(30, 2),
				},
			},
			locked: map[string]bool{
				vals[0]: true,
			},
			expected: types.RebalanceTargets{
				{
					Source: vals[1],
					Target: vals[2],
					Amount: sdkmath.NewInt(10),
				},
				{
					Source: vals[1],
					Target: vals[3],
					Amount: sdkmath.NewInt(5),
				},
			},
		},
		{
			name: "v0 missing, v1 zero; one new vals. Should delegate v0: -50; v1: -25; v2: +25; v3: +50",
			allocations: map[string]sdkmath.Int{
				vals[0]: sdkmath.NewInt(50),
				vals[1]: sdkmath.NewInt(50),
				vals[2]: sdkmath.NewInt(50),
			},
			target: types.ValidatorIntents{
				&types.ValidatorIntent{
					ValoperAddress: vals[1],
					Weight:         sdkmath.LegacyZeroDec(),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[2],
					Weight:         sdkmath.LegacyNewDecWithPrec(5, 1),
				},
				&types.ValidatorIntent{
					ValoperAddress: vals[3],
					Weight:         sdkmath.LegacyNewDecWithPrec(5, 1),
				},
			},
			locked: map[string]bool{},
			expected: types.RebalanceTargets{
				{
					Source: vals[0],
					Target: vals[2],
					Amount: sdkmath.NewInt(25),
				},
				{
					Source: vals[0],
					Target: vals[3],
					Amount: sdkmath.NewInt(25),
				},
				{
					Source: vals[1],
					Target: vals[3],
					Amount: sdkmath.NewInt(25),
				},
			},
		},
	}

	for _, tt := range tcs {
		t.Run(tt.name, func(t *testing.T) {
			currentSum, lockedSum := func(in map[string]sdkmath.Int, locked map[string]bool) (sum, lockedSum sdkmath.Int) {
				sum = sdkmath.ZeroInt()
				lockedSum = sdkmath.ZeroInt()
				for k, v := range in {
					sum = sum.Add(v)
					if locked[k] {
						lockedSum = lockedSum.Add(v)
					}
				}
				return sum, lockedSum
			}(tt.allocations, tt.locked)

			actual := types.DetermineAllocationsForRebalancing(
				tt.allocations, tt.locked, currentSum, lockedSum, tt.target, make(map[string]sdkmath.Int), nil,
			)

			require.ElementsMatch(t, tt.expected, actual)
		})
	}
}
