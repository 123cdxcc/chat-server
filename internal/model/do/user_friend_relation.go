// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFriendRelation is the golang structure of table user_friend_relation for DAO operations like Where/Data.
type UserFriendRelation struct {
	g.Meta    `orm:"table:user_friend_relation, do:true"`
	Id        interface{} // 关系ID
	UserId    interface{} // 用户ID
	FriendId  interface{} // 好友ID
	CreatedAt *gtime.Time // 创建时间
}
