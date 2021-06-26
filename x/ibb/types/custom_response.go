package types

type LoadPoolResponse struct {
	Asset            string
	CollatoralFactor int32
	Liquidity        int32
	DepositApy       int32
	BorrowApy        int32
}
