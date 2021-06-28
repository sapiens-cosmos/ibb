package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetWithdrawCount get the total number of withdraw
func (k Keeper) GetWithdrawCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawCountKey))
	byteKey := types.KeyPrefix(types.WithdrawCountKey)
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

// SetWithdrawCount set the total number of withdraw
func (k Keeper) SetWithdrawCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawCountKey))
	byteKey := types.KeyPrefix(types.WithdrawCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendWithdraw appends a withdraw in the store with a new id and update the count
func (k Keeper) AppendWithdraw(
	ctx sdk.Context,
	creator string,
	blockHeight int32,
	asset string,
	amount int32,
	denom string,
) uint64 {
	// Create the withdraw
	count := k.GetWithdrawCount(ctx)
	var withdraw = types.Withdraw{
		Creator:     creator,
		Id:          count,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKey))
	value := k.cdc.MustMarshalBinaryBare(&withdraw)
	store.Set(GetWithdrawIDBytes(withdraw.Id), value)

	// Update withdraw count
	k.SetWithdrawCount(ctx, count+1)

	return count
}

// SetWithdraw set a specific withdraw in the store
func (k Keeper) SetWithdraw(ctx sdk.Context, withdraw types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKey))
	b := k.cdc.MustMarshalBinaryBare(&withdraw)
	store.Set(GetWithdrawIDBytes(withdraw.Id), b)
}

// GetWithdraw returns a withdraw from its id
func (k Keeper) GetWithdraw(ctx sdk.Context, id uint64) types.Withdraw {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKey))
	var withdraw types.Withdraw
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetWithdrawIDBytes(id)), &withdraw)
	return withdraw
}

// HasWithdraw checks if the withdraw exists in the store
func (k Keeper) HasWithdraw(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKey))
	return store.Has(GetWithdrawIDBytes(id))
}

// GetWithdrawOwner returns the creator of the withdraw
func (k Keeper) GetWithdrawOwner(ctx sdk.Context, id uint64) string {
	return k.GetWithdraw(ctx, id).Creator
}

// RemoveWithdraw removes a withdraw from the store
func (k Keeper) RemoveWithdraw(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKey))
	store.Delete(GetWithdrawIDBytes(id))
}

// GetAllWithdraw returns all withdraw
func (k Keeper) GetAllWithdraw(ctx sdk.Context) (list []types.Withdraw) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WithdrawKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Withdraw
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetWithdrawIDBytes returns the byte representation of the ID
func GetWithdrawIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetWithdrawIDFromBytes returns ID in uint64 format from a byte array
func GetWithdrawIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
