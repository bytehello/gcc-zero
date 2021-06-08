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
	ccClusterFieldNames          = builderx.RawFieldNames(&CcCluster{})
	ccClusterRows                = strings.Join(ccClusterFieldNames, ",")
	ccClusterRowsExpectAutoSet   = strings.Join(stringx.Remove(ccClusterFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ccClusterRowsWithPlaceHolder = strings.Join(stringx.Remove(ccClusterFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	CcClusterModel interface {
		Insert(data CcCluster) (sql.Result, error)
		FindOne(id int64) (*CcCluster, error)
		FindAll(current int64, pageSize int64) (*[]CcCluster, error)
		Update(data CcCluster) error
		Delete(id int64) error
		Count() (int64, error)
	}

	defaultCcClusterModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CcCluster struct {
		ClusterName string       `db:"cluster_name"` // 集群名
		Desc        string       `db:"desc"`         // a描述
		CreateTime  time.Time    `db:"create_time"`
		UpdateTime  time.Time    `db:"update_time"`
		DeletedTime sql.NullTime `db:"deleted_time"`
		Id          int64        `db:"id"`
		AppId       int64        `db:"app_id"` // cc_app id
	}
)

func NewCcClusterModel(conn sqlx.SqlConn) CcClusterModel {
	return &defaultCcClusterModel{
		conn:  conn,
		table: "`cc_cluster`",
	}
}

func (m *defaultCcClusterModel) Insert(data CcCluster) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, ccClusterRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ClusterName, data.Desc, data.DeletedTime, data.AppId)
	return ret, err
}

func (m *defaultCcClusterModel) FindOne(id int64) (*CcCluster, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ccClusterRows, m.table)
	var resp CcCluster
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
func (m *defaultCcClusterModel) FindAll(current int64, pageSize int64) (*[]CcCluster, error) {
	if current < 1 {
		current = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	query := fmt.Sprintf("SELECT %s FROM %s LIMIT %d,%d", ccClusterRows, m.table, (current-1)*pageSize, pageSize)
	var resp []CcCluster
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

func (m *defaultCcClusterModel) Update(data CcCluster) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ccClusterRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ClusterName, data.Desc, data.DeletedTime, data.AppId, data.Id)
	return err
}

func (m *defaultCcClusterModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultCcClusterModel) Count() (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) AS total FROM %s", m.table)
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
