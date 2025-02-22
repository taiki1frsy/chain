package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"

	"github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc/types"
)

func (k msgServer) RestoreInterchainAccount(goCtx context.Context, msg *types.MsgRestoreInterchainAccount) (*types.MsgRestoreInterchainAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	hostZone, found := k.GetHostZone(ctx, msg.ChainId)
	if !found {
		k.Logger(ctx).Error(fmt.Sprintf("Host Zone not found: %s", msg.ChainId))
		return nil, types.ErrInvalidHostZone
	}

	// Get ConnectionEnd (for counterparty connection)
	connectionEnd, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, hostZone.ConnectionId)
	if !found {
		errMsg := fmt.Sprintf("invalid connection id from host %s, %s not found", msg.ChainId, hostZone.ConnectionId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, errMsg)
	}
	counterpartyConnection := connectionEnd.Counterparty

	owner := types.FormatICAAccountOwner(msg.ChainId, msg.AccountType)

	// only allow restoring an account if it already exists
	portID, err := icatypes.NewControllerPortID(owner)
	if err != nil {
		errMsg := fmt.Sprintf("could not create portID for ICA controller account address: %s", owner)
		k.Logger(ctx).Error(errMsg)
		return nil, err
	}
	_, exists := k.ICAControllerKeeper.GetInterchainAccountAddress(ctx, hostZone.ConnectionId, portID)
	if !exists {
		errMsg := fmt.Sprintf("ICA controller account address not found: %s", owner)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrapf(types.ErrInvalidInterchainAccountAddress, errMsg)
	}

	appVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: hostZone.ConnectionId,
		HostConnectionId:       counterpartyConnection.ConnectionId,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))

	if err := k.RegisterInterchainAccount(ctx, hostZone.ConnectionId, owner, appVersion); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("unable to register %s account : %s", msg.AccountType.String(), err))
		return nil, err
	}

	return &types.MsgRestoreInterchainAccountResponse{}, nil
}
