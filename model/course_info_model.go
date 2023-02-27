package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"igeek/pkg"
)

var _ CourseInfoModel = (*customCourseInfoModel)(nil)

type (
	// CourseInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCourseInfoModel.
	CourseInfoModel interface {
		courseInfoModel
		FindAll(ctx context.Context) ([]*CourseInfo, error)
		FindByTitle(ctx context.Context, title string) (*CourseInfo, error)
		CreateOne(ctx context.Context, data *CourseInfo) error
	}

	customCourseInfoModel struct {
		*defaultCourseInfoModel
	}
)

func (c customCourseInfoModel) FindByTitle(ctx context.Context, title string) (*CourseInfo, error) {
	query := fmt.Sprintf("select %s from %s where `title` = ? limit 1", courseInfoRows, c.table)
	var resp CourseInfo
	err := c.conn.QueryRowCtx(ctx, &resp, query, title)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customCourseInfoModel) CreateOne(ctx context.Context, data *CourseInfo) error {
	data.CreateAt = pkg.CommonTime()
	_, err := c.Insert(ctx, data)
	return err
}

func (c customCourseInfoModel) FindAll(ctx context.Context) ([]*CourseInfo, error) {
	query := fmt.Sprintf("select %s from %s ", courseInfoRows, c.table)
	resp := make([]*CourseInfo, 0)
	err := c.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewCourseInfoModel returns a model for the database table.
func NewCourseInfoModel(conn sqlx.SqlConn) CourseInfoModel {
	return &customCourseInfoModel{
		defaultCourseInfoModel: newCourseInfoModel(conn),
	}
}
