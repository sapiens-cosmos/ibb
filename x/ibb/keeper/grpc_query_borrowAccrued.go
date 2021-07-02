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

func (k Keeper) BorrowAccruedAll(c context.Context, req *types.QueryAllBorrowAccruedRequest) (*types.QueryAllBorrowAccruedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var borrowAccrueds []*types.BorrowAccrued
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	borrowAccruedStore := prefix.NewStore(store, types.KeyPrefix(types.BorrowAccruedKey))

	pageRes, err := query.Paginate(borrowAccruedStore, req.Pagination, func(key []byte, value []byte) error {
		var borrowAccrued types.BorrowAccrued
		if err := k.cdc.UnmarshalBinaryBare(value, &borrowAccrued); err != nil {
			return err
		}

		borrowAccrueds = append(borrowAccrueds, &borrowAccrued)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBorrowAccruedResponse{BorrowAccrued: borrowAccrueds, Pagination: pageRes}, nil
}

func (k Keeper) BorrowAccrued(c context.Context, req *types.QueryGetBorrowAccruedRequest) (*types.QueryGetBorrowAccruedResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var borrowAccrued types.BorrowAccrued
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasBorrowAccrued(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccruedKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBorrowAccruedIDBytes(req.Id)), &borrowAccrued)

	return &types.QueryGetBorrowAccruedResponse{BorrowAccrued: &borrowAccrued}, nil
}
