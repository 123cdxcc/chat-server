// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ws

import (
	"context"

	"im-chat/api/ws/v1"
)

type IWsV1 interface {
	ChatChannel(ctx context.Context, req *v1.ChatChannelReq) (res *v1.ChatChannelRes, err error)
}
