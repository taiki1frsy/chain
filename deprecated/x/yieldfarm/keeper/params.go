package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/deprecated/x/yieldfarm/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	params := types.Params{}
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
