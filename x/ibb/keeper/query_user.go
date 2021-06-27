package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/oracle"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func loadUser(ctx sdk.Context, walletAddress string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	//query user
	userList := keeper.GetAllUser(ctx)
	var queryUser types.User
	for _, user := range userList {
		if user.Creator == walletAddress {
			queryUser = user
		}
	}

	//TODO : error when walletAddress is not a valid address

	//TODO: add append user logic when the user has not been registered before
	//the creator var for user should be set as walletAddress variable we received from parameter
	if queryUser.Creator == "" {
		queryUser = types.User{}
		queryUser.Creator = walletAddress
		queryUser.Collateral = []bool{false, false, false, false, false, false}
		queryUser.Deposit = []*types.Deposit{}
		queryUser.Borrow = []*types.Borrow{}
		keeper.AppendUser(ctx, queryUser)
	}

	var userAssetList []types.LoadUserRestResponse
	var userAsset types.LoadUserRestResponse

	userAddress, err := sdk.AccAddressFromBech32(walletAddress)
	if err != nil {
		return nil, err
	}
	assetBalances := keeper.bankKeeper.GetAllBalances(ctx, userAddress)
	assetBalances = assetBalances[1:]
	assetPrices := oracle.GetAllPrices()

	for i := 0; i < assetBalances.Len(); i++ {
		userAsset.AssetBalance = assetBalances[i].Amount.Int64()
		userAsset.AssetDenom = assetBalances[i].Denom
		userAsset.AssetPrice = int32(assetPrices[i] * 1000000)
		userAsset.Collateral = queryUser.Collateral[i]
		// userAsset.AssetBorrow = queryUser.Deposit[0]

		userAssetList = append(userAssetList, userAsset)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, userAssetList)
	// bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, queryUser)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
