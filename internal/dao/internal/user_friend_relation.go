// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserFriendRelationDao is the data access object for the table user_friend_relation.
type UserFriendRelationDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of the current DAO.
	columns UserFriendRelationColumns // columns contains all the column names of Table for convenient usage.
}

// UserFriendRelationColumns defines and stores column names for the table user_friend_relation.
type UserFriendRelationColumns struct {
	Id        string // 关系ID
	UserId    string // 用户ID
	FriendId  string // 好友ID
	CreatedAt string // 创建时间
}

// userFriendRelationColumns holds the columns for the table user_friend_relation.
var userFriendRelationColumns = UserFriendRelationColumns{
	Id:        "id",
	UserId:    "user_id",
	FriendId:  "friend_id",
	CreatedAt: "created_at",
}

// NewUserFriendRelationDao creates and returns a new DAO object for table data access.
func NewUserFriendRelationDao() *UserFriendRelationDao {
	return &UserFriendRelationDao{
		group:   "default",
		table:   "user_friend_relation",
		columns: userFriendRelationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserFriendRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserFriendRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserFriendRelationDao) Columns() UserFriendRelationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserFriendRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserFriendRelationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *UserFriendRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
