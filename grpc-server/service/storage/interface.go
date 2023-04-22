package storage

import (
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ClientInterface interface {
	ListBuckets() (*s3.ListBucketsOutput, error)
	CreateBucket() (*s3.CreateBucketOutput, error)
	PutObject(key string, body io.Reader) (*s3.PutObjectOutput, error)
	PutObjectWithString(key, content string) (*s3.PutObjectOutput, error)
	GetObject(key string) (*s3.GetObjectOutput, error)
	GetObjectWithString(key string) (string, error)
	DeleteObject(key string) (*s3.DeleteObjectOutput, error)
	ListObjects(prefix string, nexttoken *string) (*s3.ListObjectsV2Output, error)
}
