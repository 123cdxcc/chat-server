package friend

import (
	"context"
	"im-chat/internal/dao"
	"im-chat/internal/model/do"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "im-chat/api/friend/v1"
)

func (c *ControllerV1) AddFriend(ctx context.Context, req *v1.AddFriendReq) (res *v1.AddFriendRes, err error) {
	res = &v1.AddFriendRes{}
	userID := auth.GetSessionUserID(ctx)

	// 检查是否已经是好友
	exist, err := dao.UserFriendRelation.Ctx(ctx).Where("user_id = ? AND friend_id = ?", userID, req.UserId).Exist()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	if exist {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "已经是好友关系")
	}

	err = dao.UserFriendRelation.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 创建双向好友关系
		_, err = tx.Model(dao.UserFriendRelation.Table()).Insert(do.UserFriendRelation{
			UserId:   userID,
			FriendId: req.UserId,
		})
		if err != nil {
			return gerror.WrapCode(gcode.CodeInternalError, err)
		}

		_, err = tx.Model(dao.UserFriendRelation.Table()).Insert(do.UserFriendRelation{
			UserId:   req.UserId,
			FriendId: userID,
		})
		if err != nil {
			return gerror.WrapCode(gcode.CodeInternalError, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
