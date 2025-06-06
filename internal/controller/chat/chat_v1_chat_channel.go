package chat

import (
	"context"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gorilla/websocket"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "im-chat/api/chat/v1"
)

func (c *ControllerV1) ChatChannel(ctx context.Context, _ *v1.ChatChannelReq) (res *v1.ChatChannelRes, err error) {
	userID := auth.GetSessionUserID(ctx)
	if userID == 0 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized)
	}
	request := g.RequestFromCtx(ctx)
	ws, err := c.upgrader.Upgrade(request.Response.Writer, request.Request, nil)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "升级websocket失败")
	}
	defer ws.Close()
	c.chatManager.AddChannel(ctx, userID, ws)
	defer c.chatManager.RemoveChannel(ctx, userID, ws)
	for {
		data := new(v1.Message)
		err := ws.ReadJSON(data)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseMessage, websocket.CloseNoStatusReceived, websocket.CloseNormalClosure) {
				break
			}
			glog.Warning(ctx, err)
			break
		}
		c.chatManager.HandleMessage(ctx, data)
	}
	return new(v1.ChatChannelRes), nil
}
