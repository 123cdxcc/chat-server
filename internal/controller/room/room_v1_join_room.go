package room

import (
	"context"
	"database/sql"
	"errors"
	"im-chat/internal/dao"
	"im-chat/internal/model/do"
	"im-chat/internal/model/entity"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"im-chat/api/room/v1"
)

func (c *ControllerV1) JoinRoom(ctx context.Context, req *v1.JoinRoomReq) (res *v1.JoinRoomRes, err error) {
	room := new(entity.Room)
	err = dao.Room.Ctx(ctx).WherePri(req.RoomId).Scan(room)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, gerror.NewCode(gcode.CodeNotFound)
		}
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	userID := auth.GetSessionUserID(ctx)
	exist, err := dao.UserRoomRelation.Ctx(ctx).Where("room_id = ? and user_id = ?", req.RoomId, userID).Exist()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	if !exist {
		_, err := dao.UserRoomRelation.Ctx(ctx).Data(do.UserRoomRelation{
			RoomId:     req.RoomId,
			UserId:     userID,
			Role:       "member",
			Subscribed: true,
		}).Insert()
		if err != nil {
			return nil, gerror.WrapCode(gcode.CodeInternalError, err)
		}
	}
	return &v1.JoinRoomRes{Room: room}, nil
}
