package storage

import (
	"context"
	"io"
	"sample/s3-grpc-server/entitiy/storage/model"
)

type RepositoryInternalInterface interface {
	CreateBucket(ctx context.Context, req *model.CreateBucket) error
	ListBuckets(ctx context.Context, req *model.ListBuckets) ([]model.Bucket, error)
	PutObject(ctx context.Context, req *model.PutObject, content io.Reader) error
	PutObjectWithString(ctx context.Context, req *model.PutObject) error
	GetObject(ctx context.Context, req *model.GetObject) (io.Reader, error)
	GetObjectWithString(ctx context.Context, req *model.GetObject) (string, error)
	DeleteObject(ctx context.Context, req *model.DeleteObject) error
	ListObjects(ctx context.Context, req *model.ListObjects) ([]model.Object, error)
}

type RepositoryInterface interface {
	CreateBucket(ctx context.Context, req *model.CreateBucket) (*model.CreateBucket, error)
	ListBuckets(ctx context.Context, req *model.ListBuckets) (*model.ListBuckets, error)
	PutObject(ctx context.Context, req *model.PutObject) (*model.PutObject, error)
	GetObject(ctx context.Context, req *model.GetObject) (*model.GetObject, error)
	DeleteObject(ctx context.Context, req *model.DeleteObject) (*model.DeleteObject, error)
	ListObjects(ctx context.Context, req *model.ListObjects) (*model.ListObjects, error)
}
