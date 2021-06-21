package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/stretchr/testify/assert"
)

func createNPool(keeper *Keeper, ctx sdk.Context, n int) []types.Pool {
	items := make([]types.Pool, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendPool(ctx, items[i])
	}
	return items
}

func TestPoolGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetPool(ctx, item.Id))
	}
}

func TestPoolExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasPool(ctx, item.Id))
	}
}

func TestPoolRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPool(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePool(ctx, item.Id)
		assert.False(t, keeper.HasPool(ctx, item.Id))
	}
}

func TestPoolGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPool(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllPool(ctx))
}

func TestPoolCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNPool(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetPoolCount(ctx))
}
