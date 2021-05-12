package main

import (
	"flag"
	"fmt"
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/config"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/handler"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	httpx.SetErrorHandler(func(err error) (i int, i2 interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			logx.Info("StatusInternalServerError err:", err)
			return http.StatusInternalServerError, nil
		}
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
