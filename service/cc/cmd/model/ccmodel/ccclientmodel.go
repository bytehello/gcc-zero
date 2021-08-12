package ccmodel

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	ccClientFieldNames          = builderx.RawFieldNames(&CcClient{})
	ccClientRows                = strings.Join(ccClientFieldNames, ",")
	ccClientRowsExpectAutoSet   = strings.Join(stringx.Remove(ccClientFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	ccClientRowsWithPlaceHolder = strings.Join(stringx.Remove(ccClientFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	CcClientModel interface {
		FindAllByKvId(kvId int64) ([]*CcClient, error)
		Insert(data CcClient) (sql.Result, error)
		FindOne(id int64) (*CcClient, error)
		Update(data CcClient) error
		Delete(id int64) error
	}

	defaultCcClientModel struct {
		conn  sqlx.SqlConn
		gorm  *gorm.DB
		table string
	}

	CcClient struct {
		Id          int64        `db:"id"`
		Ip          string       `db:"ip"`           // 客户端ip
		KvId        int64        `db:"kv_id"`        // cc_kv id
		AppId       int64        `db:"app_id"`       // app_id
		ClusterId   int64        `db:"cluster_id"`   // cluster_id
		ReleaseTime sql.NullTime `db:"release_time"` // 下发配置时间
		VisitedTime sql.NullTime `db:"visited_time"` // 访问机器时间
		CreateTime  time.Time    `db:"create_time"`
		UpdateTime  time.Time    `db:"update_time"`
	}
)

func NewCcClientModel(conn sqlx.SqlConn, gorm *gorm.DB) CcClientModel {
	return &defaultCcClientModel{
		conn:  conn,
		gorm:  gorm,
		table: "`cc_client`",
	}
}
func (CcClient) TableName() string {
	return "cc_client"
}

func (m *defaultCcClientModel) Insert(data CcClient) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, ccClientRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.Ip, data.KvId, data.AppId, data.ClusterId, data.ReleaseTime, data.VisitedTime)
	return ret, err
}

func (m *defaultCcClientModel) FindAllByKvId(kvId int64) ([]*CcClient, error) {
	var resp []*CcClient
	err := m.gorm.Where("`kv_id` = ?", kvId).Find(&resp).Error
	switch err {
	case nil:
		return resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCcClientModel) FindOne(id int64) (*CcClient, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", ccClientRows, m.table)
	var resp CcClient
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

func (m *defaultCcClientModel) Update(data CcClient) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, ccClientRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.Ip, data.KvId, data.AppId, data.ClusterId, data.ReleaseTime, data.VisitedTime, data.Id)
	return err
}

func (m *defaultCcClientModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}
