package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AprAll(c context.Context, req *types.QueryAllAprRequest) (*types.QueryAllAprResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var aprs []*types.Apr
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	aprStore := prefix.NewStore(store, types.KeyPrefix(types.AprKey))

	pageRes, err := query.Paginate(aprStore, req.Pagination, func(key []byte, value []byte) error {
		var apr types.Apr
		if err := k.cdc.UnmarshalBinaryBare(value, &apr); err != nil {
			return err
		}

		aprs = append(aprs, &apr)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAprResponse{Apr: aprs, Pagination: pageRes}, nil
}

func (k Keeper) Apr(c context.Context, req *types.QueryGetAprRequest) (*types.QueryGetAprResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var apr types.Apr
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasApr(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetAprIDBytes(req.Id)), &apr)

	return &types.QueryGetAprResponse{Apr: &apr}, nil
}
