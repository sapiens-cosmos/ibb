package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/stretchr/testify/assert"
)

func createNInterfaceApr(keeper *Keeper, ctx sdk.Context, n int) []types.InterfaceApr {
	items := make([]types.InterfaceApr, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendInterfaceApr(ctx, items[i])
	}
	return items
}

func TestInterfaceAprGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNInterfaceApr(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetInterfaceApr(ctx, item.Id))
	}
}

func TestInterfaceAprExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNInterfaceApr(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasInterfaceApr(ctx, item.Id))
	}
}

func TestInterfaceAprRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNInterfaceApr(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInterfaceApr(ctx, item.Id)
		assert.False(t, keeper.HasInterfaceApr(ctx, item.Id))
	}
}

func TestInterfaceAprGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNInterfaceApr(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllInterfaceApr(ctx))
}

func TestInterfaceAprCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNInterfaceApr(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetInterfaceAprCount(ctx))
}
