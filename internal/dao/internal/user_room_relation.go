// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserRoomRelationDao is the data access object for the table user_room_relation.
type UserRoomRelationDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of the current DAO.
	columns UserRoomRelationColumns // columns contains all the column names of Table for convenient usage.
}

// UserRoomRelationColumns defines and stores column names for the table user_room_relation.
type UserRoomRelationColumns struct {
	Id         string // 关系ID
	UserId     string // 用户ID
	RoomId     string // 房间ID
	Role       string // 用户在房间中的角色
	Subscribed string // 是否订阅消息
	JoinedAt   string // 加入时间
}

// userRoomRelationColumns holds the columns for the table user_room_relation.
var userRoomRelationColumns = UserRoomRelationColumns{
	Id:         "id",
	UserId:     "user_id",
	RoomId:     "room_id",
	Role:       "role",
	Subscribed: "subscribed",
	JoinedAt:   "joined_at",
}

// NewUserRoomRelationDao creates and returns a new DAO object for table data access.
func NewUserRoomRelationDao() *UserRoomRelationDao {
	return &UserRoomRelationDao{
		group:   "default",
		table:   "user_room_relation",
		columns: userRoomRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserRoomRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserRoomRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserRoomRelationDao) Columns() UserRoomRelationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserRoomRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserRoomRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserRoomRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
