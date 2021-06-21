package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func TestUserMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateUser(ctx, &types.MsgCreateUser{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestUserMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateUser
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateUser{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUser{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateUser{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateUser(ctx, &types.MsgCreateUser{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateUser(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteUser
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteUser{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteUser{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteUser{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateUser(ctx, &types.MsgCreateUser{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteUser(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
