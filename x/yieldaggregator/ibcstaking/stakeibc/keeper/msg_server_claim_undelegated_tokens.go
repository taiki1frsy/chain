package keeper

import (
	"context"
	"fmt"

	"github.com/spf13/cast"

	recordstypes "github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/records/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"

	epochstypes "github.com/UnUniFi/chain/x/epochs/types"
	"github.com/UnUniFi/chain/x/yieldaggregator/ibcstaking/stakeibc/types"
)

type IcaTx struct {
	ConnectionId string
	Msgs         []sdk.Msg
	Account      types.ICAAccount
	Timeout      uint64
}

// WithdrawUndelegatedTokensToChain should be called automatically to withdraw via IBC transfer rather than direct transfer
func (k Keeper) WithdrawUndelegatedTokensToChain(ctx sdk.Context, msg *types.MsgClaimUndelegatedTokens) (*types.MsgClaimUndelegatedTokensResponse, error) {
	k.Logger(ctx).Info(fmt.Sprintf("WithdrawUndelegatedTokensToChain %v", msg))
	userRedemptionRecord, err := k.GetClaimableRedemptionRecord(ctx, msg)
	if err != nil {
		errMsg := fmt.Sprintf("unable to find claimable redemption record for msg: %v, error %s", msg, err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrRecordNotFound, errMsg)
	}

	icaTx, err := k.GetRedemptionIBCTransferMsg(ctx, userRedemptionRecord, msg.HostZoneId)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "unable to build redemption transfer message")
	}

	// add callback data
	claimCallback := types.ClaimCallback{
		UserRedemptionRecordId: userRedemptionRecord.Id,
		ChainId:                msg.HostZoneId,
		EpochNumber:            msg.Epoch,
	}
	marshalledCallbackArgs, err := k.MarshalClaimCallbackArgs(ctx, claimCallback)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "unable to marshal claim callback args")
	}
	_, err = k.SubmitTxs(ctx, icaTx.ConnectionId, icaTx.Msgs, icaTx.Account, icaTx.Timeout, CLAIM, marshalledCallbackArgs)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Submit tx error: %s", err.Error()))
		return nil, sdkerrors.Wrap(err, "unable to submit ICA redemption tx")
	}

	// Set claimIsPending to true, so that the record can't be double claimed
	userRedemptionRecord.ClaimIsPending = true
	k.RecordsKeeper.SetUserRedemptionRecord(ctx, *userRedemptionRecord)

	return &types.MsgClaimUndelegatedTokensResponse{}, nil
}

func (k msgServer) ClaimUndelegatedTokens(goCtx context.Context, msg *types.MsgClaimUndelegatedTokens) (*types.MsgClaimUndelegatedTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.Logger(ctx).Info(fmt.Sprintf("ClaimUndelegatedTokens %v", msg))
	userRedemptionRecord, err := k.GetClaimableRedemptionRecord(ctx, msg)
	if err != nil {
		errMsg := fmt.Sprintf("unable to find claimable redemption record for msg: %v, error %s", msg, err.Error())
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrRecordNotFound, errMsg)
	}

	icaTx, err := k.GetRedemptionTransferMsg(ctx, userRedemptionRecord, msg.HostZoneId)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "unable to build redemption transfer message")
	}

	// add callback data
	claimCallback := types.ClaimCallback{
		UserRedemptionRecordId: userRedemptionRecord.Id,
		ChainId:                msg.HostZoneId,
		EpochNumber:            msg.Epoch,
	}
	marshalledCallbackArgs, err := k.MarshalClaimCallbackArgs(ctx, claimCallback)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "unable to marshal claim callback args")
	}
	_, err = k.SubmitTxs(ctx, icaTx.ConnectionId, icaTx.Msgs, icaTx.Account, icaTx.Timeout, CLAIM, marshalledCallbackArgs)
	if err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("Submit tx error: %s", err.Error()))
		return nil, sdkerrors.Wrap(err, "unable to submit ICA redemption tx")
	}

	// Set claimIsPending to true, so that the record can't be double claimed
	userRedemptionRecord.ClaimIsPending = true
	k.RecordsKeeper.SetUserRedemptionRecord(ctx, *userRedemptionRecord)

	return &types.MsgClaimUndelegatedTokensResponse{}, nil
}

