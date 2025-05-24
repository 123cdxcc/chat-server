package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MessageType int

const (
	MessageTypeChatData MessageType = iota
	MessageTypeHeartbeat
	MessageTypeStreamChatData
)

type ChatObjectType string

const (
	ChatObjectTypeUser ChatObjectType = "user"
	ChatObjectTypeRoom ChatObjectType = "room"
)

type Message struct {
	Type MessageType `json:"type"` // 消息类型
	Data any         `json:"data"` // 消息数据
}

type Sender struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Receiver struct {
	ID   int64          `json:"id"`
	Name string         `json:"name"`
	Type ChatObjectType `json:"type"`
}

type ChatDataInput struct {
	ClientSeqID string    `json:"client_seq_id"` // 客户端序列号
	SenderID    int64     `json:"sender_id"`     // 发送者ID(用户ID)
	Receiver    *Receiver `json:"receiver"`      // 接收者(用户或房间)
	Content     string    `json:"content"`       // 消息内容
}

type ChatDataOutput struct {
	ID             string    `json:"id"`            // 服务端生成的消息ID
	AckClientSeqID string    `json:"client_seq_id"` // 客户端序列号
	Sender         *Sender   `json:"sender"`        // 发送者(用户)
	Receiver       *Receiver `json:"receiver"`      // 接收者(用户或房间)
	Content        string    `json:"content"`       // 消息内容
}

type ChatChannelReq struct {
	g.Meta `path:"/chat/channel" tags:"Chat" summary:"聊天消息传输通道"`
}

type ChatChannelRes struct {
}

// 聊天列表item对象
type ChatListItem struct {
	ID              string         `json:"id"`
	Type            ChatObjectType `json:"type"`
	Name            string         `json:"name"`
	LastMessage     string         `json:"last_message"`
	LastMessageTime string         `json:"last_message_time"`
	UnreadCount     int            `json:"unread_count"`
}

// 聊天列表
type ChatListReq struct {
	g.Meta `path:"/chat" tags:"Chat" method:"get" summary:"聊天列表"`
}

type ChatListRes struct {
	List []*ChatListItem `json:"list" dc:"聊天列表"`
}
