// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomDao is the data access object for the table room.
type RoomDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of the current DAO.
	columns RoomColumns // columns contains all the column names of Table for convenient usage.
}

// RoomColumns defines and stores column names for the table room.
type RoomColumns struct {
	Id   string // 房间ID
	Name string // 房间名称
}

// roomColumns holds the columns for the table room.
var roomColumns = RoomColumns{
	Id:   "id",
	Name: "name",
}

// NewRoomDao creates and returns a new DAO object for table data access.
func NewRoomDao() *RoomDao {
	return &RoomDao{
		group:   "default",
		table:   "room",
		columns: roomColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RoomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RoomDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RoomDao) Columns() RoomColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RoomDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RoomDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *RoomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
