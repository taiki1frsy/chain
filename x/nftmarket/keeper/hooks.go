package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/nftmarket/types"
)

var _ types.NftmarketHooks = Keeper{}

func (k Keeper) AfterNftListed(ctx sdk.Context, nftIdentifier []byte, txMemo string) {
	if k.hooks != nil {
		k.hooks.AfterNftListed(ctx, nftIdentifier, txMemo)
	}
}

func (k Keeper) AfterNftPaymentWithCommission(ctx sdk.Context, nftIdentifier []byte, fee sdk.Coin) {
	if k.hooks != nil {
		k.hooks.AfterNftPaymentWithCommission(ctx, nftIdentifier, fee)
	}
}

func (k Keeper) AfterNftUnlistedWithoutPayment(ctx sdk.Context, nftIdentifier []byte) {
	if k.hooks != nil {
		k.hooks.AfterNftUnlistedWithoutPayment(ctx, nftIdentifier)
	}
}
