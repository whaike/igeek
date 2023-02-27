package logic

import (
	"context"

	"igeek/internal/svc"
	"igeek/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListChapterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListChapterLogic {
	return &ListChapterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListChapterLogic) ListChapter(req *types.ListChapterReq) (resp *types.ApiResp, err error) {
	resp = &types.ApiResp{}
	rows, err := l.svcCtx.ChapterModel.FindAll(req.CourseID)
	if err != nil {
		resp.Code = 500
		resp.Message = err.Error()
		return resp, nil
	}
	resp.Message = "success"
	resp.Data = rows

	return
}
