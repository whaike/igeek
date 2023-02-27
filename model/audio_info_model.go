package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AudioInfoModel = (*customAudioInfoModel)(nil)

type (
	// AudioInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAudioInfoModel.
	AudioInfoModel interface {
		audioInfoModel
	}

	customAudioInfoModel struct {
		*defaultAudioInfoModel
	}
)

// NewAudioInfoModel returns a model for the database table.
func NewAudioInfoModel(conn sqlx.SqlConn) AudioInfoModel {
	return &customAudioInfoModel{
		defaultAudioInfoModel: newAudioInfoModel(conn),
	}
}
