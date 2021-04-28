package model

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	adminUserFieldNames          = builderx.RawFieldNames(&AdminUser{})
	adminUserRows                = strings.Join(adminUserFieldNames, ",")
	adminUserRowsExpectAutoSet   = strings.Join(stringx.Remove(adminUserFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	adminUserRowsWithPlaceHolder = strings.Join(stringx.Remove(adminUserFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	AdminUserModel interface {
		Insert(data AdminUser) (sql.Result, error)
		FindOne(id int64) (*AdminUser, error)
		Update(data AdminUser) error
		Delete(id int64) error
		FindOneByUsername(userName string) (*AdminUser, error)
	}

	defaultAdminUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	AdminUser struct {
		Password   string    `db:"password"` // 用户密码
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
		Id         int64     `db:"id"`
		UserName   string    `db:"user_name"` // 用户名称
	}
)

func NewAdminUserModel(conn sqlx.SqlConn) AdminUserModel {
	return &defaultAdminUserModel{
		conn:  conn,
		table: "`admin_user`",
	}
}

func (m *defaultAdminUserModel) Insert(data AdminUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, adminUserRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Password, data.UserName)
	return ret, err
}

func (m *defaultAdminUserModel) FindOne(id int64) (*AdminUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminUserRows, m.table)
	var resp AdminUser
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminUserModel) Update(data AdminUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adminUserRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Password, data.UserName, data.Id)
	return err
}

func (m *defaultAdminUserModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultAdminUserModel) FindOneByUsername(userName string) (*AdminUser, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `user_name` = ? limit 1", adminUserRows, m.table)
	var resp AdminUser
	err := m.conn.QueryRow(&resp, query, userName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
