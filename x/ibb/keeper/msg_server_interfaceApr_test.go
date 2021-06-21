package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func TestInterfaceAprMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateInterfaceApr(ctx, &types.MsgCreateInterfaceApr{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestInterfaceAprMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateInterfaceApr
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateInterfaceApr{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateInterfaceApr{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateInterfaceApr{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateInterfaceApr(ctx, &types.MsgCreateInterfaceApr{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateInterfaceApr(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestInterfaceAprMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteInterfaceApr
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteInterfaceApr{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteInterfaceApr{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteInterfaceApr{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateInterfaceApr(ctx, &types.MsgCreateInterfaceApr{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteInterfaceApr(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
