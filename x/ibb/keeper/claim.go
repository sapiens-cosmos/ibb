package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"strconv"
)

// GetClaimCount get the total number of claim
func (k Keeper) GetClaimCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimCountKey))
	byteKey := types.KeyPrefix(types.ClaimCountKey)
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

// SetClaimCount set the total number of claim
func (k Keeper) SetClaimCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimCountKey))
	byteKey := types.KeyPrefix(types.ClaimCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendClaim appends a claim in the store with a new id and update the count
func (k Keeper) AppendClaim(
	ctx sdk.Context,
	creator string,
	blockHeight int32,
	asset string,
	amount int32,
	denom string,
) uint64 {
	// Create the claim
	count := k.GetClaimCount(ctx)
	var claim = types.Claim{
		Creator:     creator,
		Id:          count,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	value := k.cdc.MustMarshalBinaryBare(&claim)
	store.Set(GetClaimIDBytes(claim.Id), value)

	// Update claim count
	k.SetClaimCount(ctx, count+1)

	return count
}

// SetClaim set a specific claim in the store
func (k Keeper) SetClaim(ctx sdk.Context, claim types.Claim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	b := k.cdc.MustMarshalBinaryBare(&claim)
	store.Set(GetClaimIDBytes(claim.Id), b)
}

// GetClaim returns a claim from its id
func (k Keeper) GetClaim(ctx sdk.Context, id uint64) types.Claim {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	var claim types.Claim
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetClaimIDBytes(id)), &claim)
	return claim
}

// HasClaim checks if the claim exists in the store
func (k Keeper) HasClaim(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	return store.Has(GetClaimIDBytes(id))
}

// GetClaimOwner returns the creator of the claim
func (k Keeper) GetClaimOwner(ctx sdk.Context, id uint64) string {
	return k.GetClaim(ctx, id).Creator
}

// RemoveClaim removes a claim from the store
func (k Keeper) RemoveClaim(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	store.Delete(GetClaimIDBytes(id))
}

// GetAllClaim returns all claim
func (k Keeper) GetAllClaim(ctx sdk.Context) (list []types.Claim) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ClaimKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Claim
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetClaimIDBytes returns the byte representation of the ID
func GetClaimIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetClaimIDFromBytes returns ID in uint64 format from a byte array
func GetClaimIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
