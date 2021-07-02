package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBorrowAccrued{}

func NewMsgCreateBorrowAccrued(creator string, blockHeight int32, asset string, amount int32, denom string) *MsgCreateBorrowAccrued {
	return &MsgCreateBorrowAccrued{
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgCreateBorrowAccrued) Route() string {
	return RouterKey
}

func (msg *MsgCreateBorrowAccrued) Type() string {
	return "CreateBorrowAccrued"
}

func (msg *MsgCreateBorrowAccrued) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBorrowAccrued) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBorrowAccrued) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBorrowAccrued{}

func NewMsgUpdateBorrowAccrued(creator string, id uint64, blockHeight int32, asset string, amount int32, denom string) *MsgUpdateBorrowAccrued {
	return &MsgUpdateBorrowAccrued{
		Id:          id,
		Creator:     creator,
		BlockHeight: blockHeight,
		Asset:       asset,
		Amount:      amount,
		Denom:       denom,
	}
}

func (msg *MsgUpdateBorrowAccrued) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBorrowAccrued) Type() string {
	return "UpdateBorrowAccrued"
}

func (msg *MsgUpdateBorrowAccrued) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBorrowAccrued) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBorrowAccrued) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateBorrowAccrued{}

func NewMsgDeleteBorrowAccrued(creator string, id uint64) *MsgDeleteBorrowAccrued {
	return &MsgDeleteBorrowAccrued{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteBorrowAccrued) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBorrowAccrued) Type() string {
	return "DeleteBorrowAccrued"
}

func (msg *MsgDeleteBorrowAccrued) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBorrowAccrued) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBorrowAccrued) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
