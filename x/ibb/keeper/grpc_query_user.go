package keeper

import (
	"context"

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

	for i := 0; i < assetBalances.Len(); i++ {
		var userAsset types.LoadUserResponse
		userAsset.AssetBalance = assetBalances[i].Amount.Int64()
		userAsset.AssetDenom = assetBalances[i].Denom
		userAsset.AssetPrice = int32(assetPrices[i] * 1000000)
		userAsset.Collateral = queryUser.Collateral[i]
		userAsset.AssetDeposit = queryUser.Deposit[i].Amount
		userAsset.AssetBorrow = queryUser.Borrow[i].Amount

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
