package keeper

import (
	"math"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/oracle"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func loadUser(ctx sdk.Context, walletAddress string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	//query user
	userList := keeper.GetAllUser(ctx)
	var queryUser types.User
	for _, user := range userList {
		if user.Creator == walletAddress {
			queryUser = user
		}
	}

	//TODO : error when walletAddress is not a valid address

	//TODO: add append user logic when the user has not been registered before
	//the creator var for user should be set as walletAddress variable we received from parameter
	// id := uint64(3)
	if queryUser.Creator == "" {
		initialAktDeposit := types.Deposit{Asset: "akt", Denom: "uakt"}
		initialAtomDeposit := types.Deposit{Asset: "atom", Denom: "uatom"}
		initialCroDeposit := types.Deposit{Asset: "cro", Denom: "ucro"}
		initialDvpnDeposit := types.Deposit{Asset: "dvpn", Denom: "udvpn"}
		initialIrisDeposit := types.Deposit{Asset: "iris", Denom: "uiris"}
		initialXprtDeposit := types.Deposit{Asset: "xprt", Denom: "uxprt"}
		initialAktBorrow := types.Borrow{Asset: "akt", Denom: "uakt"}
		initialAtomBorrow := types.Borrow{Asset: "atom", Denom: "uatom"}
		initialCroBorrow := types.Borrow{Asset: "cro", Denom: "ucro"}
		initialDvpnBorrow := types.Borrow{Asset: "dvpn", Denom: "udvpn"}
		initialIrisBorrow := types.Borrow{Asset: "iris", Denom: "uiris"}
		initialXprtBorrow := types.Borrow{Asset: "xprt", Denom: "uxprt"}
		initialAktDepositEarned := types.DepositEarned{Asset: "AKT", Denom: "uakt"}
		initialAtomDepositEarned := types.DepositEarned{Asset: "ATOM", Denom: "uatom"}
		initialCroDepositEarned := types.DepositEarned{Asset: "CRO", Denom: "ucro"}
		initialDvpnDepositEarned := types.DepositEarned{Asset: "IRIS", Denom: "uiris"}
		initialIrisDepositEarned := types.DepositEarned{Asset: "IRIS", Denom: "uiris"}
		initialXprtDepositEarned := types.DepositEarned{Asset: "XPRT", Denom: "uxprt"}
		initialAktBorrowEarned := types.BorrowAccrued{Asset: "AKT", Denom: "uakt"}
		initialAtomBorrowEarned := types.BorrowAccrued{Asset: "ATOM", Denom: "uatom"}
		initialCroBorrowEarned := types.BorrowAccrued{Asset: "CRO", Denom: "ucro"}
		initialDvpnBorrowEarned := types.BorrowAccrued{Asset: "IRIS", Denom: "uiris"}
		initialIrisBorrowEarned := types.BorrowAccrued{Asset: "IRIS", Denom: "uiris"}
		initialXprtBorrowEarned := types.BorrowAccrued{Asset: "XPRT", Denom: "uxprt"}

		argsCollateral := []bool{false, false, false, false, false, false}
		argsDeposit := []*types.Deposit{&initialAktDeposit, &initialAtomDeposit, &initialCroDeposit, &initialDvpnDeposit, &initialIrisDeposit, &initialXprtDeposit}
		argsBorrow := []*types.Borrow{&initialAktBorrow, &initialAtomBorrow, &initialCroBorrow, &initialDvpnBorrow, &initialIrisBorrow, &initialXprtBorrow}
		argsAssetBalances := []int32{2000, 1000, 1000, 1000, 1000, 1000, 1000, 1000}
		argsDepositEarneds := []*types.DepositEarned{&initialAktDepositEarned, &initialAtomDepositEarned, &initialCroDepositEarned, &initialDvpnDepositEarned, &initialIrisDepositEarned, &initialXprtDepositEarned}
		argsBorrowAccrueds := []*types.BorrowAccrued{&initialAktBorrowEarned, &initialAtomBorrowEarned, &initialCroBorrowEarned, &initialDvpnBorrowEarned, &initialIrisBorrowEarned, &initialXprtBorrowEarned}

		queryUser = types.User{}
		queryUser.Creator = walletAddress
		queryUser.Collateral = argsCollateral
		queryUser.Deposit = argsDeposit
		queryUser.Borrow = argsBorrow
		queryUser.AssetBalances = argsAssetBalances
		queryUser.DepositEarneds = argsDepositEarneds
		queryUser.BorrowAccrueds = argsBorrowAccrueds

		// id = keeper.AppendUser(ctx, queryUser)

		return nil, sdkerrors.ErrKeyNotFound
	}

	var userAssetList []types.LoadUserRestResponse
	var userAsset types.LoadUserRestResponse

	userAddress, err := sdk.AccAddressFromBech32(walletAddress)
	if err != nil {
		return nil, err
	}
	assetBalances := keeper.bankKeeper.GetAllBalances(ctx, userAddress)[1:]
	assetPrices := oracle.GetAllPrices()
	poolList := keeper.GetAllPool(ctx)

	for i, msg := range poolList {
		userAsset.AssetBalance = assetBalances[i].Amount.Int64()
		userAsset.AssetDenom = assetBalances[i].Denom
		userAsset.AssetPrice = int32(assetPrices[i] * 1000000)
		userAsset.Collateral = queryUser.Collateral[i]
		userAsset.AssetDeposit = queryUser.Deposit[i].Amount
		userAsset.AssetBorrow = queryUser.Borrow[i].Amount

		currentTargetBorrowRatio := float64(msg.BorrowBalance) / float64(msg.DepositBalance)
		currentDepositApy := types.DepositInterest + types.DepositInterest*(currentTargetBorrowRatio-float64(types.TargetBorrowRatio)*0.01)*types.InterestFactor
		userAsset.Asset = msg.Asset
		userAsset.CollatoralFactor = msg.CollatoralFactor
		userAsset.Liquidity = msg.DepositBalance - msg.BorrowBalance
		userAsset.DepositApy = int32(currentDepositApy * 1000000)
		userAsset.BorrowApy = int32(currentDepositApy / currentTargetBorrowRatio * 1000000)
		userAsset.AssetPrice = int32(assetPrices[i] * 1000000)

		checkpoint := queryUser.DepositEarneds[i].BlockHeight
		var currentBlockHeight = checkpoint
		var totalDeposit int32 = 0
		var totalDeposits []int32
		var totalBorrow int32 = 0
		var totalBorrows []int32
		for _, txHistory := range queryUser.TxHistories {
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
		for currentBlockHeight <= int32(ctx.BlockHeight()) {
			totalDeposits = append(totalDeposits, totalDeposit)
			totalBorrows = append(totalBorrows, totalBorrow)

			currentBlockHeight++
		}

		currentBlockHeight = checkpoint
		var depositApy int32 = 0
		var depositApys []int32
		var borrowApy int32 = 0
		var borrowApys []int32
		for _, apr := range msg.Aprs {
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
		for currentBlockHeight <= int32(ctx.BlockHeight()) {
			depositApys = append(depositApys, depositApy)
			borrowApys = append(borrowApys, borrowApy)

			currentBlockHeight++
		}

		var depositEarnedAmount float64 = 0
		var borrowAccruedAmount float64 = 0
		for j := 0; j < int(math.Min(float64(len(totalDeposits)), float64(len(depositApys)))); j++ {
			// On assumption that 1 block is 1 miniute, 525600 blocks are 1 year.
			// Then, 1 block is 1 / 525600 == 0.00000190258 year.
			depositEarnedAmount += float64(totalDeposits[j]) * float64(depositApys[j]) * 0.00000190258
			borrowAccruedAmount += float64(totalBorrows[j]) * float64(borrowApys[j]) * 0.00000190258
		}

		userAsset.DepositEarned = int32(depositEarnedAmount)
		userAsset.BorrowAccrued = int32(borrowAccruedAmount)

		userAssetList = append(userAssetList, userAsset)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, userAssetList)
	// bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, id)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
