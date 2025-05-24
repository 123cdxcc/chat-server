// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatMessage is the golang structure for table chat_message.
type ChatMessage struct {
	Id           string      `json:"id"           orm:"id"            description:"消息ID"`                                // 消息ID
	ClientSeqId  string      `json:"clientSeqId"  orm:"client_seq_id" description:"客户端序列号"`                              // 客户端序列号
	SenderId     int64       `json:"senderId"     orm:"sender_id"     description:"发送者ID(用户ID)"`                         // 发送者ID(用户ID)
	ReceiverId   int64       `json:"receiverId"   orm:"receiver_id"   description:"接收者ID(用户ID或房间ID, 根据receiver_type确定)"` // 接收者ID(用户ID或房间ID, 根据receiver_type确定)
	ReceiverType string      `json:"receiverType" orm:"receiver_type" description:"接收者类型(user/room)"`                    // 接收者类型(user/room)
	Content      string      `json:"content"      orm:"content"       description:"消息内容"`                                // 消息内容
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`                                // 创建时间
}
