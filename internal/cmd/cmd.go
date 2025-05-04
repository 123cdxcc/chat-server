package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"im-chat/internal/controller/auth"
	"im-chat/internal/controller/hello"
	"im-chat/internal/controller/room"
	"im-chat/internal/controller/user"
	"im-chat/internal/controller/ws"
	authUtil "im-chat/utility/auth"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareCORS, ghttp.MiddlewareHandlerResponse)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Bind(auth.NewV1())
				})
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(authUtil.SessionAuth)
					group.Bind(
						hello.NewV1(),
						user.NewV1(),
						room.NewV1(),
						ws.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
