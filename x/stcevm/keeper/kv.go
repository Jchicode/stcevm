package keeper

import (
	"github.com/Jchicode/stcevm/x/stcevm/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetKv set kv in the store
func (k Keeper) SetKv(ctx sdk.Context, kv types.Kv) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvKey))
	b := k.cdc.MustMarshal(&kv)
	store.Set([]byte{0}, b)
}

// GetKv returns kv
func (k Keeper) GetKv(ctx sdk.Context) (val types.Kv, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKv removes kv from the store
func (k Keeper) RemoveKv(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KvKey))
	store.Delete([]byte{0})
}
