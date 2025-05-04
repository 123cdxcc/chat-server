package user

import (
	"context"
	"database/sql"
	"errors"
	"im-chat/internal/dao"
	"im-chat/utility/auth"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"im-chat/api/user/v1"
)

func (c *ControllerV1) GetUserDetail(ctx context.Context, _ *v1.GetUserDetailReq) (res *v1.GetUserDetailRes, err error) {
	userID := auth.GetSessionUserID(ctx)
	res = &v1.GetUserDetailRes{}
	err = dao.User.Ctx(ctx).WherePri(userID).Scan(res)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, gerror.NewCode(gcode.CodeNotFound)
		}
		return nil, gerror.WrapCode(gcode.CodeInternalError, err)
	}
	return res, nil
}
