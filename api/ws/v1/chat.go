package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"im-chat/internal/model/entity"
)

type ChatData struct {
	ID      string       `json:"id"`
	From    *entity.User `json:"from"`
	To      *entity.Room `json:"to"`
	Content string       `json:"content"`
}

type ChatChannelReq struct {
	g.Meta `path:"/chat/channel" tags:"WebSocket" method:"get" summary:"聊天消息传输通道"`
}

type ChatChannelRes struct {
}
