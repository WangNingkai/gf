// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// User2Dao is the data access object for the table user2.
type User2Dao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of the current DAO.
	columns User2Columns // columns contains all the column names of Table for convenient usage.
}

// User2Columns defines and stores column names for the table user2.
type User2Columns struct {
	Id       string // User ID
	Passport string // User Passport
	Password string // User Password
	Nickname string // User Nickname
	Score    string // Total score amount.
	CreateAt string // Created Time
	UpdateAt string // Updated Time
}

// user2Columns holds the columns for the table user2.
var user2Columns = User2Columns{
	Id:       "id",
	Passport: "passport",
	Password: "password",
	Nickname: "nickname",
	Score:    "score",
	CreateAt: "create_at",
	UpdateAt: "update_at",
}

// NewUser2Dao creates and returns a new DAO object for table data access.
func NewUser2Dao() *User2Dao {
	return &User2Dao{
		group:   "sys",
		table:   "user2",
		columns: user2Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *User2Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *User2Dao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *User2Dao) Columns() User2Columns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *User2Dao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *User2Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *User2Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
