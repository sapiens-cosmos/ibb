package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"github.com/stretchr/testify/assert"
)

func createNDeposit(keeper *Keeper, ctx sdk.Context, n int) []types.Deposit {
	items := make([]types.Deposit, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendDeposit(ctx, items[i])
	}
	return items
}

func TestDepositGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetDeposit(ctx, item.Id))
	}
}

func TestDepositExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasDeposit(ctx, item.Id))
	}
}

func TestDepositRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeposit(ctx, item.Id)
		assert.False(t, keeper.HasDeposit(ctx, item.Id))
	}
}

func TestDepositGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllDeposit(ctx))
}

func TestDepositCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNDeposit(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetDepositCount(ctx))
}
