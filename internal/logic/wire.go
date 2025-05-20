package logic

import (
	"im-chat/internal/logic/chat"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	chat.NewChannelManager,
)
