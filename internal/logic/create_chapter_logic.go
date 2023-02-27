package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"igeek/internal/svc"
	"igeek/internal/types"
	"igeek/model"
	"igeek/pkg"
)

type CreateChapterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChapterLogic {
	return &CreateChapterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateChapterLogic) CreateChapter(req *types.CreateChapterReq) (resp *types.ApiResp, err error) {
	resp = &types.ApiResp{}
	row, err := l.svcCtx.ChapterModel.Find(req.Title, req.Hash)
	if err != nil || row != nil {
		resp.Code = 500
		resp.Message = err.Error()
		return resp, nil
	}
	_, err = l.svcCtx.ChapterModel.Insert(l.ctx, &model.ChapterInfo{
		CourseId: req.CourseID,
		AudioId:  req.AudioID,
		Title:    req.Title,
		Hash:     req.Hash,
		Link:     req.Link,
		Ext:      req.Ext,
		Content:  req.Content,
		CreateAt: pkg.CommonTime(),
	})
	if err != nil {
		resp.Code = 500
		resp.Message = "Failed"
	}
	resp.Code = 0
	resp.Message = "success"

	return
}
