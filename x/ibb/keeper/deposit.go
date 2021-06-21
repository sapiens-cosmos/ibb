package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetDepositCount get the total number of deposit
func (k Keeper) GetDepositCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositCountKey))
	byteKey := types.KeyPrefix(types.DepositCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetDepositCount set the total number of deposit
func (k Keeper) SetDepositCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositCountKey))
	byteKey := types.KeyPrefix(types.DepositCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendDeposit appends a deposit in the store with a new id and update the count
func (k Keeper) AppendDeposit(
	ctx sdk.Context,
	deposit types.Deposit,
) uint64 {
	// Create the deposit
	count := k.GetDepositCount(ctx)

	// Set the ID of the appended value
	deposit.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&deposit)
	store.Set(GetDepositIDBytes(deposit.Id), appendedValue)

	// Update deposit count
	k.SetDepositCount(ctx, count+1)

	return count
}

// SetDeposit set a specific deposit in the store
func (k Keeper) SetDeposit(ctx sdk.Context, deposit types.Deposit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	b := k.cdc.MustMarshalBinaryBare(&deposit)
	store.Set(GetDepositIDBytes(deposit.Id), b)
}

// GetDeposit returns a deposit from its id
func (k Keeper) GetDeposit(ctx sdk.Context, id uint64) types.Deposit {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	var deposit types.Deposit
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetDepositIDBytes(id)), &deposit)
	return deposit
}

// HasDeposit checks if the deposit exists in the store
func (k Keeper) HasDeposit(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	return store.Has(GetDepositIDBytes(id))
}

// GetDepositOwner returns the creator of the deposit
func (k Keeper) GetDepositOwner(ctx sdk.Context, id uint64) string {
	return k.GetDeposit(ctx, id).Creator
}

// RemoveDeposit removes a deposit from the store
func (k Keeper) RemoveDeposit(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	store.Delete(GetDepositIDBytes(id))
}

// GetAllDeposit returns all deposit
func (k Keeper) GetAllDeposit(ctx sdk.Context) (list []types.Deposit) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Deposit
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDepositIDBytes returns the byte representation of the ID
func GetDepositIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDepositIDFromBytes returns ID in uint64 format from a byte array
func GetDepositIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
