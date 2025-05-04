// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRoomRelation is the golang structure for table user_room_relation.
type UserRoomRelation struct {
	Id         int64       `json:"id"         orm:"id"         description:"关系ID"`      // 关系ID
	UserId     int64       `json:"userId"     orm:"user_id"    description:"用户ID"`      // 用户ID
	RoomId     int64       `json:"roomId"     orm:"room_id"    description:"房间ID"`      // 房间ID
	Role       string      `json:"role"       orm:"role"       description:"用户在房间中的角色"` // 用户在房间中的角色
	Subscribed int         `json:"subscribed" orm:"subscribed" description:"是否订阅消息"`    // 是否订阅消息
	JoinedAt   *gtime.Time `json:"joinedAt"   orm:"joined_at"  description:"加入时间"`      // 加入时间
}
