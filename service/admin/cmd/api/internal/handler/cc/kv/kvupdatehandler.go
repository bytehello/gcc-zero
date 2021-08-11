package handler

import (
	"net/http"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/logic/cc/kv"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func KvUpdateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.KvUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewKvUpdateLogic(r.Context(), ctx)
		resp, err := l.KvUpdate(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
