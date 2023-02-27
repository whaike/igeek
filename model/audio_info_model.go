package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AudioInfoModel = (*customAudioInfoModel)(nil)

type (
	// AudioInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAudioInfoModel.
	AudioInfoModel interface {
		audioInfoModel
		Find(title, hash string) (*AudioInfo, error)
	}

	customAudioInfoModel struct {
		*defaultAudioInfoModel
	}
)

func (c customAudioInfoModel) Find(title, hash string) (*AudioInfo, error) {
	query := fmt.Sprintf("select %s from %s where `title` = ? and `hash` = ? limit 1", audioInfoRows, c.table)
	var resp AudioInfo
	err := c.conn.QueryRow(&resp, query, title, hash)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

// NewAudioInfoModel returns a model for the database table.
func NewAudioInfoModel(conn sqlx.SqlConn) AudioInfoModel {
	return &customAudioInfoModel{
		defaultAudioInfoModel: newAudioInfoModel(conn),
	}
}
