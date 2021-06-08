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
	ccKvRevisionFieldNames          = builderx.RawFieldNames(&CcKvRevision{})
	ccKvRevisionRows                = strings.Join(ccKvRevisionFieldNames, ",")
	ccKvRevisionRowsExpectAutoSet   = strings.Join(stringx.Remove(ccKvRevisionFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ccKvRevisionRowsWithPlaceHolder = strings.Join(stringx.Remove(ccKvRevisionFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	CcKvRevisionModel interface {
		Insert(data CcKvRevision) (sql.Result, error)
		FindOne(id int64) (*CcKvRevision, error)
		Update(data CcKvRevision) error
		Delete(id int64) error
	}

	defaultCcKvRevisionModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CcKvRevision struct {
		Value          string       `db:"value"`           // 值
		Version        int64        `db:"version"`         // 版本
		CreateRevision int64        `db:"create_revision"` // create_revision
		CreateTime     time.Time    `db:"create_time"`
		UpdateTime     time.Time    `db:"update_time"`
		DeletedTime    sql.NullTime `db:"deleted_time"`
		Id             int64        `db:"id"`
		KvId           int64        `db:"kv_id"`        // cc_kv主键
		Key            string       `db:"key"`          // 键名
		ModRevision    int64        `db:"mod_revision"` // mod_revision
	}
)

func NewCcKvRevisionModel(conn sqlx.SqlConn) CcKvRevisionModel {
	return &defaultCcKvRevisionModel{
		conn:  conn,
		table: "`cc_kv_revision`",
	}
}

func (m *defaultCcKvRevisionModel) Insert(data CcKvRevision) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, ccKvRevisionRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Value, data.Version, data.CreateRevision, data.DeletedTime, data.KvId, data.Key, data.ModRevision)
	return ret, err
}

func (m *defaultCcKvRevisionModel) FindOne(id int64) (*CcKvRevision, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ccKvRevisionRows, m.table)
	var resp CcKvRevision
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

func (m *defaultCcKvRevisionModel) Update(data CcKvRevision) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ccKvRevisionRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Value, data.Version, data.CreateRevision, data.DeletedTime, data.KvId, data.Key, data.ModRevision, data.Id)
	return err
}

func (m *defaultCcKvRevisionModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
