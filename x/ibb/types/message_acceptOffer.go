package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAcceptOffer{}

func NewMsgAcceptOffer(creator string, nftId int32, offerId int32) *MsgAcceptOffer {
	return &MsgAcceptOffer{
		Creator: creator,
		NftId:   nftId,
		OfferId: offerId,
	}
}

func (msg *MsgAcceptOffer) Route() string {
	return RouterKey
}

func (msg *MsgAcceptOffer) Type() string {
	return "AcceptOffer"
}

func (msg *MsgAcceptOffer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptOffer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
