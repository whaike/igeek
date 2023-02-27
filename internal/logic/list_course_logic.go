package logic

import (
	"context"

	"igeek/internal/svc"
	"igeek/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCourseLogic {
	return &ListCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCourseLogic) ListCourse(req *types.ListCourseReq) (resp *types.ApiResp, err error) {
	resp = &types.ApiResp{}
	res, err := l.svcCtx.CourseModel.FindAll(l.ctx)
	if err != nil {
		resp.Message = "Failed"
		resp.Code = 500
		l.Logger.Errorf("查询错误: %s", err.Error())
	} else {
		resp.Message = "Success"
		resp.Data = res
	}
	return resp, nil
}
