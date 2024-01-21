package participationrewards_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	simapp "github.com/quicksilver-zone/quicksilver/v7/app"
	"github.com/quicksilver-zone/quicksilver/v7/x/participationrewards"
	"github.com/quicksilver-zone/quicksilver/v7/x/participationrewards/types"
)

func TestParticipationRewardsExportGenesis(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false)

	chainStartTime := ctx.BlockTime()

	pool := types.OsmosisPoolProtocolData{
		PoolID:      1,
		PoolName:    "atom/osmo",
		LastUpdated: chainStartTime,
	}

	bz, err := json.Marshal(pool)
	if err != nil {
		t.Fatalf("unable to marshal protocol data: %v", err)
	}
	protocolData := types.NewProtocolData(types.ProtocolDataType_name[int32(types.ProtocolDataTypeOsmosisPool)], bz)

	app.ParticipationRewardsKeeper.SetProtocolData(ctx, []byte(fmt.Sprintf("%d", pool.PoolID)), protocolData)

	genesis := participationrewards.ExportGenesis(ctx, app.ParticipationRewardsKeeper)

	// 0,0,0,4 (binary encoded types.ProtocolDataTypeOsmosisPool)
	// 49 (ASCII value of '1')
	require.Equal(t, string([]byte{0, 0, 0, 0, 0, 0, 0, 4, 49}), genesis.ProtocolData[0].Key)
	require.Equal(t, types.ProtocolDataType_name[int32(types.ProtocolDataTypeOsmosisPool)], genesis.ProtocolData[0].ProtocolData.Type)
}

func TestParticipationRewardsInitGenesis(t *testing.T) {
	// setup params
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false)

	now := time.Now()
	ctx = ctx.WithBlockHeight(1)
	ctx = ctx.WithBlockTime(now)

	validOsmosisData := `{
	"poolid": 1,
	"poolname": "atom/osmo",
	"pooltype": "balancer",
	"denoms": {
		"ibc/00000000000000000000000000000000": {"denom": "ustake", "chainid": "testzone-1"}
	}
}`

	kpd := &types.KeyedProtocolData{
		Key: "6",
		ProtocolData: &types.ProtocolData{
			Type: types.ProtocolDataType_name[int32(types.ProtocolDataTypeOsmosisPool)],
			Data: []byte(validOsmosisData),
		},
	}

	// test genesisState validation
	genesisState := types.GenesisState{
		Params: types.Params{
			DistributionProportions: types.DistributionProportions{
				ValidatorSelectionAllocation: sdkmath.LegacyNewDecWithPrec(5, 1),
				HoldingsAllocation:           sdkmath.LegacyNewDecWithPrec(5, 1),
				LockupAllocation:             sdkmath.LegacyZeroDec(),
			},
		},
		ProtocolData: []*types.KeyedProtocolData{kpd},
	}
	require.NoError(t, genesisState.Validate(), "genesis validation failed")

	participationrewards.InitGenesis(ctx, app.ParticipationRewardsKeeper, genesisState)

	require.Equal(t, app.ParticipationRewardsKeeper.GetParams(ctx).DistributionProportions.ValidatorSelectionAllocation, sdkmath.LegacyNewDecWithPrec(5, 1))
	require.Equal(t, app.ParticipationRewardsKeeper.GetParams(ctx).DistributionProportions.HoldingsAllocation, sdkmath.LegacyNewDecWithPrec(5, 1))
	require.Equal(t, app.ParticipationRewardsKeeper.GetParams(ctx).DistributionProportions.LockupAllocation, sdkmath.LegacyZeroDec())

	pd, found := app.ParticipationRewardsKeeper.GetProtocolData(ctx, types.ProtocolDataTypeOsmosisPool, "6")
	require.True(t, found)
	require.Equal(t, types.ProtocolDataType_name[int32(types.ProtocolDataTypeOsmosisPool)], pd.Type)
}
