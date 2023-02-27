package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"igeek/internal/logic"
	"igeek/internal/svc"
	"igeek/internal/types"
)

func ListCourseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCourseReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListCourseLogic(r.Context(), svcCtx)
		resp, err := l.ListCourse(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
