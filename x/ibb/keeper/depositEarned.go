package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetDepositEarnedCount get the total number of depositEarned
func (k Keeper) GetDepositEarnedCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedCountKey))
	byteKey := types.KeyPrefix(types.DepositEarnedCountKey)
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

// SetDepositEarnedCount set the total number of depositEarned
func (k Keeper) SetDepositEarnedCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedCountKey))
	byteKey := types.KeyPrefix(types.DepositEarnedCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendDepositEarned appends a depositEarned in the store with a new id and update the count
func (k Keeper) AppendDepositEarned(
	ctx sdk.Context,
	creator string,
	blockHeight int32,
	asset string,
	amount int32,
	denom string,
) uint64 {
	// Create the depositEarned
	count := k.GetDepositEarnedCount(ctx)
	var depositEarned = types.DepositEarned{
		Creator:     creator,
		Id:          count,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	value := k.cdc.MustMarshalBinaryBare(&depositEarned)
	store.Set(GetDepositEarnedIDBytes(depositEarned.Id), value)

	// Update depositEarned count
	k.SetDepositEarnedCount(ctx, count+1)

	return count
}

// SetDepositEarned set a specific depositEarned in the store
func (k Keeper) SetDepositEarned(ctx sdk.Context, depositEarned types.DepositEarned) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	b := k.cdc.MustMarshalBinaryBare(&depositEarned)
	store.Set(GetDepositEarnedIDBytes(depositEarned.Id), b)
}

// GetDepositEarned returns a depositEarned from its id
func (k Keeper) GetDepositEarned(ctx sdk.Context, id uint64) types.DepositEarned {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	var depositEarned types.DepositEarned
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetDepositEarnedIDBytes(id)), &depositEarned)
	return depositEarned
}

// HasDepositEarned checks if the depositEarned exists in the store
func (k Keeper) HasDepositEarned(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	return store.Has(GetDepositEarnedIDBytes(id))
}

// GetDepositEarnedOwner returns the creator of the depositEarned
func (k Keeper) GetDepositEarnedOwner(ctx sdk.Context, id uint64) string {
	return k.GetDepositEarned(ctx, id).Creator
}

// RemoveDepositEarned removes a depositEarned from the store
func (k Keeper) RemoveDepositEarned(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	store.Delete(GetDepositEarnedIDBytes(id))
}

// GetAllDepositEarned returns all depositEarned
func (k Keeper) GetAllDepositEarned(ctx sdk.Context) (list []types.DepositEarned) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DepositEarned
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDepositEarnedIDBytes returns the byte representation of the ID
func GetDepositEarnedIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDepositEarnedIDFromBytes returns ID in uint64 format from a byte array
func GetDepositEarnedIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
