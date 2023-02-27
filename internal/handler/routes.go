// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"igeek/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/course",
				Handler: CreateCourseHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/course",
				Handler: ListCourseHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload",
				Handler: FileUploadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/chapter",
				Handler: ListChapterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/chapter",
				Handler: CreateChapterHandler(serverCtx),
			},
		},
	)
}
