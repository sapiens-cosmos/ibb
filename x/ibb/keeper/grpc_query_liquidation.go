package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sapiens-cosmos/ibb/x/ibb/oracle"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LiquidationList(c context.Context, req *types.QueryLiquidationListRequest) (*types.QueryLiquidationListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var liquidations []*types.Liquidation

	ctx := sdk.UnwrapSDKContext(c)
	assetPrices := oracle.GetAllPrices()

	userList := k.GetAllUser(ctx)
	for _, user := range userList {
		var liquidation types.Liquidation
		var liquidationAssetList []*types.LiquidationAsset
		valueDeposit := float64(0)
		valueBorrow := float64(0)

		//using enums
		for i, deposit := range user.Deposit {
			valueDeposit = valueDeposit + float64(deposit.Amount)*assetPrices[i]
		}
		for i, borrow := range user.Borrow {
			valueBorrow = valueBorrow + float64(borrow.Amount)*assetPrices[i]
		}

		if (valueBorrow / valueDeposit) > types.LiquidationRatio {
			for _, deposit := range user.Deposit {
				var liquidationAsset types.LiquidationAsset
				liquidationAsset.Amount = int64(deposit.Amount)
				liquidationAsset.Asset = deposit.Asset

				liquidationAssetList = append(liquidationAssetList, &liquidationAsset)
			}
			liquidation.LiquidationAsset = liquidationAssetList
			liquidation.WalletAddress = user.Creator
			liquidations = append(liquidations, &liquidation)
		}
	}

	return &types.QueryLiquidationListResponse{Liquidation: liquidations}, nil
}
