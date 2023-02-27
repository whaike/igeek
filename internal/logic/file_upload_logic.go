package logic

import (
	"context"
	"igeek/internal/svc"
	"igeek/internal/types"
	"igeek/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadLogic {
	return &FileUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadLogic) FileUpload(req *types.FileUploadReq) (resp *types.FileUploadResp, err error) {
	resp = &types.FileUploadResp{}
	// 根据名称和hash查重，有重复的直接返回,没有重复则上传之后将信息补充到mysql并返回
	switch req.Types {
	case "mp3":
		row, err := l.svcCtx.AudioModel.Find(req.Title, req.Hash)
		if err != nil {
			resp.Code = 500
			resp.Message = err.Error()
			return resp, nil
		}
		if row != nil {
			resp.Code = 0
			resp.Message = "duplicated file"
			resp.Link = row.Link
			resp.Hash = row.Hash
			resp.Types = req.Types
			resp.Ext = req.Ext
			resp.AudioID = row.Id
			return resp, nil
		}
		location, err := l.svcCtx.AudioS3.Upload(context.Background(), req.Hash, req.Bits)
		if err != nil {
			resp.Code = 500
			resp.Message = err.Error()
			return resp, nil
		}
		res, err := l.svcCtx.AudioModel.Insert(l.ctx, &model.AudioInfo{
			Title: req.Title,
			Hash:  req.Hash,
			Link:  location,
			Size:  req.Size,
			Ext:   req.Ext,
		})
		if err != nil {
			resp.Code = 500
			resp.Message = err.Error()
			return resp, nil
		}
		resp.Code = 0
		resp.Message = "success"
		resp.Link = location
		resp.Hash = req.Hash
		resp.Types = req.Types
		resp.Ext = req.Ext
		resp.AudioID, _ = res.LastInsertId()

	case "pdf", "html", "txt":
		row, err := l.svcCtx.ChapterModel.Find(req.Title, req.Hash)
		if err != nil {
			resp.Code = 500
			resp.Message = err.Error()
			return resp, nil
		}
		if row != nil {
			resp.Code = 0
			resp.Message = "duplicated file"
			resp.Link = row.Link
			resp.Hash = row.Hash
			resp.Types = req.Types
			resp.Ext = req.Ext
			return resp, nil
		}
		location, err := l.svcCtx.ChapterS3.Upload(context.Background(), req.Hash, req.Bits)
		if err != nil {
			resp.Code = 500
			resp.Message = err.Error()
			return resp, nil
		}
		resp.Code = 0
		resp.Message = "success"
		resp.Link = location
		resp.Types = req.Types
		resp.Ext = req.Ext
	}
	return resp, nil
}
