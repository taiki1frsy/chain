package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	nftmarkettypes "github.com/UnUniFi/chain/x/nftbackedloan/types"
)

var (
	DafaultRewardParams = []*RewardParams{
		{
			ModuleName: nftmarkettypes.ModuleName,
			RewardRate: []RewardRate{
				{
					RewardType: RewardType_FRONTEND_DEVELOPERS,
					Rate:       sdk.MustNewDecFromStr("0.5"),
				},
			},
		},
		{
			ModuleName: nftmarkettypes.ModuleName,
			RewardRate: []RewardRate{
				{
					RewardType: RewardType_STAKERS,
					Rate:       sdk.MustNewDecFromStr("0.5"),
				},
			},
		},
	}
	KeyRewardParams = []byte("RewardParams")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return DefaultParams()
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		RewardParams: DafaultRewardParams,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyRewardParams, &p.RewardParams, validateRewardParams),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {

	if err := validateRewardParams(p.RewardParams); err != nil {
		return err
	}

	return nil
}

func validateRewardParams(i interface{}) error {
	rewardParams, ok := i.([]*RewardParams)
	if !ok {
		return fmt.Errorf("invalid paramter type: %T", i)
	}

	for _, rewardParam := range rewardParams {
		for _, rate := range rewardParam.RewardRate {
			if rate.Rate.GT(sdk.OneDec()) {
				return fmt.Errorf("each reward rate must be less than 1 dec")
			}

			if rate.Rate.IsNegative() {
				return fmt.Errorf("each reward rate must be positive")
			}
		}
	}

	return nil
}
