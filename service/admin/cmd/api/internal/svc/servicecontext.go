package svc

import (
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/config"
	"github.com/bytehello/gcc-zero/service/admin/model"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/ccclient"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	AdminUser   model.AdminUserModel
	CcRpcClient ccclient.Cc
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		AdminUser:   model.NewAdminUserModel(conn),
		CcRpcClient: ccclient.NewCc(zrpc.MustNewClient(c.CcRpc)),
	}
}
