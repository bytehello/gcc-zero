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
	ccReleaseLogFieldNames          = builderx.RawFieldNames(&CcReleaseLog{})
	ccReleaseLogRows                = strings.Join(ccReleaseLogFieldNames, ",")
	ccReleaseLogRowsExpectAutoSet   = strings.Join(stringx.Remove(ccReleaseLogFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ccReleaseLogRowsWithPlaceHolder = strings.Join(stringx.Remove(ccReleaseLogFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	CcReleaseLogModel interface {
		Insert(data CcReleaseLog) (sql.Result, error)
		FindOne(id int64) (*CcReleaseLog, error)
		Update(data CcReleaseLog) error
		Delete(id int64) error
	}

	defaultCcReleaseLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CcReleaseLog struct {
		Id           int64        `db:"id"`
		Ip           string       `db:"ip"`             // 客户端ip
		KvId         int64        `db:"kv_id"`          // cc_kv id
		AppId        int64        `db:"app_id"`         // app_id
		ClusterId    int64        `db:"cluster_id"`     // cluster_id
		KvRevisionId int64        `db:"kv_revision_id"` // cc_kv_revision id
		ReleaseTime  sql.NullTime `db:"release_time"`   // 下发配置时间
		CreateTime   time.Time    `db:"create_time"`
		UpdateTime   time.Time    `db:"update_time"`
	}
)

func NewCcReleaseLogModel(conn sqlx.SqlConn) CcReleaseLogModel {
	return &defaultCcReleaseLogModel{
		conn:  conn,
		table: "`cc_release_log`",
	}
}

func (m *defaultCcReleaseLogModel) Insert(data CcReleaseLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, ccReleaseLogRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Ip, data.KvId, data.AppId, data.ClusterId, data.KvRevisionId, data.ReleaseTime)
	return ret, err
}

func (m *defaultCcReleaseLogModel) FindOne(id int64) (*CcReleaseLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ccReleaseLogRows, m.table)
	var resp CcReleaseLog
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

func (m *defaultCcReleaseLogModel) Update(data CcReleaseLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ccReleaseLogRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Ip, data.KvId, data.AppId, data.ClusterId, data.KvRevisionId, data.ReleaseTime, data.Id)
	return err
}

func (m *defaultCcReleaseLogModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
