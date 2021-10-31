package handler

import (
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
	"suite-zero-svr/internal/logic"
	"suite-zero-svr/internal/svc"
	"suite-zero-svr/internal/types"
)

func suiteCmdCallbackGetHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MsgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSuiteCmdCallbackGetLogic(r.Context(), ctx)
		resp, err := l.SuiteCmdCallbackGet(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			w.Write(resp.Data)
			httpx.Ok(w)
		}
	}
}
