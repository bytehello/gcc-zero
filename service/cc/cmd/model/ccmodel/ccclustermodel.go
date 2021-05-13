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
		Update(data CcCluster) error
		Delete(id int64) error
	}

	defaultCcClusterModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CcCluster struct {
		Id          int64        `db:"id"`
		ClusterName string       `db:"cluster_name"` // 集群名
		Desc        string       `db:"desc"`         // a描述
		CreateTime  time.Time    `db:"create_time"`
		UpdateTime  time.Time    `db:"update_time"`
		DeletedTime sql.NullTime `db:"deleted_time"`
	}
)

func NewCcClusterModel(conn sqlx.SqlConn) CcClusterModel {
	return &defaultCcClusterModel{
		conn:  conn,
		table: "`cc_cluster`",
	}
}

func (m *defaultCcClusterModel) Insert(data CcCluster) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, ccClusterRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.ClusterName, data.Desc, data.DeletedTime)
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

func (m *defaultCcClusterModel) Update(data CcCluster) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ccClusterRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.ClusterName, data.Desc, data.DeletedTime, data.Id)
	return err
}

func (m *defaultCcClusterModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
