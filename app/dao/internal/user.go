// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

// UserDao is the manager for logic model data accessing and custom defined data operations functions management.
type UserDao struct {
	Table   string      // Table is the underlying table name of the DAO.
	Group   string      // Group is the database configuration group name of current DAO.
	Columns UserColumns // Columns is the short type for Columns, which contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	Id       string // 用户id（主键）
	Username string // 用户名称
	RealName string // 用户真实名称
	RoleId   string // 用户角色，1表示普通用户
	Password string // 用户密码
	Phone    string // 用户电话
	Balance  string // 用户余额
	Status   string // 用户状态，1表示正常，0表示暂停
	Created  string // 创建时间
	Updated  string // 更新时间
}

//  userColumns holds the columns for table user.
var userColumns = UserColumns{
	Id:       "id",
	Username: "username",
	RealName: "real_name",
	RoleId:   "role_id",
	Password: "password",
	Phone:    "phone",
	Balance:  "balance",
	Status:   "status",
	Created:  "created",
	Updated:  "updated",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		Group:   "default",
		Table:   "user",
		Columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.Group)
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.Table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
