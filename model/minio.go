package model

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"igeek/pkg"
	"io"
	"log"
)

type S3Model interface {
	Upload(ctx context.Context, filename string, data []byte) (string, error)
	FileExist(filename string) (string, bool) // 判断文件是否存在,如果存在则返回key和true
	GetObject(filename string) ([]byte, error)
}

type miniomodel struct {
	client *minio.Client
	bucket string
}

func (m miniomodel) FileExist(filename string) (string, bool) {
	b, err := m.GetObject(filename)
	if err != nil || b == nil {
		return "", false
	}
	hash, err := pkg.FileMd5(b)
	if err != nil {
		return "", false
	}
	return hash, true
}

func (m miniomodel) GetObject(filename string) ([]byte, error) {
	object, err := m.client.GetObject(context.Background(), m.bucket, filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	buf := &bytes.Buffer{}

	if _, err = io.Copy(buf, object); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}

func (m miniomodel) Upload(ctx context.Context, filename string, data []byte) (string, error) {
	res, err := m.client.PutObject(ctx, m.bucket, filename, bytes.NewReader(data), -1, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return res.Location, nil
}

func NewMinioModel(client *minio.Client, bucket string) S3Model {
	// 如果没有这个bucket就新建一个
	err := client.MakeBucket(context.Background(), bucket, minio.MakeBucketOptions{}) // 本地使用，未设region参数
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := client.BucketExists(context.Background(), bucket)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucket)
		} else {
			panic(err.Error())
		}
	} else {
		log.Printf("Successfully created %s\n", bucket)
	}
	return &miniomodel{
		client: client,
		bucket: bucket,
	}
}
