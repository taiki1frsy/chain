package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/UnUniFi/chain/x/derivatives/types"
)

func RandomGenesisBool(r *rand.Rand) bool {
	// 90% chance
	return r.Int63n(100) < 90
}

func RandomizedGenState(simState *module.SimulationState) {
	sdk.NewCoins()
	// numAccs := int64(len(simState.Accounts))

	derivativesGenesis := types.GenesisState{
		Params: types.Params{
			PoolParams: types.PoolParams{
				QuoteTicker:                 "usd",
				BaseLptMintFee:              sdk.NewDecWithPrec(1, 3),
				BaseLptRedeemFee:            sdk.NewDecWithPrec(1, 3),
				BorrowingFeeRatePerHour:     sdk.NewDecWithPrec(1, 6),
				ReportLiquidationRewardRate: sdk.NewDecWithPrec(3, 1),
				ReportLevyPeriodRewardRate:  sdk.NewDecWithPrec(3, 1),
				AcceptedAssetsConf: []types.PoolAssetConf{
					{
						Denom:        "btc",
						TargetWeight: sdk.NewDecWithPrec(1, 2),
					},
					{
						Denom:        "eth",
						TargetWeight: sdk.NewDecWithPrec(1, 3),
					},
				},
			},
			PerpetualFutures: types.PerpetualFuturesParams{
				CommissionRate:        sdk.NewDecWithPrec(1, 3),
				MarginMaintenanceRate: sdk.NewDecWithPrec(5, 1),
				ImaginaryFundingRateProportionalCoefficient: sdk.NewDecWithPrec(5, 4),
				Markets: []*types.Market{
					{
						BaseDenom:  "ubtc",
						QuoteDenom: "uusdc",
					},
					{
						BaseDenom:  "ueth",
						QuoteDenom: "uusdc",
					},
				},
				MaxLeverage: 30,
			},
		},
	}

	paramsBytes, err := json.MarshalIndent(&derivativesGenesis.Params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated derivatives parameters:\n%s\n", paramsBytes)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&derivativesGenesis)
}
