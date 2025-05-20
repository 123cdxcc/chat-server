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
)

func (m MessageType) String() string {
	return []string{"chat_data", "heartbeat"}[m]
}

type Message struct {
	From        *entity.User `json:"-"`
	MessageType MessageType  `json:"message_type"`
	Data        any          `json:"data"`
}

// json反序列化时将data单独处理
func (m *Message) UnmarshalJSON(data []byte) error {
	var temp struct {
		MessageType MessageType     `json:"message_type"`
		Data        json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.MessageType = temp.MessageType
	switch temp.MessageType {
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
