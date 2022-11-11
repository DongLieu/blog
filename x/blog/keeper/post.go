package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

//	func (k Keeper) AppendPost() uint64 {
//	  count := k.GetPostCount()
//	  store.Set()
//	  k.SetPostCount()
//	  return count
//	}
func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	// Get the current number of posts in the store
	count := k.GetPostCount(ctx)

	// Assign an ID to the post based on the number of posts in the store
	post.Id = count

	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))

	// Convert the post ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, post.Id)

	// Marshal the post into bytes
	appendedValue := k.cdc.MustMarshal(&post)

	// Insert the post bytes using post ID as a key
	store.Set(byteKey, appendedValue)

	// Update the post count
	k.SetPostCount(ctx, count+1)
	return count
}
