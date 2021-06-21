package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetInterfaceAprCount get the total number of interfaceApr
func (k Keeper) GetInterfaceAprCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprCountKey))
	byteKey := types.KeyPrefix(types.InterfaceAprCountKey)
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

// SetInterfaceAprCount set the total number of interfaceApr
func (k Keeper) SetInterfaceAprCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprCountKey))
	byteKey := types.KeyPrefix(types.InterfaceAprCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendInterfaceApr appends a interfaceApr in the store with a new id and update the count
func (k Keeper) AppendInterfaceApr(
	ctx sdk.Context,
	interfaceApr types.InterfaceApr,
) uint64 {
	// Create the interfaceApr
	count := k.GetInterfaceAprCount(ctx)

	// Set the ID of the appended value
	interfaceApr.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&interfaceApr)
	store.Set(GetInterfaceAprIDBytes(interfaceApr.Id), appendedValue)

	// Update interfaceApr count
	k.SetInterfaceAprCount(ctx, count+1)

	return count
}

// SetInterfaceApr set a specific interfaceApr in the store
func (k Keeper) SetInterfaceApr(ctx sdk.Context, interfaceApr types.InterfaceApr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprKey))
	b := k.cdc.MustMarshalBinaryBare(&interfaceApr)
	store.Set(GetInterfaceAprIDBytes(interfaceApr.Id), b)
}

// GetInterfaceApr returns a interfaceApr from its id
func (k Keeper) GetInterfaceApr(ctx sdk.Context, id uint64) types.InterfaceApr {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprKey))
	var interfaceApr types.InterfaceApr
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetInterfaceAprIDBytes(id)), &interfaceApr)
	return interfaceApr
}

// HasInterfaceApr checks if the interfaceApr exists in the store
func (k Keeper) HasInterfaceApr(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprKey))
	return store.Has(GetInterfaceAprIDBytes(id))
}

// GetInterfaceAprOwner returns the creator of the interfaceApr
func (k Keeper) GetInterfaceAprOwner(ctx sdk.Context, id uint64) string {
	return k.GetInterfaceApr(ctx, id).Creator
}

// RemoveInterfaceApr removes a interfaceApr from the store
func (k Keeper) RemoveInterfaceApr(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprKey))
	store.Delete(GetInterfaceAprIDBytes(id))
}

// GetAllInterfaceApr returns all interfaceApr
func (k Keeper) GetAllInterfaceApr(ctx sdk.Context) (list []types.InterfaceApr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InterfaceAprKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.InterfaceApr
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetInterfaceAprIDBytes returns the byte representation of the ID
func GetInterfaceAprIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetInterfaceAprIDFromBytes returns ID in uint64 format from a byte array
func GetInterfaceAprIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
