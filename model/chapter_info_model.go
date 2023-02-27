package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChapterInfoModel = (*customChapterInfoModel)(nil)

type (
	// ChapterInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChapterInfoModel.
	ChapterInfoModel interface {
		chapterInfoModel
		Find(title, hash string) (*ChapterInfo, error)
		FindAll(courseID int64) ([]*ChapterInfo, error)
	}

	customChapterInfoModel struct {
		*defaultChapterInfoModel
	}
)

func (c customChapterInfoModel) FindAll(courseID int64) ([]*ChapterInfo, error) {
	query := fmt.Sprintf("select %s from %s where `course_id` = ?", chapterInfoRows, c.table)
	resp := make([]*ChapterInfo, 0)
	err := c.conn.QueryRows(&resp, query, courseID)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customChapterInfoModel) Find(title, hash string) (*ChapterInfo, error) {
	query := fmt.Sprintf("select %s from %s where `name` = ? and `hash` = ? limit 1", chapterInfoRows, c.table)
	var resp ChapterInfo
	err := c.conn.QueryRow(&resp, query, title, hash)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewChapterInfoModel returns a model for the database table.
func NewChapterInfoModel(conn sqlx.SqlConn) ChapterInfoModel {
	return &customChapterInfoModel{
		defaultChapterInfoModel: newChapterInfoModel(conn),
	}
}
