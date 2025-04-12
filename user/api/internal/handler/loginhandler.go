package handler

import (
	"img_analyse/user/commn/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"img_analyse/user/api/internal/logic"
	"img_analyse/user/api/internal/svc"
	"img_analyse/user/api/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(r, w, resp, err)
	}
}
