package storage

import (
	"context"
	"io"
	"sample/s3-grpc-server/infra/storage/model"
)

type StorageInterface interface {
	ListBuckets(ctx context.Context) ([]model.Bucket, error)
	CreateBucket(ctx context.Context) error
	PutObject(ctx context.Context, key string, content io.Reader) error
	PutObjectWithString(ctx context.Context, key, content string) error
	GetObject(ctx context.Context, key string) (io.Reader, error)
	GetObjectWithString(ctx context.Context, key string) (string, error)
	DeleteObject(ctx context.Context, key string) error
	ListObjects(ctx context.Context, prefix string, nexttoken *string) ([]model.Object, error)
}
