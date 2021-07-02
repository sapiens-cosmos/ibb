package keeper

import (
	"context"
	"math"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sapiens-cosmos/ibb/x/ibb/oracle"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) UserAll(c context.Context, req *types.QueryAllUserRequest) (*types.QueryAllUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var users []*types.User
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	userStore := prefix.NewStore(store, types.KeyPrefix(types.UserKey))

	pageRes, err := query.Paginate(userStore, req.Pagination, func(key []byte, value []byte) error {
		var user types.User
		if err := k.cdc.UnmarshalBinaryBare(value, &user); err != nil {
			return err
		}

		users = append(users, &user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserResponse{User: users, Pagination: pageRes}, nil
}

func (k Keeper) User(c context.Context, req *types.QueryGetUserRequest) (*types.QueryGetUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var user types.User
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasUser(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetUserIDBytes(req.Id)), &user)

	return &types.QueryGetUserResponse{User: &user}, nil
}

func (k Keeper) UserLoad(c context.Context, req *types.QueryLoadUserRequest) (*types.QueryLoadUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// var user types.User
	ctx := sdk.UnwrapSDKContext(c)

	userList := k.GetAllUser(ctx)

	var queryUser types.User
	for _, user := range userList {
		if user.Creator == req.Id {
			queryUser = user
		}
	}
	//TODO: keynotfound error when user has not been found
	var userAssetList []*types.LoadUserResponse

	userAddress, err := sdk.AccAddressFromBech32(req.Id)
	if err != nil {
		return nil, err
	}

	assetBalances := k.bankKeeper.GetAllBalances(ctx, userAddress)[1:]
	assetPrices := oracle.GetAllPrices()
	poolList := k.GetAllPool(ctx)

	for i, msg := range poolList {
		var userAsset types.LoadUserResponse
		userAsset.AssetBalance = assetBalances[i].Amount.Int64()
		userAsset.AssetDenom = assetBalances[i].Denom
		userAsset.AssetPrice = int32(assetPrices[i] * 1000000)
		userAsset.Collateral = queryUser.Collateral[i]
		userAsset.AssetDeposit = queryUser.Deposit[i].Amount
		userAsset.AssetBorrow = queryUser.Borrow[i].Amount

		currentTargetBorrowRatio := float64(msg.BorrowBalance) / float64(msg.DepositBalance)
		currentDepositApy := types.DepositInterest + types.DepositInterest*(currentTargetBorrowRatio-float64(types.TargetBorrowRatio)*0.01)*types.InterestFactor
		currentDepositApy = math.Max(currentDepositApy, types.MinimumDepositInterest)
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

		userAssetList = append(userAssetList, &userAsset)
	}

	return &types.QueryLoadUserResponse{LoadUserResponse: userAssetList}, nil
}

func (k Keeper) UserBalance(c context.Context, req *types.QueryBalanceUserRequest) (*types.QueryBalanceUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	//TODO: omit keynotfound error when user has not been found
	userAddress, err := sdk.AccAddressFromBech32(req.Id)
	if err != nil {
		return nil, err
	}

	assetBalances := k.bankKeeper.GetAllBalances(ctx, userAddress)[1:]

	var userBalance types.UserAssetBalances
	userBalance.Uakt = assetBalances[0].Amount.Int64()
	userBalance.Uatom = assetBalances[1].Amount.Int64()
	userBalance.Ucro = assetBalances[2].Amount.Int64()
	userBalance.Udvpn = assetBalances[3].Amount.Int64()
	userBalance.Uiris = assetBalances[4].Amount.Int64()
	userBalance.Uxprt = assetBalances[5].Amount.Int64()

	return &types.QueryBalanceUserResponse{UserAssetBalances: &userBalance}, nil
}
