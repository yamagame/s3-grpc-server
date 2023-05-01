package storage

import (
	"io"
)

type StorageInterface interface {
	ListBuckets() ([]Bucket, error)
	CreateBucket() error
	PutObject(key string, content io.Reader) error
	PutObjectWithString(key, content string) error
	GetObject(key string) (io.Reader, error)
	GetObjectWithString(key string) (string, error)
	DeleteObject(key string) error
	ListObjects(prefix string, nexttoken *string) ([]Object, error)
}
