package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateApr{}

func NewMsgCreateApr(creator string, blockHeight int32, depositApy int32, borrowApy int32) *MsgCreateApr {
	return &MsgCreateApr{
		Creator:     creator,
		BlockHeight: blockHeight,
		DepositApy:  depositApy,
		BorrowApy:   borrowApy,
	}
}

func (msg *MsgCreateApr) Route() string {
	return RouterKey
}

func (msg *MsgCreateApr) Type() string {
	return "CreateApr"
}

func (msg *MsgCreateApr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateApr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateApr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateApr{}

func NewMsgUpdateApr(creator string, id uint64, blockHeight int32, depositApy int32, borrowApy int32) *MsgUpdateApr {
	return &MsgUpdateApr{
		Id:          id,
		Creator:     creator,
		BlockHeight: blockHeight,
		DepositApy:  depositApy,
		BorrowApy:   borrowApy,
	}
}

func (msg *MsgUpdateApr) Route() string {
	return RouterKey
}

func (msg *MsgUpdateApr) Type() string {
	return "UpdateApr"
}

func (msg *MsgUpdateApr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateApr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateApr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateApr{}

func NewMsgDeleteApr(creator string, id uint64) *MsgDeleteApr {
	return &MsgDeleteApr{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteApr) Route() string {
	return RouterKey
}

func (msg *MsgDeleteApr) Type() string {
	return "DeleteApr"
}

func (msg *MsgDeleteApr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteApr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteApr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