func (k Keeper) GetClaimableRedemptionRecord(ctx sdk.Context, msg *types.MsgClaimUndelegatedTokens) (*recordstypes.UserRedemptionRecord, error) {
	// grab the UserRedemptionRecord from the store
	userRedemptionRecordKey := recordstypes.UserRedemptionRecordKeyFormatter(msg.HostZoneId, msg.Epoch, msg.Sender)
	userRedemptionRecord, found := k.RecordsKeeper.GetUserRedemptionRecord(ctx, userRedemptionRecordKey)
	if !found {
		errMsg := fmt.Sprintf("User redemption record %s not found on host zone %s", userRedemptionRecordKey, msg.HostZoneId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrInvalidUserRedemptionRecord, errMsg)
	}

	// check that the record is claimable
	hostZoneUnbonding, found := k.RecordsKeeper.GetHostZoneUnbondingByChainId(ctx, userRedemptionRecord.EpochNumber, msg.HostZoneId)
	if !found {
		errMsg := fmt.Sprintf("Host zone unbonding record %s not found on host zone %s", userRedemptionRecordKey, msg.HostZoneId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrapf(types.ErrInvalidUserRedemptionRecord, errMsg)
	}
	// records associated with host zone unbondings are claimable after the host zone unbonding tokens have been transferred to the redemption account
	if hostZoneUnbonding.Status != recordstypes.HostZoneUnbonding_TRANSFERRED {
		errMsg := fmt.Sprintf("User redemption record %s is not claimable, host zone unbonding has status: %s, requires status TRANSFERRED", userRedemptionRecord.Id, hostZoneUnbonding.Status)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrapf(types.ErrInvalidUserRedemptionRecord, errMsg)
	}
	// records that have claimIsPending set to True have already been claimed (and are pending an ack)
	if userRedemptionRecord.ClaimIsPending {
		errMsg := fmt.Sprintf("User redemption record %s is not claimable, pending ack", userRedemptionRecord.Id)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrapf(types.ErrInvalidUserRedemptionRecord, errMsg)
	}
	return &userRedemptionRecord, nil
}

func (k Keeper) GetRedemptionTransferMsg(ctx sdk.Context, userRedemptionRecord *recordstypes.UserRedemptionRecord, hostZoneId string) (*IcaTx, error) {
	// grab necessary fields to construct ICA call
	hostZone, found := k.GetHostZone(ctx, hostZoneId)
	if !found {
		errMsg := fmt.Sprintf("Host zone %s not found", hostZoneId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrInvalidHostZone, errMsg)
	}
	redemptionAccount, found := k.GetRedemptionAccount(ctx, hostZone)
	if !found {
		errMsg := fmt.Sprintf("Redemption account not found for host zone %s", hostZoneId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrInvalidHostZone, errMsg)
	}

	var msgs []sdk.Msg
	rrAmt, err := cast.ToInt64E(userRedemptionRecord.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidUserRedemptionRecord, err.Error())
	}
	msgs = append(msgs, &bankTypes.MsgSend{
		FromAddress: redemptionAccount.Address,
		ToAddress:   "cosmos1mjk79fjjgpplak5wq838w0yd982gzkyfrk07am", // userRedemptionRecord.Receiver,
		Amount:      sdk.NewCoins(sdk.NewInt64Coin(userRedemptionRecord.Denom, rrAmt)),
	})

	// Give claims a 10 minute timeout
	epochTracker, found := k.GetEpochTracker(ctx, epochstypes.BASE_EPOCH)
	if !found {
		errMsg := fmt.Sprintf("Epoch tracker not found for epoch %s", epochstypes.BASE_EPOCH)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrEpochNotFound, errMsg)
	}
	icaTimeOutNanos := k.GetParams(ctx).IcaTimeoutNanos
	nextEpochStarttime := epochTracker.NextEpochStartTime
	timeout := nextEpochStarttime + icaTimeOutNanos

	icaTx := IcaTx{
		ConnectionId: hostZone.GetConnectionId(),
		Msgs:         msgs,
		Account:      *redemptionAccount,
		Timeout:      timeout,
	}

	return &icaTx, nil
}

