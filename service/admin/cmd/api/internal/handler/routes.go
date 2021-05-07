// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	ccapp "github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/handler/cc/app"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/cc/app/add",
				Handler: ccapp.AppAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cc/app/list",
				Handler: ccapp.AppListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cc/app/update",
				Handler: ccapp.AppUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/cc/app/del",
				Handler: ccapp.AppDelHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
