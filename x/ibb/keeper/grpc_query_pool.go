package keeper

import (
	"context"
	"math"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PoolAll(c context.Context, req *types.QueryAllPoolRequest) (*types.QueryAllPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pools []*types.Pool
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	poolStore := prefix.NewStore(store, types.KeyPrefix(types.PoolKey))

	pageRes, err := query.Paginate(poolStore, req.Pagination, func(key []byte, value []byte) error {
		var pool types.Pool
		if err := k.cdc.UnmarshalBinaryBare(value, &pool); err != nil {
			return err
		}

		pools = append(pools, &pool)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPoolResponse{Pool: pools, Pagination: pageRes}, nil
}

func (k Keeper) Pool(c context.Context, req *types.QueryGetPoolRequest) (*types.QueryGetPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pool types.Pool
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasPool(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetPoolIDBytes(req.Id)), &pool)

	return &types.QueryGetPoolResponse{Pool: &pool}, nil
}

func (k Keeper) PoolLoad(c context.Context, req *types.QueryLoadPoolRequest) (*types.QueryLoadPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// var pool types.Pool
	ctx := sdk.UnwrapSDKContext(c)

	poolList := k.GetAllPool(ctx)
	var loadPoolList []*types.LoadPoolResponse
	var loadPool types.LoadPoolResponse
	for _, msg := range poolList {
		currentTargetBorrowRatio := float64(msg.BorrowBalance) / float64(msg.DepositBalance)
		currentDepositApy := types.DepositInterest + types.DepositInterest*math.Abs(currentTargetBorrowRatio-float64(types.TargetBorrowRatio)*0.01)*types.InterestFactor
		loadPool.Asset = msg.Asset
		loadPool.CollatoralFactor = msg.CollatoralFactor
		loadPool.Liquidity = msg.DepositBalance - msg.BorrowBalance
		loadPool.DepositApy = int32(currentDepositApy * 1000000)
		loadPool.BorrowApy = int32(currentDepositApy / currentTargetBorrowRatio * 1000000)

		loadPoolList = append(loadPoolList, &loadPool)
	}
	// k.cdc.MustUnmarshalBinaryBare(store.Get(GetPoolIDBytes(req.Id)), &pool)

	return &types.QueryLoadPoolResponse{LoadPoolResponse: loadPoolList}, nil
}
