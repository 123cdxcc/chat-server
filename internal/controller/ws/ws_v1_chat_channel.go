package ws

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"im-chat/internal/dao"
	"im-chat/internal/model/entity"
	"im-chat/utility"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"im-chat/api/ws/v1"
)

func (c *ControllerV1) ChatChannel(ctx context.Context, _ *v1.ChatChannelReq) (res *v1.ChatChannelRes, err error) {
	userID := auth.GetSessionUserID(ctx)
	if userID == 0 {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized)
	}
	user := new(entity.User)
	err = dao.User.Ctx(ctx).WherePri(userID).Scan(user)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取用户信息失败")
	}
	request := g.RequestFromCtx(ctx)
	ws, err := c.upgrader.Upgrade(request.Response.Writer, request.Request, nil)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "升级websocket失败")
	}
	defer ws.Close()
	c.chatManager.AddChannel(ctx, userID, ws)
	defer c.chatManager.RemoveChannel(ctx, userID, ws)
	roomCol := dao.UserRoomRelation.Columns()
	for {
		data := new(v1.ChatData)
		err := ws.ReadJSON(data)
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseMessage, websocket.CloseNoStatusReceived, websocket.CloseNormalClosure) {
				break
			}
			glog.Warning(ctx, err)
			break
		}
		if data.To == nil {
			continue
		}
		data.From = user
		data.ID = utility.NewID()
		vals, err := dao.UserRoomRelation.Ctx(ctx).Where("room_id = ?", data.To.Id).Fields([]string{roomCol.UserId}).Array()
		if err != nil {
			return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取房间用户失败")
		}
		userIDs := gconv.Int64s(vals)
		err = c.chatManager.SendUsersMessage(ctx, userIDs, data)
		if err != nil {
			glog.Warningf(ctx, "消息[%s]发送失败", data.ID)
		}
	}
	return new(v1.ChatChannelRes), nil
}
