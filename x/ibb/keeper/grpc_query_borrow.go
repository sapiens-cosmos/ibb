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

func (k Keeper) BorrowAll(c context.Context, req *types.QueryAllBorrowRequest) (*types.QueryAllBorrowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var borrows []*types.Borrow
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	borrowStore := prefix.NewStore(store, types.KeyPrefix(types.BorrowKey))

	pageRes, err := query.Paginate(borrowStore, req.Pagination, func(key []byte, value []byte) error {
		var borrow types.Borrow
		if err := k.cdc.UnmarshalBinaryBare(value, &borrow); err != nil {
			return err
		}

		borrows = append(borrows, &borrow)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBorrowResponse{Borrow: borrows, Pagination: pageRes}, nil
}

func (k Keeper) Borrow(c context.Context, req *types.QueryGetBorrowRequest) (*types.QueryGetBorrowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var borrow types.Borrow
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasBorrow(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBorrowIDBytes(req.Id)), &borrow)

	return &types.QueryGetBorrowResponse{Borrow: &borrow}, nil
}
