package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc/types"
)

func (k msgServer) ChangeValidatorWeight(goCtx context.Context, msg *types.MsgChangeValidatorWeight) (*types.MsgChangeValidatorWeightResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hostZone, found := k.GetHostZone(ctx, msg.HostZone)
	if !found {
		k.Logger(ctx).Error(fmt.Sprintf("Host Zone %s not found", msg.HostZone))
		return nil, types.ErrInvalidHostZone
	}
	validators := hostZone.Validators
	for _, validator := range validators {
		if validator.GetAddress() == msg.ValAddr {
			validator.Weight = msg.Weight
			k.SetHostZone(ctx, hostZone)
			return &types.MsgChangeValidatorWeightResponse{}, nil
		}
	}

	k.Logger(ctx).Error(fmt.Sprintf("Validator %s not found on Host Zone %s", msg.ValAddr, msg.HostZone))
	return nil, types.ErrValidatorNotFound
}
