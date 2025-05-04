package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/auth/login" tags:"授权" method:"post" summary:"登陆"`
	UserName string `v:"required" dc:"用户名"`
}

type LoginRes struct {
}
