package types

type LoadPoolRestResponse struct {
	Asset            string
	CollatoralFactor int32
	Liquidity        int32
	DepositApy       int32
	BorrowApy        int32
}

type LoadUserResponse struct {
	AssetApy     int32
	AssetBalance int32
	AssetDeposit int32
	AssetBorrow  int32
	Collateral   bool
}
