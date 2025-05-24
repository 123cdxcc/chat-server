// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chat

import (
	"context"

	"im-chat/api/chat/v1"
)

type IChatV1 interface {
	ChatChannel(ctx context.Context, req *v1.ChatChannelReq) (res *v1.ChatChannelRes, err error)
	ChatList(ctx context.Context, req *v1.ChatListReq) (res *v1.ChatListRes, err error)
}
