package handler

import (
	"net/http"

	"college/api/internal/logic"
	"college/api/internal/svc"
	"college/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CollegeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollegeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUniversityLogic(r.Context(), ctx)
		resp, err := l.University(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
