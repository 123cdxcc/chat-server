// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRoomRelation is the golang structure of table user_room_relation for DAO operations like Where/Data.
type UserRoomRelation struct {
	g.Meta     `orm:"table:user_room_relation, do:true"`
	Id         interface{} // 关系ID
	UserId     interface{} // 用户ID
	RoomId     interface{} // 房间ID
	Role       interface{} // 用户在房间中的角色
	Subscribed interface{} // 是否订阅消息
	JoinedAt   *gtime.Time // 加入时间
}
