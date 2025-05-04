package room

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"im-chat/internal/dao"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"im-chat/api/room/v1"
)

func (c *ControllerV1) GetRoomList(ctx context.Context, req *v1.GetRoomListReq) (res *v1.GetRoomListRes, err error) {
	res = &v1.GetRoomListRes{}
	userID := auth.GetSessionUserID(ctx)
	col := dao.UserRoomRelation.Columns()
	vals, err := dao.UserRoomRelation.Ctx(ctx).Where("user_id = ?", userID).Fields(col.RoomId).Array()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	roomIDs := gconv.Int64s(vals)
	err = dao.Room.Ctx(ctx).WhereIn("id", roomIDs).Scan(&res.List)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	return res, nil
}
