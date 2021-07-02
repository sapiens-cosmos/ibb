package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateOffer{}

func NewMsgCreateOffer(creator string, denom string, amount int32, paybackAmount int32, paybackDuration int32, offerStartAt int64, nftId uint64) *MsgCreateOffer {
	return &MsgCreateOffer{
		Creator:         creator,
		Denom:           denom,
		Amount:          amount,
		PaybackAmount:   paybackAmount,
		PaybackDuration: paybackDuration,
		OfferStartAt:    offerStartAt,
		NftId:           nftId,
	}
}

func (msg *MsgCreateOffer) Route() string {
	return RouterKey
}

func (msg *MsgCreateOffer) Type() string {
	return "CreateOffer"
}

func (msg *MsgCreateOffer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOffer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOffer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
