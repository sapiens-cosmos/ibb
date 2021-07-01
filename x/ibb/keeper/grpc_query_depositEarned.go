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

func (k Keeper) DepositEarnedAll(c context.Context, req *types.QueryAllDepositEarnedRequest) (*types.QueryAllDepositEarnedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var depositEarneds []*types.DepositEarned
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	depositEarnedStore := prefix.NewStore(store, types.KeyPrefix(types.DepositEarnedKey))

	pageRes, err := query.Paginate(depositEarnedStore, req.Pagination, func(key []byte, value []byte) error {
		var depositEarned types.DepositEarned
		if err := k.cdc.UnmarshalBinaryBare(value, &depositEarned); err != nil {
			return err
		}

		depositEarneds = append(depositEarneds, &depositEarned)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDepositEarnedResponse{DepositEarned: depositEarneds, Pagination: pageRes}, nil
}

func (k Keeper) DepositEarned(c context.Context, req *types.QueryGetDepositEarnedRequest) (*types.QueryGetDepositEarnedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var depositEarned types.DepositEarned
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasDepositEarned(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetDepositEarnedIDBytes(req.Id)), &depositEarned)

	return &types.QueryGetDepositEarnedResponse{DepositEarned: &depositEarned}, nil
}
