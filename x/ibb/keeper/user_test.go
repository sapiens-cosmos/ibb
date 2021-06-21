package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/stretchr/testify/assert"
)

func createNUser(keeper *Keeper, ctx sdk.Context, n int) []types.User {
	items := make([]types.User, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendUser(ctx, items[i])
	}
	return items
}

func TestUserGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUser(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetUser(ctx, item.Id))
	}
}

func TestUserExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUser(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasUser(ctx, item.Id))
	}
}

func TestUserRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUser(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUser(ctx, item.Id)
		assert.False(t, keeper.HasUser(ctx, item.Id))
	}
}

func TestUserGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUser(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllUser(ctx))
}

func TestUserCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNUser(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetUserCount(ctx))
}
