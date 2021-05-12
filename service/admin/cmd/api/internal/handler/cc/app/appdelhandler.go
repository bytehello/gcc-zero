package handler

import (
	"github.com/bytehello/gcc-zero/common/errorx"
	"net/http"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/logic/cc/app"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AppDelHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AppDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.DefaultCodeError(err.Error()))
			return
		}

		l := logic.NewAppDelLogic(r.Context(), ctx)
		resp, err := l.AppDel(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
