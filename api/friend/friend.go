// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package friend

import (
	"context"

	"im-chat/api/friend/v1"
)

type IFriendV1 interface {
	AddFriend(ctx context.Context, req *v1.AddFriendReq) (res *v1.AddFriendRes, err error)
	GetFriendList(ctx context.Context, req *v1.GetFriendListReq) (res *v1.GetFriendListRes, err error)
}
