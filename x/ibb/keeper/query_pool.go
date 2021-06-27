package keeper

import (
	"math"
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
	poolList := keeper.GetAllPool(ctx)
	var loadPoolList []types.LoadPoolRestResponse
	var loadPool types.LoadPoolRestResponse
	for _, msg := range poolList {
		currentTargetBorrowRatio := float64(msg.BorrowBalance) / float64(msg.DepositBalance)
		currentDepositApy := types.DepositInterest + types.DepositInterest*math.Abs(currentTargetBorrowRatio-float64(types.TargetBorrowRatio)*0.01)*types.InterestFactor
		loadPool.Asset = msg.Asset
		loadPool.CollatoralFactor = msg.CollatoralFactor
		loadPool.Liquidity = msg.DepositBalance - msg.BorrowBalance
		loadPool.DepositApy = int32(currentDepositApy * 1000000)
		loadPool.BorrowApy = int32(currentDepositApy / currentTargetBorrowRatio * 1000000)

		loadPoolList = append(loadPoolList, loadPool)
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, loadPoolList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
