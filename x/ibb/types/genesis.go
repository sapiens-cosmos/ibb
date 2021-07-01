package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		DepositEarnedList: []*DepositEarned{},
		AprList:           []*Apr{},
		RepayList:         []*Repay{},
		WithdrawList:      []*Withdraw{},
		UserList:          []*User{},
		BorrowList:        []*Borrow{},
		DepositList:       []*Deposit{},
		PoolList:          []*Pool{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in depositEarned
	depositEarnedIdMap := make(map[uint64]bool)

	for _, elem := range gs.DepositEarnedList {
		if _, ok := depositEarnedIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for depositEarned")
		}
		depositEarnedIdMap[elem.Id] = true
	}
	// Check for duplicated ID in apr
	aprIdMap := make(map[uint64]bool)

	for _, elem := range gs.AprList {
		if _, ok := aprIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for apr")
		}
		aprIdMap[elem.Id] = true
	}
	// Check for duplicated ID in repay
	repayIdMap := make(map[uint64]bool)

	for _, elem := range gs.RepayList {
		if _, ok := repayIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for repay")
		}
		repayIdMap[elem.Id] = true
	}
	// Check for duplicated ID in withdraw
	withdrawIdMap := make(map[uint64]bool)

	for _, elem := range gs.WithdrawList {
		if _, ok := withdrawIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for withdraw")
		}
		withdrawIdMap[elem.Id] = true
	}
	// Check for duplicated ID in user
	userIdMap := make(map[uint64]bool)

	for _, elem := range gs.UserList {
		if _, ok := userIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for user")
		}
		userIdMap[elem.Id] = true
	}
	// Check for duplicated ID in borrow
	borrowIdMap := make(map[uint64]bool)

	for _, elem := range gs.BorrowList {
		if _, ok := borrowIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for borrow")
		}
		borrowIdMap[elem.Id] = true
	}
	// Check for duplicated ID in deposit
	depositIdMap := make(map[uint64]bool)

	for _, elem := range gs.DepositList {
		if _, ok := depositIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for deposit")
		}
		depositIdMap[elem.Id] = true
	}
	// Check for duplicated ID in pool
	poolIdMap := make(map[uint64]bool)

	for _, elem := range gs.PoolList {
		if _, ok := poolIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for pool")
		}
		poolIdMap[elem.Id] = true
	}

	return nil
}
