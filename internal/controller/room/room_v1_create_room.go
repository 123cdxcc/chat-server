package room

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"im-chat/internal/dao"
	"im-chat/internal/model/do"
	"im-chat/internal/model/entity"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"im-chat/api/room/v1"
)

func (c *ControllerV1) CreateRoom(ctx context.Context, req *v1.CreateRoomReq) (res *v1.CreateRoomRes, err error) {
	var roomID int64
	err = dao.Room.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := tx.Model(dao.Room.Table()).Data(do.Room{
			Name: req.Name,
		}).InsertAndGetId()
		if err != nil {
			return gerror.WrapCode(gcode.CodeInternalError, err)
		}
		_, err = tx.Model(dao.UserRoomRelation.Table()).Data(do.UserRoomRelation{
			RoomId: id,
			UserId: auth.GetSessionUserID(ctx),
			Role:   "created",
		}).InsertIgnore()
		if err != nil {
			return gerror.WrapCode(gcode.CodeInternalError, err)
		}
		roomID = id
		return nil
	})
	if err != nil {
		return nil, err
	}
	room := entity.Room{}
	err = dao.Room.Ctx(ctx).WherePri(roomID).Scan(&room)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	return &v1.CreateRoomRes{
		Room: &room,
	}, nil
}
