package svc

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"igeek/internal/config"
	"igeek/model"
)

type ServiceContext struct {
	Config       config.Config
	CourseModel  model.CourseInfoModel  // 课程信息表
	ChapterModel model.ChapterInfoModel // 文章信息表
	AudioModel   model.AudioInfoModel   // 音频信息表
	ChapterS3    model.S3Model          // 章节内容(文件)对象存储（临时）
	AudioS3      model.S3Model          // 章节内容(音频)对象存储
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	chapterS3, err := minio.New(c.Minio.Endpoint, &minio.Options{
		Secure: c.Minio.UseSSL,
		Creds:  credentials.NewStaticV4(c.Minio.AccessKeyID, c.Minio.SecretAccessKey, ""),
	})
	if err != nil {
		panic(err.Error())
	}

	return &ServiceContext{
		Config:       c,
		CourseModel:  model.NewCourseInfoModel(conn),
		ChapterModel: model.NewChapterInfoModel(conn),
		AudioModel:   model.NewAudioInfoModel(conn),
		ChapterS3:    model.NewMinioModel(chapterS3, "chapters"),
		AudioS3:      model.NewMinioModel(chapterS3, "audios"),
	}
}
