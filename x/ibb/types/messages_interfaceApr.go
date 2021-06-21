package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateInterfaceApr{}

func NewMsgCreateInterfaceApr(creator string, blockTime string, timeApr int32) *MsgCreateInterfaceApr {
	return &MsgCreateInterfaceApr{
		Creator:   creator,
		BlockTime: blockTime,
		TimeApr:   timeApr,
	}
}

func (msg *MsgCreateInterfaceApr) Route() string {
	return RouterKey
}

func (msg *MsgCreateInterfaceApr) Type() string {
	return "CreateInterfaceApr"
}

func (msg *MsgCreateInterfaceApr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateInterfaceApr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateInterfaceApr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateInterfaceApr{}

func NewMsgUpdateInterfaceApr(creator string, id uint64, blockTime string, timeApr int32) *MsgUpdateInterfaceApr {
	return &MsgUpdateInterfaceApr{
		Id:        id,
		Creator:   creator,
		BlockTime: blockTime,
		TimeApr:   timeApr,
	}
}

func (msg *MsgUpdateInterfaceApr) Route() string {
	return RouterKey
}

func (msg *MsgUpdateInterfaceApr) Type() string {
	return "UpdateInterfaceApr"
}

func (msg *MsgUpdateInterfaceApr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateInterfaceApr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateInterfaceApr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteInterfaceApr{}

func NewMsgDeleteInterfaceApr(creator string, id uint64) *MsgDeleteInterfaceApr {
	return &MsgDeleteInterfaceApr{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteInterfaceApr) Route() string {
	return RouterKey
}

func (msg *MsgDeleteInterfaceApr) Type() string {
	return "DeleteInterfaceApr"
}

func (msg *MsgDeleteInterfaceApr) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteInterfaceApr) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteInterfaceApr) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
