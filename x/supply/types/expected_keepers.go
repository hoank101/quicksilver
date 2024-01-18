package types // noalias

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the contract required for account APIs.
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	IterateAccounts(ctx context.Context, cb func(account sdk.AccountI) (stop bool))
}

// BankKeeper defines the contract needed to be fulfilled for banking and supply
// dependencies.
type BankKeeper interface {
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	GetSupply(ctx context.Context, denom string) sdk.Coin
	LockedCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
}

type StakingKeeper interface {
	BondDenom(ctx context.Context) (string, error)
}
