package svc

import (
	"github.com/bytehello/gcc-zero/common/mysql"
	"github.com/bytehello/gcc-zero/internal"
	"github.com/bytehello/gcc-zero/service/cc/cmd/model/ccmodel"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/config"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	AppModel     ccmodel.CcAppModel
	ClusterModel ccmodel.CcClusterModel
	KvModel      ccmodel.CcKvModel
	KVer         internal.KVerInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	gorm := mysql.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		AppModel:     ccmodel.NewCcAppModel(sqlConn),
		ClusterModel: ccmodel.NewCcClusterModel(sqlConn),
		KvModel:      ccmodel.NewCcKvModel(sqlConn, gorm),
		KVer:         internal.NewKVer(c.Etcd.Hosts),
	}
}
