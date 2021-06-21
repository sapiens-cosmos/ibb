package ibb

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/keeper"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the pool
	for _, elem := range genState.PoolList {
		k.SetPool(ctx, *elem)
	}

	// Set pool count
	k.SetPoolCount(ctx, genState.PoolCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all pool
	poolList := k.GetAllPool(ctx)
	for _, elem := range poolList {
		elem := elem
		genesis.PoolList = append(genesis.PoolList, &elem)
	}

	// Set the current count
	genesis.PoolCount = k.GetPoolCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}