package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/yieldaggregator module sentinel errors
var (
	ErrInvalidFeeDenom            = sdkerrors.Register(ModuleName, 1, "invalid fee denom")
	ErrInsufficientFee            = sdkerrors.Register(ModuleName, 2, "insufficient fee")
	ErrInvalidDepositDenom        = sdkerrors.Register(ModuleName, 3, "invalid deposit denom")
	ErrInsufficientDeposit        = sdkerrors.Register(ModuleName, 4, "insufficient deposit")
	ErrInvalidVaultId             = sdkerrors.Register(ModuleName, 5, "invalid vault id")
	ErrNotVaultOwner              = sdkerrors.Register(ModuleName, 6, "not a vault owner")
	ErrVaultHasPositiveBalance    = sdkerrors.Register(ModuleName, 7, "vault has positive balance")
	ErrInvalidCommissionRate      = sdkerrors.Register(ModuleName, 8, "invalid commission rate")
	ErrDuplicatedStrategy         = sdkerrors.Register(ModuleName, 9, "duplicated strategy")
	ErrInvalidStrategyWeightSum   = sdkerrors.Register(ModuleName, 10, "invalid strategy weight sum")
	ErrInvalidStrategyInvolved    = sdkerrors.Register(ModuleName, 11, "invalid strategy id involved")
	ErrInvalidWithdrawReserveRate = sdkerrors.Register(ModuleName, 12, "invalid withdraw reserve rate")
)
