package types_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/quicksilver-zone/quicksilver/v7/x/mint/types"
)

// Benchmarking :)
// previously using sdkmath.Int operations:
// BenchmarkEpochProvision-4 5000000 220 ns/op
//
// using sdkmath.LegacyDec operations: (current implementation)
// BenchmarkEpochProvision-4 3000000 429 ns/op.
func BenchmarkEpochProvision(b *testing.B) {
	b.ReportAllocs()
	minter := types.InitialMinter()
	params := types.DefaultParams()

	s1 := rand.NewSource(100)
	r1 := rand.New(s1)
	minter.EpochProvisions = sdkmath.LegacyNewDec(r1.Int63n(1000000))

	// run the EpochProvision function b.N times
	for n := 0; n < b.N; n++ {
		minter.EpochProvision(params)
	}
}

// Next epoch provisions benchmarking
// BenchmarkNextEpochProvisions-4 5000000 251 ns/op.
func BenchmarkNextEpochProvisions(b *testing.B) {
	b.ReportAllocs()
	minter := types.InitialMinter()
	params := types.DefaultParams()

	// run the NextEpochProvisions function b.N times
	for n := 0; n < b.N; n++ {
		minter.NextEpochProvisions(params)
	}
}

func TestEpochProvision(t *testing.T) {
	params := types.DefaultParams()
	minter := types.DefaultInitialMinter()
	minter.EpochProvisions = sdkmath.LegacyNewDecWithPrec(75, 2)

	actual := minter.NextEpochProvisions(params)
	require.Equal(t, sdkmath.LegacyNewDecWithPrec(5625, 4), actual)
	require.Equal(t, sdk.NewCoin(params.MintDenom, actual.TruncateInt()), minter.EpochProvision(params))
}

func TestMinterValidate(t *testing.T) {
	testcases := []struct {
		name    string
		minter  types.Minter
		isValid bool
	}{
		{
			"valid",
			types.InitialMinter(),
			true,
		},
		{
			"negative",
			types.Minter{
				EpochProvisions: sdkmath.LegacyNewDec(-1),
			},
			false,
		},
		{
			"nil",
			types.Minter{},
			false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.minter.Validate()
			if !tc.isValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
