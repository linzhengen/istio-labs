package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"helloworld/api/internal/logic"
	"helloworld/api/internal/svc"
	"helloworld/api/internal/types"
)

func GetHelloWorldHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetHelloWorldLogic(r.Context(), svcCtx)
		resp, err := l.GetHelloWorld(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
