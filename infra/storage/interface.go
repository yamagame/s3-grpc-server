package storage

import (
	"io"
	"sample/s3-grpc-server/infra/storage/model"
)

type StorageInterface interface {
	ListBuckets() ([]model.Bucket, error)
	CreateBucket() error
	PutObject(key string, content io.Reader) error
	PutObjectWithString(key, content string) error
	GetObject(key string) (io.Reader, error)
	GetObjectWithString(key string) (string, error)
	DeleteObject(key string) error
	ListObjects(prefix string, nexttoken *string) ([]model.Object, error)
}
