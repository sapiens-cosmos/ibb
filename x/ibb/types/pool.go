package types

type LoadPoolResponse struct {
	Asset            string
	Id               uint64
	Creator          string
	CollatoralFactor int32
	Liquidity        int32
}