func (k Keeper) GetRedemptionIBCTransferMsg(ctx sdk.Context, userRedemptionRecord *recordstypes.UserRedemptionRecord, hostZoneId string) (*IcaTx, error) {
	// grab necessary fields to construct ICA call
	hostZone, found := k.GetHostZone(ctx, hostZoneId)
	if !found {
		errMsg := fmt.Sprintf("Host zone %s not found", hostZoneId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrInvalidHostZone, errMsg)
	}
	redemptionAccount, found := k.GetRedemptionAccount(ctx, hostZone)
	if !found {
		errMsg := fmt.Sprintf("Redemption account not found for host zone %s", hostZoneId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrInvalidHostZone, errMsg)
	}

	var msgs []sdk.Msg
	rrAmt, err := cast.ToInt64E(userRedemptionRecord.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidUserRedemptionRecord, err.Error())
	}

	ibcTransferTimeoutNanos := k.GetParams(ctx).IbcTransferTimeoutNanos
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + ibcTransferTimeoutNanos

	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, ibctransfertypes.PortID, hostZone.TransferChannelId)
	if !found {
		return nil, sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", ibctransfertypes.PortID, hostZone.TransferChannelId)
	}

	transferChannelId := sourceChannelEnd.GetCounterparty().GetChannelID()
	msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		SourcePort:       ibctransfertypes.PortID,
		SourceChannel:    transferChannelId,
		Token:            sdk.NewInt64Coin(hostZone.HostDenom, rrAmt),
		Sender:           redemptionAccount.Address,
		Receiver:         userRedemptionRecord.Receiver,
		TimeoutTimestamp: timeoutTimestamp,
	})

	// Give claims a 10 minute timeout
	epochTracker, found := k.GetEpochTracker(ctx, epochstypes.BASE_EPOCH)
	if !found {
		errMsg := fmt.Sprintf("Epoch tracker not found for epoch %s", epochstypes.BASE_EPOCH)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrEpochNotFound, errMsg)
	}
	icaTimeOutNanos := k.GetParams(ctx).IcaTimeoutNanos
	nextEpochStarttime := epochTracker.NextEpochStartTime
	timeout := nextEpochStarttime + icaTimeOutNanos

	icaTx := IcaTx{
		ConnectionId: hostZone.GetConnectionId(),
		Msgs:         msgs,
		Account:      *redemptionAccount,
		Timeout:      timeout,
	}

	return &icaTx, nil
}

func (k Keeper) GetRedemptionIBCTransferMsg2(ctx sdk.Context, msg *types.MsgClaimUndelegatedTokens, hostZoneId string) (*IcaTx, error) {
	// grab necessary fields to construct ICA call
	hostZone, found := k.GetHostZone(ctx, hostZoneId)
	if !found {
		errMsg := fmt.Sprintf("Host zone %s not found", hostZoneId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrInvalidHostZone, errMsg)
	}
	redemptionAccount, found := k.GetRedemptionAccount(ctx, hostZone)
	if !found {
		errMsg := fmt.Sprintf("Redemption account not found for host zone %s", hostZoneId)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrInvalidHostZone, errMsg)
	}

	var msgs []sdk.Msg
	rrAmt := int64(111)

	ibcTransferTimeoutNanos := k.GetParams(ctx).IbcTransferTimeoutNanos
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + ibcTransferTimeoutNanos

	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, ibctransfertypes.PortID, hostZone.TransferChannelId)
	if !found {
		return nil, sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", ibctransfertypes.PortID, hostZone.TransferChannelId)
	}

	transferChannelId := sourceChannelEnd.GetCounterparty().GetChannelID()
	msgs = append(msgs, &ibctransfertypes.MsgTransfer{
		SourcePort:       ibctransfertypes.PortID,
		SourceChannel:    transferChannelId,
		Token:            sdk.NewInt64Coin(hostZone.HostDenom, rrAmt),
		Sender:           redemptionAccount.Address,
		Receiver:         msg.Creator,
		TimeoutTimestamp: timeoutTimestamp,
	})

	// Give claims a 10 minute timeout
	epochTracker, found := k.GetEpochTracker(ctx, epochstypes.BASE_EPOCH)
	if !found {
		errMsg := fmt.Sprintf("Epoch tracker not found for epoch %s", epochstypes.BASE_EPOCH)
		k.Logger(ctx).Error(errMsg)
		return nil, sdkerrors.Wrap(types.ErrEpochNotFound, errMsg)
	}
	icaTimeOutNanos := k.GetParams(ctx).IcaTimeoutNanos
	nextEpochStarttime := epochTracker.NextEpochStartTime
	timeout := nextEpochStarttime + icaTimeOutNanos

	icaTx := IcaTx{
		ConnectionId: hostZone.GetConnectionId(),
		Msgs:         msgs,
		Account:      *redemptionAccount,
		Timeout:      timeout,
	}

	return &icaTx, nil
}
