package handler

import (
	"net/http"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/logic/cc/app"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AppListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAppListLogic(r.Context(), ctx)
		resp, err := l.AppList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
