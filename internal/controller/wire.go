package controller

import (
	"im-chat/internal/controller/auth"
	"im-chat/internal/controller/chat"
	"im-chat/internal/controller/friend"
	"im-chat/internal/controller/room"
	"im-chat/internal/controller/user"

	"github.com/google/wire"
)

var AuthProviderSet = wire.NewSet(
	user.NewV1,
	room.NewV1,
	chat.NewV1,
	friend.NewV1,
)

var NoAuthProviderSet = wire.NewSet(
	auth.NewV1,
)
