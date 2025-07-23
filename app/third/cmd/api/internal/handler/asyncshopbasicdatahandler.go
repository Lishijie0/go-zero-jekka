package handler

import (
	"jekka-api-go/pkg/response/xresp"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"jekka-api-go/app/third/cmd/api/internal/logic"
	"jekka-api-go/app/third/cmd/api/internal/svc"
	"jekka-api-go/app/third/cmd/api/internal/types"
)

func asyncShopBasicDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AsyncShopBasicDataReq
		if err := httpx.Parse(r, &req); err != nil {
			xresp.ParamErrorResult(r, w, err)
			return
		}

		l := logic.NewAsyncShopBasicDataLogic(r.Context(), svcCtx)
		resp, err := l.AsyncShopBasicData(&req)
		xresp.HttpResult(r, w, resp, err)
	}
}
