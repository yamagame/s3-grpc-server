package service

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/domain"
)

type StorageScannerInterface interface {
	CreateBucket() *domain.CreateBucketEntity
	ListBuckets() *domain.ListBucketsEntity
	PutObject() *domain.PutObjectEntity
	GetObject() *domain.GetObjectEntity
	DeleteObject() *domain.DeleteObjectEntity
	ListObjects() *domain.ListObjectsEntity
}

type StorageScanner struct {
	scanner *bufio.Scanner
}

func NewAwsScanner(scanner *bufio.Scanner) *StorageScanner {
	return &StorageScanner{
		scanner: scanner,
	}
}

func (x *StorageScanner) CreateBucket() *domain.CreateBucketEntity {
	return &domain.CreateBucketEntity{}
}

func (x *StorageScanner) ListBuckets() *domain.ListBucketsEntity {
	return &domain.ListBucketsEntity{}
}

func (x *StorageScanner) PutObject() *domain.PutObjectEntity {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	fmt.Print("content >")
	x.scanner.Scan()
	content := x.scanner.Text()
	return &domain.PutObjectEntity{
		Key:     key,
		Content: content,
	}
}

func (x *StorageScanner) GetObject() *domain.GetObjectEntity {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &domain.GetObjectEntity{
		Key: key,
	}
}

func (x *StorageScanner) DeleteObject() *domain.DeleteObjectEntity {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &domain.DeleteObjectEntity{
		Key: key,
	}
}

func (x *StorageScanner) ListObjects() *domain.ListObjectsEntity {
	fmt.Print("prefix >")
	x.scanner.Scan()
	prefix := x.scanner.Text()
	return &domain.ListObjectsEntity{
		Prefix: prefix,
	}
}
