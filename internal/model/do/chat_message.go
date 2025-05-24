// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChatMessage is the golang structure of table chat_message for DAO operations like Where/Data.
type ChatMessage struct {
	g.Meta       `orm:"table:chat_message, do:true"`
	Id           interface{} // 消息ID
	ClientSeqId  interface{} // 客户端序列号
	SenderId     interface{} // 发送者ID(用户ID)
	ReceiverId   interface{} // 接收者ID(用户ID或房间ID, 根据receiver_type确定)
	ReceiverType interface{} // 接收者类型(user/room)
	Content      interface{} // 消息内容
	CreatedAt    *gtime.Time // 创建时间
}
