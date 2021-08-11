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
	ccKvFieldNames          = builderx.RawFieldNames(&CcKv{})
	ccKvRows                = strings.Join(ccKvFieldNames, ",")
	ccKvRowsExpectAutoSet   = strings.Join(stringx.Remove(ccKvFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ccKvRowsWithPlaceHolder = strings.Join(stringx.Remove(ccKvFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	CcKvModel interface {
		Count(appId int64, clusterId int64) int64
		FindAll(appId int64, clusterId int64, current int64, pageSize int64) *[]CcKv
		Insert(data CcKv) (sql.Result, error)
		FindOne(id int64) (*CcKv, error)
		Update(data CcKv) error
		Delete(id int64) error
		FindOneByAppIdClusterIdKey(appId int64, clusterId int64, key string) (*CcKv, error)
	}

	defaultCcKvModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CcKv struct {
		Id             int64        `db:"id"`
		ClusterId      int64        `db:"cluster_id"`  // cc_cluster id
		PushStatus     int64        `db:"push_status"` // 推送状态0未推送1已经推送
		DeletedTime    sql.NullTime `db:"deleted_time"`
		AppId          int64        `db:"app_id"`  // cc_app id
		Value          string       `db:"value"`   // 值
		Version        int64        `db:"version"` // 版本
		Format         string       `db:"format"`  // value 格式
		CreateTime     time.Time    `db:"create_time"`
		Desc           string       `db:"desc"`         // 配置描述
		ModRevision    int64        `db:"mod_revision"` // mod_revision
		UpdateTime     time.Time    `db:"update_time"`
		PushedTime     sql.NullTime `db:"pushed_time"`     // 推送时间
		Key            string       `db:"key"`             // 键名
		CreateRevision int64        `db:"create_revision"` // create_revision
	}
)

func NewCcKvModel(conn sqlx.SqlConn) CcKvModel {
	return &defaultCcKvModel{
		conn:  conn,
		table: "`cc_kv`",
	}
}

func (m *defaultCcKvModel) Count(appId int64, clusterId int64) int64 {
	query := fmt.Sprintf("SELECT COUNT(*) AS count FROM %s WHERE `app_id` = ? AND `cluster_id` = ?", m.table)
	var count int64
	_ = m.conn.QueryRowPartial(&count, query, appId, clusterId)
	return count
}

func (m *defaultCcKvModel) FindAll(appId int64, clusterId int64, current int64, pageSize int64) *[]CcKv {
	if current < 1 {
		current = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `app_id` = ? AND `cluster_id` = ? LIMIT %d,%d", ccKvRows, m.table, (current-1)*pageSize, pageSize)
	var resp []CcKv
	_ = m.conn.QueryRows(&resp, query, appId, clusterId)
	return &resp
}

func (m *defaultCcKvModel) Insert(data CcKv) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, ccKvRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ClusterId, data.PushStatus, data.DeletedTime, data.AppId, data.Value, data.Version, data.Format, data.Desc, data.ModRevision, data.PushedTime, data.Key, data.CreateRevision)
	return ret, err
}

func (m *defaultCcKvModel) FindOne(id int64) (*CcKv, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ccKvRows, m.table)
	var resp CcKv
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

func (m *defaultCcKvModel) Update(data CcKv) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ccKvRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ClusterId, data.PushStatus, data.DeletedTime, data.AppId, data.Value, data.Version, data.Format, data.Desc, data.ModRevision, data.PushedTime, data.Key, data.CreateRevision, data.Id)
	return err
}

func (m *defaultCcKvModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultCcKvModel) FindOneByAppIdClusterIdKey(appId int64, clusterId int64, key string) (*CcKv, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `app_id` = ? AND `cluster_id` = ? AND `key` = ?", ccKvRows, m.table)
	var resp CcKv
	err := m.conn.QueryRow(&resp, query, appId, clusterId, key)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
