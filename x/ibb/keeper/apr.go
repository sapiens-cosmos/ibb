package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetAprCount get the total number of apr
func (k Keeper) GetAprCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprCountKey))
	byteKey := types.KeyPrefix(types.AprCountKey)
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

// SetAprCount set the total number of apr
func (k Keeper) SetAprCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprCountKey))
	byteKey := types.KeyPrefix(types.AprCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendApr appends a apr in the store with a new id and update the count
func (k Keeper) AppendApr(
	ctx sdk.Context,
	creator string,
	blockHeight int32,
	depositApy int32,
	borrowApy int32,
) uint64 {
	// Create the apr
	count := k.GetAprCount(ctx)
	var apr = types.Apr{
		Creator:     creator,
		Id:          count,
		BlockHeight: blockHeight,
		DepositApy:  depositApy,
		BorrowApy:   borrowApy,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	value := k.cdc.MustMarshalBinaryBare(&apr)
	store.Set(GetAprIDBytes(apr.Id), value)

	// Update apr count
	k.SetAprCount(ctx, count+1)

	return count
}

// SetApr set a specific apr in the store
func (k Keeper) SetApr(ctx sdk.Context, apr types.Apr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	b := k.cdc.MustMarshalBinaryBare(&apr)
	store.Set(GetAprIDBytes(apr.Id), b)
}

// GetApr returns a apr from its id
func (k Keeper) GetApr(ctx sdk.Context, id uint64) types.Apr {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	var apr types.Apr
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAprIDBytes(id)), &apr)
	return apr
}

// HasApr checks if the apr exists in the store
func (k Keeper) HasApr(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	return store.Has(GetAprIDBytes(id))
}

// GetAprOwner returns the creator of the apr
func (k Keeper) GetAprOwner(ctx sdk.Context, id uint64) string {
	return k.GetApr(ctx, id).Creator
}

// RemoveApr removes a apr from the store
func (k Keeper) RemoveApr(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	store.Delete(GetAprIDBytes(id))
}

// GetAllApr returns all apr
func (k Keeper) GetAllApr(ctx sdk.Context) (list []types.Apr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Apr
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAprIDBytes returns the byte representation of the ID
func GetAprIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAprIDFromBytes returns ID in uint64 format from a byte array
func GetAprIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
