package types

const (
	// ModuleName defines the module name
	ModuleName = "ibb"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ibb"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PoolKey      = "Pool-value-"
	PoolCountKey = "Pool-count-"
)

const (
	DepositKey      = "Deposit-value-"
	DepositCountKey = "Deposit-count-"
)

const (
	BorrowKey      = "Borrow-value-"
	BorrowCountKey = "Borrow-count-"
)

const (
	UserKey      = "User-value-"
	UserCountKey = "User-count-"
)

const (
	WithdrawKey      = "Withdraw-value-"
	WithdrawCountKey = "Withdraw-count-"
)

const (
	RepayKey      = "Repay-value-"
	RepayCountKey = "Repay-count-"
)

const (
	AprKey      = "Apr-value-"
	AprCountKey = "Apr-count-"
)

const (
	DepositEarnedKey      = "DepositEarned-value-"
	DepositEarnedCountKey = "DepositEarned-count-"
)

const (
	BorrowAccruedKey      = "BorrowAccrued-value-"
	BorrowAccruedCountKey = "BorrowAccrued-count-"
)
