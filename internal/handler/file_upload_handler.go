package handler

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
	"igeek/internal/logic"
	"igeek/internal/svc"
	"igeek/internal/types"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取上传的文件（FormData）
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}

		// md5
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))

		// 往 logic 传递 request
		req.Ext = path.Ext(fileHeader.Filename)
		req.Title = strings.TrimRight(fileHeader.Filename, req.Ext)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Types = strings.ToLower(strings.Trim(req.Ext, "."))
		req.Bits = b

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
