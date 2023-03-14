package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/UnUniFi/chain/x/yield-aggregator/types"
)

// GetStrategyCount get the total number of Strategy
func (k Keeper) GetStrategyCount(ctx sdk.Context, vaultDenom string) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixStrategyCount(vaultDenom)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetStrategyCount set the total number of Strategy
func (k Keeper) SetStrategyCount(ctx sdk.Context, vaultDenom string, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefixStrategyCount(vaultDenom)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendStrategy appends a Strategy in the store with a new id and update the count
func (k Keeper) AppendStrategy(
	ctx sdk.Context,
	vaultDenom string,
	Strategy types.Strategy,
) uint64 {
	// Create the Strategy
	count := k.GetStrategyCount(ctx, vaultDenom)

	// Set the ID of the appended value
	Strategy.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	appendedValue := k.cdc.MustMarshal(&Strategy)
	store.Set(GetStrategyIDBytes(Strategy.Id), appendedValue)

	// Update Strategy count
	k.SetStrategyCount(ctx, vaultDenom, count+1)

	return count
}

// SetStrategy set a specific Strategy in the store
func (k Keeper) SetStrategy(ctx sdk.Context, vaultDenom string, Strategy types.Strategy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	b := k.cdc.MustMarshal(&Strategy)
	store.Set(GetStrategyIDBytes(Strategy.Id), b)
}

// GetStrategy returns a Strategy from its id
func (k Keeper) GetStrategy(ctx sdk.Context, vaultDenom string, id uint64) (val types.Strategy, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	b := store.Get(GetStrategyIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStrategy removes a Strategy from the store
func (k Keeper) RemoveStrategy(ctx sdk.Context, vaultDenom string, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	store.Delete(GetStrategyIDBytes(id))
}

// GetAllStrategy returns all Strategy
func (k Keeper) GetAllStrategy(ctx sdk.Context, vaultDenom string) (list []types.Strategy) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStrategy(vaultDenom))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Strategy
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetStrategyIDBytes returns the byte representation of the ID
func GetStrategyIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetStrategyIDFromBytes returns ID in uint64 format from a byte array
func GetStrategyIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

// stake into strategy
func (k Keeper) StakeToStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy, amount sdk.Int) error {
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		vaultModName := types.GetVaultModuleAccountName(vault.Id)
		vaultModAddr := authtypes.NewModuleAddress(vaultModName)

		return k.stakeibcKeeper.LiquidStake(
			ctx,
			vaultModAddr,
			sdk.NewCoin(vault.Denom, amount),
		)
	default:
		panic("not implemented")
	}

	// TODO: call `stake` function of the strategy contract
	// 		wasmMsg := `{"stake":{}}`
	// 		contractAddr := sdk.MustAccAddressFromBech32(target.AccountAddress)
	// 		_, err := k.wasmKeeper.Execute(ctx, contractAddr, farmingUnit.GetAddress(), []byte(wasmMsg), amount)
	// 		if err != nil {
	// 			return err
	// 		}

	return nil
}

// unstake worth of withdrawal amount from the strategy
func (k Keeper) UnstakeFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy, amount sdk.Int) error {
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		{
			vaultModName := types.GetVaultModuleAccountName(vault.Id)
			vaultModAddr := authtypes.NewModuleAddress(vaultModName)
			err := k.stakeibcKeeper.RedeemStake(
				ctx,
				vaultModAddr,
				sdk.NewCoin(vault.Denom, amount),
				vaultModAddr.String(),
			)
			if err != nil {
				return err
			}

			return nil
		}
	}
	panic("not implemented")
	// TODO: call `unstake` function of the strategy contract
	// 		wasmMsg := `{"unstake":{}}`
	// 		contractAddr := sdk.MustAccAddressFromBech32(target.AccountAddress)
	// 		_, err := k.wasmKeeper.Execute(ctx, contractAddr, farmingUnit.GetAddress(), []byte(wasmMsg), sdk.Coins{})
	// 		if err != nil {
	// 			return err
	// 		}

	return nil
}

func (k Keeper) GetAmountFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy) (sdk.Coin, error) {
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		{
			vaultModName := types.GetVaultModuleAccountName(vault.Id)
			vaultModAddr := authtypes.NewModuleAddress(vaultModName)
			updatedAmount := k.stakeibcKeeper.GetUpdatedBalance(ctx, vaultModAddr, vault.Denom)
			return sdk.NewCoin(vault.Denom, updatedAmount), nil
		}
	}
	// call `amount` function of the strategy contract
	panic("not implemented")

	return sdk.Coin{}, nil
}

func (k Keeper) GetUnbondingAmountFromStrategy(ctx sdk.Context, vault types.Vault, strategy types.Strategy) (sdk.Coin, error) {
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		{
			vaultModName := types.GetVaultModuleAccountName(vault.Id)
			vaultModAddr := authtypes.NewModuleAddress(vaultModName)
			unbondingAmount := k.recordsKeeper.GetUserRedemptionRecordBySenderAndDenom(ctx, vaultModAddr, vault.Denom)
			return sdk.NewCoin(vault.Denom, unbondingAmount), nil
		}
	}
	// call `amount` function of the strategy contract
	panic("not implemented")
	// 		wasmQuery := `{"amount":{}}`
	// 		contractAddr := sdk.MustAccAddressFromBech32(target.AccountAddress)
	// 		_, err := k.wasmKeeper.Execute(ctx, contractAddr, farmingUnit.GetAddress(), []byte(wasmQuery), sdk.Coins{})
	// 		if err != nil {
	// 			return err
	// 		}

	return sdk.Coin{}, nil
}

func (k Keeper) GetAPRFromStrategy(ctx sdk.Context, strategy types.Strategy) (*sdk.Dec, error) {
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		{

			return nil, nil
		}
	}
	// call `apr` function of the strategy contract
	panic("not implemented")

	return nil, nil
}

func (k Keeper) GetInterestFeeRate(ctx sdk.Context, strategy types.Strategy) (*sdk.Dec, error) {
	switch strategy.ContractAddress {
	case "x/ibc-staking":
		{

			return nil, nil
		}
	}
	// call `interest_fee_rate` function of the strategy contract
	panic("not implemented")

	return nil, nil
}
