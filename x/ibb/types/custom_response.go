package types

type LoadPoolRestResponse struct {
	Asset            string
	CollatoralFactor int32
	Liquidity        int32
	DepositApy       int32
	BorrowApy        int32
	AssetPrice       int32
}

type LoadUserRestResponse struct {
	AssetApy     int32
	AssetDenom   string
	AssetBalance int64
	AssetDeposit int32
	AssetBorrow  int32
	AssetPrice   int32
	Collateral   bool
}
