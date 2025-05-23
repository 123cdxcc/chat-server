package friend

import (
	"context"
	"im-chat/internal/dao"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "im-chat/api/friend/v1"
)

func (c *ControllerV1) GetFriendList(ctx context.Context, req *v1.GetFriendListReq) (res *v1.GetFriendListRes, err error) {
	res = &v1.GetFriendListRes{}
	userID := auth.GetSessionUserID(ctx)

	// 获取好友ID列表
	col := dao.UserFriendRelation.Columns()
	vals, err := dao.UserFriendRelation.Ctx(ctx).Where("user_id = ?", userID).Fields(col.FriendId).Array()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}

	// 获取好友详细信息
	friendIDs := gconv.Int64s(vals)
	err = dao.User.Ctx(ctx).WhereIn("id", friendIDs).Scan(&res.List)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}

	return res, nil
}
