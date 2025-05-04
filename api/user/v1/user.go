package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"im-chat/internal/model/entity"
)

type GetUserDetailReq struct {
	g.Meta `path:"/user/detail" tags:"用户" method:"get" summary:"用户详情"`
}

type GetUserDetailRes struct {
	*entity.User `dc:"用户信息"`
}
