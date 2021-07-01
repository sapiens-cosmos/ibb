package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetTxHistoryCount get the total number of txHistory
func (k Keeper) GetTxHistoryCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryCountKey))
	byteKey := types.KeyPrefix(types.TxHistoryCountKey)
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

// SetTxHistoryCount set the total number of txHistory
func (k Keeper) SetTxHistoryCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryCountKey))
	byteKey := types.KeyPrefix(types.TxHistoryCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendTxHistory appends a txHistory in the store with a new id and update the count
func (k Keeper) AppendTxHistory(
	ctx sdk.Context,
	creator string,
	blockHeight int32,
	tx string,
	asset string,
	amount int32,
	denom string,
) uint64 {
	// Create the txHistory
	count := k.GetTxHistoryCount(ctx)
	var txHistory = types.TxHistory{
		Creator:     creator,
		Id:          count,
		BlockHeight: blockHeight,
		Tx:          tx,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	value := k.cdc.MustMarshalBinaryBare(&txHistory)
	store.Set(GetTxHistoryIDBytes(txHistory.Id), value)

	// Update txHistory count
	k.SetTxHistoryCount(ctx, count+1)

	return count
}

// SetTxHistory set a specific txHistory in the store
func (k Keeper) SetTxHistory(ctx sdk.Context, txHistory types.TxHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	b := k.cdc.MustMarshalBinaryBare(&txHistory)
	store.Set(GetTxHistoryIDBytes(txHistory.Id), b)
}

// GetTxHistory returns a txHistory from its id
func (k Keeper) GetTxHistory(ctx sdk.Context, id uint64) types.TxHistory {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	var txHistory types.TxHistory
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetTxHistoryIDBytes(id)), &txHistory)
	return txHistory
}

// HasTxHistory checks if the txHistory exists in the store
func (k Keeper) HasTxHistory(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	return store.Has(GetTxHistoryIDBytes(id))
}

// GetTxHistoryOwner returns the creator of the txHistory
func (k Keeper) GetTxHistoryOwner(ctx sdk.Context, id uint64) string {
	return k.GetTxHistory(ctx, id).Creator
}

// RemoveTxHistory removes a txHistory from the store
func (k Keeper) RemoveTxHistory(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	store.Delete(GetTxHistoryIDBytes(id))
}

// GetAllTxHistory returns all txHistory
func (k Keeper) GetAllTxHistory(ctx sdk.Context) (list []types.TxHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TxHistory
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTxHistoryIDBytes returns the byte representation of the ID
func GetTxHistoryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTxHistoryIDFromBytes returns ID in uint64 format from a byte array
func GetTxHistoryIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
