package controller

import (
	"im-chat/internal/controller/auth"
	"im-chat/internal/controller/room"
	"im-chat/internal/controller/user"
	"im-chat/internal/controller/ws"

	"github.com/google/wire"
)

var AuthProviderSet = wire.NewSet(
	user.NewV1,
	room.NewV1,
	ws.NewV1,
)

var NoAuthProviderSet = wire.NewSet(
	auth.NewV1,
)
