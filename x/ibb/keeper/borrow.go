package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetBorrowCount get the total number of borrow
func (k Keeper) GetBorrowCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowCountKey))
	byteKey := types.KeyPrefix(types.BorrowCountKey)
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

// SetBorrowCount set the total number of borrow
func (k Keeper) SetBorrowCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowCountKey))
	byteKey := types.KeyPrefix(types.BorrowCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendBorrow appends a borrow in the store with a new id and update the count
func (k Keeper) AppendBorrow(
	ctx sdk.Context,
	borrow types.Borrow,
) uint64 {
	// Create the borrow
	count := k.GetBorrowCount(ctx)

	// Set the ID of the appended value
	borrow.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&borrow)
	store.Set(GetBorrowIDBytes(borrow.Id), appendedValue)

	// Update borrow count
	k.SetBorrowCount(ctx, count+1)

	return count
}

// SetBorrow set a specific borrow in the store
func (k Keeper) SetBorrow(ctx sdk.Context, borrow types.Borrow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowKey))
	b := k.cdc.MustMarshalBinaryBare(&borrow)
	store.Set(GetBorrowIDBytes(borrow.Id), b)
}

// GetBorrow returns a borrow from its id
func (k Keeper) GetBorrow(ctx sdk.Context, id uint64) types.Borrow {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowKey))
	var borrow types.Borrow
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBorrowIDBytes(id)), &borrow)
	return borrow
}

// HasBorrow checks if the borrow exists in the store
func (k Keeper) HasBorrow(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowKey))
	return store.Has(GetBorrowIDBytes(id))
}

// GetBorrowOwner returns the creator of the borrow
func (k Keeper) GetBorrowOwner(ctx sdk.Context, id uint64) string {
	return k.GetBorrow(ctx, id).Creator
}

// RemoveBorrow removes a borrow from the store
func (k Keeper) RemoveBorrow(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowKey))
	store.Delete(GetBorrowIDBytes(id))
}

// GetAllBorrow returns all borrow
func (k Keeper) GetAllBorrow(ctx sdk.Context) (list []types.Borrow) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Borrow
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBorrowIDBytes returns the byte representation of the ID
func GetBorrowIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBorrowIDFromBytes returns ID in uint64 format from a byte array
func GetBorrowIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
