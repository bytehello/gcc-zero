package svc

import (
	"github.com/bytehello/gcc-zero/service/agent/biz/releasekv"
	"github.com/bytehello/gcc-zero/service/agent/cmd/rpc/internal/config"
)

type ServiceContext struct {
	c          config.Config
	ReleaseSvr *releasekv.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:          c,
		ReleaseSvr: releasekv.NewService(c.Cc.ReleasePath),
	}
}
