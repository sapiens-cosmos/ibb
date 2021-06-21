package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/stretchr/testify/assert"
)

func createNBorrow(keeper *Keeper, ctx sdk.Context, n int) []types.Borrow {
	items := make([]types.Borrow, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendBorrow(ctx, items[i])
	}
	return items
}

func TestBorrowGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBorrow(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetBorrow(ctx, item.Id))
	}
}

func TestBorrowExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBorrow(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasBorrow(ctx, item.Id))
	}
}

func TestBorrowRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBorrow(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBorrow(ctx, item.Id)
		assert.False(t, keeper.HasBorrow(ctx, item.Id))
	}
}

func TestBorrowGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBorrow(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllBorrow(ctx))
}

func TestBorrowCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBorrow(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetBorrowCount(ctx))
}
