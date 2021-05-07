package ccmodel

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
	ccAppFieldNames          = builderx.RawFieldNames(&CcApp{})
	ccAppRows                = strings.Join(ccAppFieldNames, ",")
	ccAppRowsExpectAutoSet   = strings.Join(stringx.Remove(ccAppFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ccAppRowsWithPlaceHolder = strings.Join(stringx.Remove(ccAppFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	CcAppModel interface {
		Insert(data CcApp) (sql.Result, error)
		FindAll(current int64, pageSize int64) (*[]CcApp, error)
		FindOne(id int64) (*CcApp, error)
		FindOneByAppKey(appKey string) (*CcApp, error)
		Update(data CcApp) error
		Delete(id int64) error
		Count() (int64, error)
	}

	defaultCcAppModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CcApp struct {
		Desc        string       `db:"desc"` // app描述
		CreateTime  time.Time    `db:"create_time"`
		UpdateTime  time.Time    `db:"update_time"`
		DeletedTime sql.NullTime `db:"deleted_time"`
		Id          int64        `db:"id"`
		AppKey      string       `db:"app_key"`  // app key
		AppName     string       `db:"app_name"` // app名称
	}
)

func NewCcAppModel(conn sqlx.SqlConn) CcAppModel {
	return &defaultCcAppModel{
		conn:  conn,
		table: "`cc_app`",
	}
}

func (m *defaultCcAppModel) Insert(data CcApp) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, ccAppRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Desc, data.DeletedTime, data.AppKey, data.AppName)
	return ret, err
}

func (m *defaultCcAppModel) FindAll(current int64, pageSize int64) (*[]CcApp, error) {
	if current < 1 {
		current = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT %d,%d", ccAppRows, m.table, (current-1)*pageSize, pageSize)
	var resp []CcApp
	err := m.conn.QueryRows(&resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCcAppModel) FindOne(id int64) (*CcApp, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ccAppRows, m.table)
	var resp CcApp
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

func (m *defaultCcAppModel) FindOneByAppKey(appKey string) (*CcApp, error) {
	var resp CcApp
	query := fmt.Sprintf("select %s from %s where `app_key` = ? limit 1", ccAppRows, m.table)
	err := m.conn.QueryRow(&resp, query, appKey)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCcAppModel) Update(data CcApp) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ccAppRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Desc, data.DeletedTime, data.AppKey, data.AppName, data.Id)
	return err
}

func (m *defaultCcAppModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultCcAppModel) Count() (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
