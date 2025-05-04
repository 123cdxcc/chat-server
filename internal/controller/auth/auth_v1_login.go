package auth

import (
	"context"
	"im-chat/internal/dao"
	"im-chat/internal/model/do"
	"im-chat/internal/model/entity"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"im-chat/api/auth/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	username := req.UserName

	ok, err := dao.User.Ctx(ctx).Where("username = ?", username).Exist()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	u := entity.User{}
	if !ok {
		id, err := dao.User.Ctx(ctx).Data(do.User{
			Username: username,
		}).InsertAndGetId()
		if err != nil {
			return nil, gerror.WrapCode(gcode.CodeInternalError, err)
		}
		u.Id = id
	} else {
		err := dao.User.Ctx(ctx).Where("username", username).Scan(&u)
		if err != nil {
			return nil, gerror.WrapCode(gcode.CodeInternalError, err)
		}
	}
	auth.SetSessionUserID(ctx, u.Id)
	return &v1.LoginRes{}, nil
}
