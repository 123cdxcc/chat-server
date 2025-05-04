// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Room is the golang structure of table room for DAO operations like Where/Data.
type Room struct {
	g.Meta `orm:"table:room, do:true"`
	Id     interface{} // 房间ID
	Name   interface{} // 房间名称
}
