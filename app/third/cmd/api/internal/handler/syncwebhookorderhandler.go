package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jekka-api-go/app/third/cmd/api/internal/logic"
	"jekka-api-go/app/third/cmd/api/internal/svc"
	"jekka-api-go/app/third/cmd/api/internal/types"
)

func syncWebhookOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SyncWebhookOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSyncWebhookOrderLogic(r.Context(), svcCtx)
		resp, err := l.SyncWebhookOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
