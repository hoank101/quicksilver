package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BankKeeper defines the expected bank keeper.
type BankKeeper interface {
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}
