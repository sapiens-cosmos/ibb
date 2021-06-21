package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func TestBorrowMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateBorrow(ctx, &types.MsgCreateBorrow{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestBorrowMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateBorrow
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateBorrow{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBorrow{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBorrow{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateBorrow(ctx, &types.MsgCreateBorrow{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateBorrow(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBorrowMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteBorrow
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteBorrow{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteBorrow{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteBorrow{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateBorrow(ctx, &types.MsgCreateBorrow{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteBorrow(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
