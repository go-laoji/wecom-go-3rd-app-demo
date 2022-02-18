package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"wecom-go-3rd-app-demo/internal/logic"
	"wecom-go-3rd-app-demo/internal/svc"
	"wecom-go-3rd-app-demo/internal/types"
)

func suiteCmdCallbackPostHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MsgRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSuiteCmdCallbackPostLogic(r.Context(), ctx)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			httpx.Error(w, err)
		}
		resp, err := l.SuiteCmdCallbackPost(req, body)
		if err != nil {
			httpx.Error(w, err)
		} else {
			w.Write(resp.Data)
			httpx.Ok(w)
		}
	}
}
