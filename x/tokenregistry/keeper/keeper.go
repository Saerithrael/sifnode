package keeper

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gogotypes "github.com/gogo/protobuf/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/Sifchain/sifnode/x/tokenregistry/types"
)

type keeper struct {
	cdc codec.BinaryMarshaler

	storeKey sdk.StoreKey
}

func NewKeeper(cdc codec.Marshaler, storeKey sdk.StoreKey) types.Keeper {
	return keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// Logger returns a module-specific logger.
func (k keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k keeper) SetAdminAccount(ctx sdk.Context, adminAccount sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminAccountStorePrefix
	store.Set(key, k.cdc.MustMarshalBinaryBare(&gogotypes.BytesValue{Value: adminAccount}))
}

func (k keeper) IsAdminAccount(ctx sdk.Context, adminAccount sdk.AccAddress) bool {
	account := k.GetAdminAccount(ctx)
	if account == nil {
		return false
	}
	return bytes.Equal(account, adminAccount)
}

func (k keeper) GetAdminAccount(ctx sdk.Context) (adminAccount sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := types.AdminAccountStorePrefix
	bz := store.Get(key)
	acc := gogotypes.BytesValue{}
	k.cdc.MustUnmarshalBinaryBare(bz, &acc)

	adminAccount = sdk.AccAddress(acc.Value)

	return adminAccount
}

func (k keeper) IsDenomWhitelisted(ctx sdk.Context, denom string) bool {
	d := k.GetDenom(ctx, denom)

	return d.IsWhitelisted
}

func (k keeper) CheckDenomPermissions(ctx sdk.Context, denom string, requiredPermissions []types.Permission) bool {
	d := k.GetDenom(ctx, denom)

	for _, requiredPermission := range requiredPermissions {
		var has bool
		for _, allowedPermission := range d.Permissions {
			if allowedPermission == requiredPermission {
				has = true
				break
			}
		}

		if !has {
			return false
		}
	}

	return true
}

func (k keeper) GetDenom(ctx sdk.Context, denom string) types.RegistryEntry {
	wl := k.GetDenomWhitelist(ctx)

	// TODO: This is inefficient, and potentially an attack vector. Investigate
	// using a map instead
	// See ticket #2065
	for i := range wl.Entries {
		if wl.Entries[i] != nil && strings.EqualFold(wl.Entries[i].Denom, denom) {
			return *wl.Entries[i]
		}
	}

	return types.RegistryEntry{
		IsWhitelisted: false,
		Denom:         denom,
	}
}

func (k keeper) SetToken(ctx sdk.Context, entry *types.RegistryEntry) {
	wl := k.GetDenomWhitelist(ctx)

	entry.Sanitize()

	for i := range wl.Entries {
		if wl.Entries[i] != nil && strings.EqualFold(wl.Entries[i].Denom, entry.Denom) {
			wl.Entries[i] = entry

			k.SetDenomWhitelist(ctx, wl)
			return
		}
	}

	wl.Entries = append(wl.Entries, entry)

	k.SetDenomWhitelist(ctx, wl)
}

func (k keeper) RemoveToken(ctx sdk.Context, denom string) {
	registry := k.GetDenomWhitelist(ctx)

	updated := make([]*types.RegistryEntry, 0)
	for _, t := range registry.Entries {
		if t != nil && !strings.EqualFold(t.Denom, denom) {
			updated = append(updated, t)
		}
	}

	k.SetDenomWhitelist(ctx, types.Registry{
		Entries: updated,
	})
}

func (k keeper) SetDenomWhitelist(ctx sdk.Context, wl types.Registry) {
	store := ctx.KVStore(k.storeKey)

	bz := k.cdc.MustMarshalBinaryBare(&wl)

	store.Set(types.WhitelistStorePrefix, bz)
}

func (k keeper) GetDenomWhitelist(ctx sdk.Context) types.Registry {
	var whitelist types.Registry
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.WhitelistStorePrefix)
	if len(bz) == 0 {
		return types.Registry{}
	}
	k.cdc.MustUnmarshalBinaryBare(bz, &whitelist)
	return whitelist
}

// Exists chec if the key existed in db.
func (k keeper) Exists(ctx sdk.Context, key []byte) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(key)
}
