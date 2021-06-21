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
	// Set all the interfaceApr
	for _, elem := range genState.InterfaceAprList {
		k.SetInterfaceApr(ctx, *elem)
	}

	// Set interfaceApr count
	k.SetInterfaceAprCount(ctx, genState.InterfaceAprCount)

	// Set all the user
	for _, elem := range genState.UserList {
		k.SetUser(ctx, *elem)
	}

	// Set user count
	k.SetUserCount(ctx, genState.UserCount)

	// Set all the borrow
	for _, elem := range genState.BorrowList {
		k.SetBorrow(ctx, *elem)
	}

	// Set borrow count
	k.SetBorrowCount(ctx, genState.BorrowCount)

	// Set all the deposit
	for _, elem := range genState.DepositList {
		k.SetDeposit(ctx, *elem)
	}

	// Set deposit count
	k.SetDepositCount(ctx, genState.DepositCount)

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
	// Get all interfaceApr
	interfaceAprList := k.GetAllInterfaceApr(ctx)
	for _, elem := range interfaceAprList {
		elem := elem
		genesis.InterfaceAprList = append(genesis.InterfaceAprList, &elem)
	}

	// Set the current count
	genesis.InterfaceAprCount = k.GetInterfaceAprCount(ctx)

	// Get all user
	userList := k.GetAllUser(ctx)
	for _, elem := range userList {
		elem := elem
		genesis.UserList = append(genesis.UserList, &elem)
	}

	// Set the current count
	genesis.UserCount = k.GetUserCount(ctx)

	// Get all borrow
	borrowList := k.GetAllBorrow(ctx)
	for _, elem := range borrowList {
		elem := elem
		genesis.BorrowList = append(genesis.BorrowList, &elem)
	}

	// Set the current count
	genesis.BorrowCount = k.GetBorrowCount(ctx)

	// Get all deposit
	depositList := k.GetAllDeposit(ctx)
	for _, elem := range depositList {
		elem := elem
		genesis.DepositList = append(genesis.DepositList, &elem)
	}

	// Set the current count
	genesis.DepositCount = k.GetDepositCount(ctx)

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
