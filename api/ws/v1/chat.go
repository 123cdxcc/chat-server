package v1

import (
	"encoding/json"
	"im-chat/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type MessageType int

const (
	MessageTypeChatData MessageType = iota
	MessageTypeHeartbeat
	MessageTypeStreamChatData
)

type Message struct {
	From *entity.User `json:"-"`
	Type MessageType  `json:"type"`
	Data any          `json:"data"`
}

// json反序列化时将data单独处理
func (m *Message) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type MessageType     `json:"type"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Type = temp.Type
	switch temp.Type {
	case MessageTypeChatData:
		var chatData ChatData
		if err := json.Unmarshal(temp.Data, &chatData); err != nil {
			return err
		}
		m.Data = chatData
	case MessageTypeHeartbeat:
		m.Data = nil
	}
	return nil
}

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
