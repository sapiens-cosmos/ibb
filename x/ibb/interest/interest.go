package interest

import (
	"math"

	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func GetInterests(asset string, blockHeight int32, txHistories []*types.TxHistory, aprs []*types.Apr) (int32, int32) {
	var checkpoint int32 = 0
	var currentBlockHeight = checkpoint
	var totalDeposit int32 = 0
	var totalDeposits []int32
	var totalBorrow int32 = 0
	var totalBorrows []int32
	for _, txHistory := range txHistories {
		if txHistory.Asset != asset {
			continue
		}

		if txHistory.BlockHeight < currentBlockHeight {
			switch txHistory.Tx {
			case "deposit":
				totalDeposit += txHistory.Amount
			case "withdraw":
				totalDeposit -= txHistory.Amount
			case "borrow":
				totalBorrow += txHistory.Amount
			case "repay":
				totalBorrow -= txHistory.Amount
			default:
				// Do nothing.
			}
		} else {
			for currentBlockHeight <= txHistory.BlockHeight {
				totalDeposits = append(totalDeposits, totalDeposit)
				totalBorrows = append(totalBorrows, totalBorrow)

				currentBlockHeight++
			}
			switch txHistory.Tx {
			case "deposit":
				totalDeposit += txHistory.Amount
			case "withdraw":
				totalDeposit -= txHistory.Amount
			case "borrow":
				totalBorrow += txHistory.Amount
			case "repay":
				totalBorrow -= txHistory.Amount
			default:
				// Do nothing.
			}
		}
	}
	for currentBlockHeight <= blockHeight {
		totalDeposits = append(totalDeposits, totalDeposit)
		totalBorrows = append(totalBorrows, totalBorrow)

		currentBlockHeight++
	}

	currentBlockHeight = checkpoint
	var depositApy int32 = 0
	var depositApys []int32
	var borrowApy int32 = 0
	var borrowApys []int32
	for _, apr := range aprs {
		if apr.BlockHeight < currentBlockHeight {
			depositApy = apr.DepositApy
			borrowApy = apr.BorrowApy
		} else {
			for currentBlockHeight <= apr.BlockHeight {
				depositApys = append(depositApys, depositApy)
				borrowApys = append(borrowApys, borrowApy)

				currentBlockHeight++
			}
			depositApy = apr.DepositApy
			borrowApy = apr.BorrowApy
			depositApys = append(depositApys, depositApy)
			borrowApys = append(borrowApys, borrowApy)
		}
	}
	for currentBlockHeight <= blockHeight {
		depositApys = append(depositApys, depositApy)
		borrowApys = append(borrowApys, borrowApy)

		currentBlockHeight++
	}

	var depositEarnedAmount float64 = 0
	var borrowAccruedAmount float64 = 0
	for j := 0; j < int(math.Min(float64(len(totalDeposits)), float64(len(depositApys)))); j++ {
		// On assumption that 1 block is 1 second, 31536000 blocks are 1 year.
		// Then, 1 block is 1 / 31536000 == 3.1709792e-8 year.
		depositEarnedAmount += float64(totalDeposits[j]) * float64(depositApys[j]) * 1e-6 * 3.1709792e-8
		borrowAccruedAmount += float64(totalBorrows[j]) * float64(borrowApys[j]) * 1e-6 * 3.1709792e-8
	}

	for _, txHistory := range txHistories {
		if txHistory.Asset != asset {
			continue
		}

		if txHistory.Tx == "claim" {
			depositEarnedAmount -= float64(txHistory.Amount)
		}
	}

	return int32(depositEarnedAmount), int32(borrowAccruedAmount)
}
