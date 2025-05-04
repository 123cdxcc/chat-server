// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Room is the golang structure for table room.
type Room struct {
	Id   int64  `json:"id"   orm:"id"   description:"房间ID"` // 房间ID
	Name string `json:"name" orm:"name" description:"房间名称"` // 房间名称
}
