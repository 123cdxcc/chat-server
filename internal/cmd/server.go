package cmd

import (
	"im-chat/api/auth"
	"im-chat/api/chat"
	"im-chat/api/friend"
	"im-chat/api/room"
	"im-chat/api/user"
	authUtil "im-chat/utility/auth"

	"github.com/gogf/gf/v2/net/ghttp"
)

type NoAuthServer func(group *ghttp.RouterGroup)

func NewNoAuthServer(authV1 auth.IAuthV1) NoAuthServer {
	return func(group *ghttp.RouterGroup) {
		group.Bind(authV1)
	}
}

type AuthServer func(group *ghttp.RouterGroup)

func NewAuthServer(userV1 user.IUserV1, chatV1 chat.IChatV1, roomV1 room.IRoomV1, friendV1 friend.IFriendV1) AuthServer {
	return func(group *ghttp.RouterGroup) {
		group.Middleware(authUtil.SessionAuth)
		group.Bind(
			userV1,
			chatV1,
			roomV1,
			friendV1,
		)
	}
}

type App struct {
	NoAuthServer NoAuthServer
	AuthServer   AuthServer
}

func NewApp(noAuth NoAuthServer, auth AuthServer) *App {
	return &App{
		NoAuthServer: noAuth,
		AuthServer:   auth,
	}
}
