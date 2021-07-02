package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTxHistory{}

func NewMsgCreateTxHistory(creator string, blockHeight int32, tx string, asset string, amount int32, denom string) *MsgCreateTxHistory {
	return &MsgCreateTxHistory{
		Creator:     creator,
		BlockHeight: blockHeight,
		Tx:          tx,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgCreateTxHistory) Route() string {
	return RouterKey
}

func (msg *MsgCreateTxHistory) Type() string {
	return "CreateTxHistory"
}

func (msg *MsgCreateTxHistory) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTxHistory) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTxHistory) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTxHistory{}

func NewMsgUpdateTxHistory(creator string, id uint64, blockHeight int32, tx string, asset string, amount int32, denom string) *MsgUpdateTxHistory {
	return &MsgUpdateTxHistory{
		Id:          id,
		Creator:     creator,
		BlockHeight: blockHeight,
		Tx:          tx,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgUpdateTxHistory) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTxHistory) Type() string {
	return "UpdateTxHistory"
}

func (msg *MsgUpdateTxHistory) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateTxHistory) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTxHistory) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateTxHistory{}

func NewMsgDeleteTxHistory(creator string, id uint64) *MsgDeleteTxHistory {
	return &MsgDeleteTxHistory{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteTxHistory) Route() string {
	return RouterKey
}

func (msg *MsgDeleteTxHistory) Type() string {
	return "DeleteTxHistory"
}

func (msg *MsgDeleteTxHistory) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteTxHistory) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteTxHistory) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
