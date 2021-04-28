package svc

import (
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/config"
	"github.com/bytehello/gcc-zero/service/admin/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	AdminUser model.AdminUserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		AdminUser: model.NewAdminUserModel(conn),
	}
}
