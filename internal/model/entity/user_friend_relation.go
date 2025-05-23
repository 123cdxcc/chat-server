// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFriendRelation is the golang structure for table user_friend_relation.
type UserFriendRelation struct {
	Id        int64       `json:"id"        orm:"id"         description:"关系ID"` // 关系ID
	UserId    int64       `json:"userId"    orm:"user_id"    description:"用户ID"` // 用户ID
	FriendId  int64       `json:"friendId"  orm:"friend_id"  description:"好友ID"` // 好友ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
}
