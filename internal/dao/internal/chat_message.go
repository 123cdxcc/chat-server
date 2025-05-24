// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChatMessageDao is the data access object for the table chat_message.
type ChatMessageDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of the current DAO.
	columns ChatMessageColumns // columns contains all the column names of Table for convenient usage.
}

// ChatMessageColumns defines and stores column names for the table chat_message.
type ChatMessageColumns struct {
	Id           string // 消息ID
	ClientSeqId  string // 客户端序列号
	SenderId     string // 发送者ID(用户ID)
	ReceiverId   string // 接收者ID(用户ID或房间ID, 根据receiver_type确定)
	ReceiverType string // 接收者类型(user/room)
	Content      string // 消息内容
	CreatedAt    string // 创建时间
}

// chatMessageColumns holds the columns for the table chat_message.
var chatMessageColumns = ChatMessageColumns{
	Id:           "id",
	ClientSeqId:  "client_seq_id",
	SenderId:     "sender_id",
	ReceiverId:   "receiver_id",
	ReceiverType: "receiver_type",
	Content:      "content",
	CreatedAt:    "created_at",
}

// NewChatMessageDao creates and returns a new DAO object for table data access.
func NewChatMessageDao() *ChatMessageDao {
	return &ChatMessageDao{
		group:   "default",
		table:   "chat_message",
		columns: chatMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ChatMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ChatMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ChatMessageDao) Columns() ChatMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ChatMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ChatMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ChatMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
