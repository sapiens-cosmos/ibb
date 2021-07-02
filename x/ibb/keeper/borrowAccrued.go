package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetBorrowAccruedCount get the total number of borrowAccrued
func (k Keeper) GetBorrowAccruedCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedCountKey))
	byteKey := types.KeyPrefix(types.BorrowAccruedCountKey)
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

// SetBorrowAccruedCount set the total number of borrowAccrued
func (k Keeper) SetBorrowAccruedCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedCountKey))
	byteKey := types.KeyPrefix(types.BorrowAccruedCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendBorrowAccrued appends a borrowAccrued in the store with a new id and update the count
func (k Keeper) AppendBorrowAccrued(
	ctx sdk.Context,
	creator string,
	blockHeight int32,
	asset string,
	amount int32,
	denom string,
) uint64 {
	// Create the borrowAccrued
	count := k.GetBorrowAccruedCount(ctx)
	var borrowAccrued = types.BorrowAccrued{
		Creator:     creator,
		Id:          count,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedKey))
	value := k.cdc.MustMarshalBinaryBare(&borrowAccrued)
	store.Set(GetBorrowAccruedIDBytes(borrowAccrued.Id), value)

	// Update borrowAccrued count
	k.SetBorrowAccruedCount(ctx, count+1)

	return count
}

// SetBorrowAccrued set a specific borrowAccrued in the store
func (k Keeper) SetBorrowAccrued(ctx sdk.Context, borrowAccrued types.BorrowAccrued) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedKey))
	b := k.cdc.MustMarshalBinaryBare(&borrowAccrued)
	store.Set(GetBorrowAccruedIDBytes(borrowAccrued.Id), b)
}

// GetBorrowAccrued returns a borrowAccrued from its id
func (k Keeper) GetBorrowAccrued(ctx sdk.Context, id uint64) types.BorrowAccrued {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedKey))
	var borrowAccrued types.BorrowAccrued
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBorrowAccruedIDBytes(id)), &borrowAccrued)
	return borrowAccrued
}

// HasBorrowAccrued checks if the borrowAccrued exists in the store
func (k Keeper) HasBorrowAccrued(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedKey))
	return store.Has(GetBorrowAccruedIDBytes(id))
}

// GetBorrowAccruedOwner returns the creator of the borrowAccrued
func (k Keeper) GetBorrowAccruedOwner(ctx sdk.Context, id uint64) string {
	return k.GetBorrowAccrued(ctx, id).Creator
}

// RemoveBorrowAccrued removes a borrowAccrued from the store
func (k Keeper) RemoveBorrowAccrued(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedKey))
	store.Delete(GetBorrowAccruedIDBytes(id))
}

// GetAllBorrowAccrued returns all borrowAccrued
func (k Keeper) GetAllBorrowAccrued(ctx sdk.Context) (list []types.BorrowAccrued) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BorrowAccrued
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBorrowAccruedIDBytes returns the byte representation of the ID
func GetBorrowAccruedIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBorrowAccruedIDFromBytes returns ID in uint64 format from a byte array
func GetBorrowAccruedIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
