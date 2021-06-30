package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/stretchr/testify/assert"
)

func createNNft(keeper *Keeper, ctx sdk.Context, n int) []types.Nft {
	items := make([]types.Nft, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendNft(ctx, items[i])
	}
	return items
}

func TestNftGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNNft(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetNft(ctx, item.Id))
	}
}

func TestNftExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNNft(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasNft(ctx, item.Id))
	}
}

func TestNftRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNNft(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNft(ctx, item.Id)
		assert.False(t, keeper.HasNft(ctx, item.Id))
	}
}

func TestNftGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNNft(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllNft(ctx))
}

func TestNftCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNNft(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetNftCount(ctx))
}
