package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func listPool(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	msgs := keeper.GetAllPool(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, msgs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func getPool(ctx sdk.Context, key string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	id, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return nil, err
	}

	if !keeper.HasPool(ctx, id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	msg := keeper.GetPool(ctx, id)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, msg)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func loadPool(ctx sdk.Context, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	msgs := keeper.GetAllPool(ctx)
	var loadPoolList []types.LoadPoolResponse
	var loadPool types.LoadPoolResponse
	for _, msg := range msgs {
		loadPool.Asset = msg.Asset
		loadPool.CollatoralFactor = msg.CollatoralFactor
		//TODO: Change logic in accordance with API Endpoint
		loadPool.Liquidity = msg.DepositBalance

		loadPoolList = append(loadPoolList, loadPool)
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, loadPoolList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
