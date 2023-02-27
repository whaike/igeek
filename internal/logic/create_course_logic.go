package logic

import (
	"context"
	"igeek/model"

	"igeek/internal/svc"
	"igeek/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCourseLogic {
	return &CreateCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCourseLogic) CreateCourse(req *types.CreateCourseReq) (resp *types.ApiResp, err error) {
	resp = &types.ApiResp{}
	err = l.svcCtx.CourseModel.CreateOne(l.ctx, &model.CourseInfo{
		Title:  req.Title,
		Author: req.Author,
		Source: req.Source,
	})
	if err != nil {
		l.Logger.Errorf("create course failed, %s", err.Error())
		resp.Code = 500
		resp.Message = "Failed"
	} else {
		resp.Message = "Success"
	}
	return
}
