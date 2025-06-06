// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package chat

import (
	apiChat "im-chat/api/chat"
	"im-chat/internal/logic/chat"
	"net/http"

	"github.com/gorilla/websocket"
)

type ControllerV1 struct {
	upgrader    *websocket.Upgrader
	chatManager *chat.ChannelManager
}

func NewV1(channelManager *chat.ChannelManager) apiChat.IChatV1 {
	return &ControllerV1{
		upgrader: &websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		chatManager: channelManager,
	}
}
