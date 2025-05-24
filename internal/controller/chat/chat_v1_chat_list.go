package chat

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"im-chat/api/chat/v1"
)

func (c *ControllerV1) ChatList(ctx context.Context, req *v1.ChatListReq) (res *v1.ChatListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
