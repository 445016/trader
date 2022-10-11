package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"trader/internal/logic"
	"trader/internal/svc"
	"trader/internal/types"
)

func TraderHandlerGetHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TradingViewHookRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTraderHandlerGetLogic(r.Context(), svcCtx)
		resp, err := l.TraderHandlerGet(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
