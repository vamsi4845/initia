package keeper

import (
	"context"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)

// Unjail calls the staking Unjail function to unjail a validator if the
// jailed period has concluded
func (k Keeper) Unjail(ctx context.Context, validatorAddr sdk.ValAddress) error {
	validator, err := k.sk.Validator(ctx, validatorAddr)
	if err != nil {
		return err
	}
	if validator == nil {
		return types.ErrNoValidatorForAddress
	}

	// cannot be unjailed if no self-delegation exists
	selfDel, err := k.sk.Delegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	if err != nil {
		return err
	}

	if selfDel == nil {
		return types.ErrMissingSelfDelegation
	}

	tokens, _ := validator.TokensFromShares(selfDel.GetShares()).TruncateDecimal()
	vp, err := k.sk.VotingPower(ctx, tokens)
	if err != nil {
		return err
	}

	cp := k.sk.VotingPowerToConsensusPower(ctx, vp)
	if cp < 1 {
		return errors.Wrapf(
			types.ErrSelfDelegationTooLowToUnjail, "%s less than 1", tokens,
		)
	}

	// cannot be unjailed if not jailed
	if !validator.IsJailed() {
		return types.ErrValidatorNotJailed
	}

	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return err
	}
	// If the validator has a ValidatorSigningInfo object that signals that the
	// validator was bonded and so we must check that the validator is not tombstoned
	// and can be unjailed at the current block.
	//
	// A validator that is jailed but has no ValidatorSigningInfo object signals
	// that the validator was never bonded and must've been jailed due to falling
	// below their minimum self-delegation. The validator can unjail at any point
	// assuming they've now bonded above their minimum self-delegation.
	info, err := k.GetValidatorSigningInfo(ctx, consAddr)
	if err == nil {
		// cannot be unjailed if tombstoned
		if info.Tombstoned {
			return types.ErrValidatorJailed
		}

		// cannot be unjailed until out of jail
		sdkCtx := sdk.UnwrapSDKContext(ctx)
		if sdkCtx.BlockHeader().Time.Before(info.JailedUntil) {
			return types.ErrValidatorJailed
		}
	}

	return k.sk.Unjail(ctx, consAddr)
}
