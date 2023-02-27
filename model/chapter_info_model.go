package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ChapterInfoModel = (*customChapterInfoModel)(nil)

type (
	// ChapterInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChapterInfoModel.
	ChapterInfoModel interface {
		chapterInfoModel
	}

	customChapterInfoModel struct {
		*defaultChapterInfoModel
	}
)

// NewChapterInfoModel returns a model for the database table.
func NewChapterInfoModel(conn sqlx.SqlConn) ChapterInfoModel {
	return &customChapterInfoModel{
		defaultChapterInfoModel: newChapterInfoModel(conn),
	}
}
