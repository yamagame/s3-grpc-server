package storage

import (
	"context"
	"io"
	"sample/s3-grpc-server/infra/storage/model"
)

type RepositoryClientInterface interface {
	ListBuckets(ctx context.Context) ([]model.Bucket, error)
	CreateBucket(ctx context.Context) error
	PutObject(ctx context.Context, key string, content io.Reader) error
	PutObjectWithString(ctx context.Context, key, content string) error
	GetObject(ctx context.Context, key string) (io.Reader, error)
	GetObjectWithString(ctx context.Context, key string) (string, error)
	DeleteObject(ctx context.Context, key string) error
	ListObjects(ctx context.Context, prefix string, nexttoken *string) ([]model.Object, error)
}

type RepositoryInterface interface {
	CreateBucket(ctx context.Context, req *model.CreateBucket) (*model.CreateBucket, error)
	ListBuckets(ctx context.Context, req *model.ListBuckets) (*model.ListBuckets, error)
	PutObject(ctx context.Context, req *model.PutObject) (*model.PutObject, error)
	GetObject(ctx context.Context, req *model.GetObject) (*model.GetObject, error)
	DeleteObject(ctx context.Context, req *model.DeleteObject) (*model.DeleteObject, error)
	ListObjects(ctx context.Context, req *model.ListObjects) (*model.ListObjects, error)
}
