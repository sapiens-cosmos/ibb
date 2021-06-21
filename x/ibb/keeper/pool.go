package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetPoolCount get the total number of pool
func (k Keeper) GetPoolCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolCountKey))
	byteKey := types.KeyPrefix(types.PoolCountKey)
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

// SetPoolCount set the total number of pool
func (k Keeper) SetPoolCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolCountKey))
	byteKey := types.KeyPrefix(types.PoolCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendPool appends a pool in the store with a new id and update the count
func (k Keeper) AppendPool(
	ctx sdk.Context,
	pool types.Pool,
) uint64 {
	// Create the pool
	count := k.GetPoolCount(ctx)

	// Set the ID of the appended value
	pool.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&pool)
	store.Set(GetPoolIDBytes(pool.Id), appendedValue)

	// Update pool count
	k.SetPoolCount(ctx, count+1)

	return count
}

// SetPool set a specific pool in the store
func (k Keeper) SetPool(ctx sdk.Context, pool types.Pool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	b := k.cdc.MustMarshalBinaryBare(&pool)
	store.Set(GetPoolIDBytes(pool.Id), b)
}

// GetPool returns a pool from its id
func (k Keeper) GetPool(ctx sdk.Context, id uint64) types.Pool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	var pool types.Pool
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetPoolIDBytes(id)), &pool)
	return pool
}

// HasPool checks if the pool exists in the store
func (k Keeper) HasPool(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	return store.Has(GetPoolIDBytes(id))
}

// GetPoolOwner returns the creator of the pool
func (k Keeper) GetPoolOwner(ctx sdk.Context, id uint64) string {
	return k.GetPool(ctx, id).Creator
}

// RemovePool removes a pool from the store
func (k Keeper) RemovePool(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	store.Delete(GetPoolIDBytes(id))
}

// GetAllPool returns all pool
func (k Keeper) GetAllPool(ctx sdk.Context) (list []types.Pool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Pool
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetPoolIDBytes returns the byte representation of the ID
func GetPoolIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetPoolIDFromBytes returns ID in uint64 format from a byte array
func GetPoolIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
