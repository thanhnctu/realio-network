package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateToken = "update_token"

var _ sdk.Msg = &MsgUpdateToken{}

func NewMsgUpdateToken(creator string, symbol string, authorizationRequired bool) *MsgUpdateToken {
	return &MsgUpdateToken{
		Creator:               creator,
		Symbol:                symbol,
		AuthorizationRequired: authorizationRequired,
	}
}

func (msg *MsgUpdateToken) Route() string {
	return RouterKey
}

func (msg *MsgUpdateToken) Type() string {
	return TypeMsgUpdateToken
}

func (msg *MsgUpdateToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
