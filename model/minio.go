package model

import (
	"bufio"
	"context"
	"github.com/minio/minio-go/v7"
	"log"
)

type S3Model interface {
	Upload(ctx context.Context, filename string, reader *bufio.Reader) (string, error)
}

type miniomodel struct {
	client *minio.Client
	bucket string
}

func (m miniomodel) Upload(ctx context.Context, filename string, reader *bufio.Reader) (string, error) {
	res, err := m.client.PutObject(ctx, m.bucket, filename, reader, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return res.Key, nil
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
