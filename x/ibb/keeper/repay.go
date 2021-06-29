package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetRepayCount get the total number of repay
func (k Keeper) GetRepayCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayCountKey))
	byteKey := types.KeyPrefix(types.RepayCountKey)
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

// SetRepayCount set the total number of repay
func (k Keeper) SetRepayCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayCountKey))
	byteKey := types.KeyPrefix(types.RepayCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendRepay appends a repay in the store with a new id and update the count
func (k Keeper) AppendRepay(
	ctx sdk.Context,
	creator string,
	blockHeight int32,
	asset string,
	amount int32,
	denom string,
) uint64 {
	// Create the repay
	count := k.GetRepayCount(ctx)
	var repay = types.Repay{
		Creator:     creator,
		Id:          count,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayKey))
	value := k.cdc.MustMarshalBinaryBare(&repay)
	store.Set(GetRepayIDBytes(repay.Id), value)

	// Update repay count
	k.SetRepayCount(ctx, count+1)

	return count
}

// SetRepay set a specific repay in the store
func (k Keeper) SetRepay(ctx sdk.Context, repay types.Repay) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayKey))
	b := k.cdc.MustMarshalBinaryBare(&repay)
	store.Set(GetRepayIDBytes(repay.Id), b)
}

// GetRepay returns a repay from its id
func (k Keeper) GetRepay(ctx sdk.Context, id uint64) types.Repay {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayKey))
	var repay types.Repay
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetRepayIDBytes(id)), &repay)
	return repay
}

// HasRepay checks if the repay exists in the store
func (k Keeper) HasRepay(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayKey))
	return store.Has(GetRepayIDBytes(id))
}

// GetRepayOwner returns the creator of the repay
func (k Keeper) GetRepayOwner(ctx sdk.Context, id uint64) string {
	return k.GetRepay(ctx, id).Creator
}

// RemoveRepay removes a repay from the store
func (k Keeper) RemoveRepay(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayKey))
	store.Delete(GetRepayIDBytes(id))
}

// GetAllRepay returns all repay
func (k Keeper) GetAllRepay(ctx sdk.Context) (list []types.Repay) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RepayKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Repay
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRepayIDBytes returns the byte representation of the ID
func GetRepayIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRepayIDFromBytes returns ID in uint64 format from a byte array
func GetRepayIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
