package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"igeek/internal/logic"
	"igeek/internal/svc"
	"igeek/internal/types"
)

func ListChapterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListChapterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListChapterLogic(r.Context(), svcCtx)
		resp, err := l.ListChapter(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
