package service

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/external/storage"
)

type StorageScannerInterface interface {
	CreateBucket() *storage.CreateBucketEntity
	ListBuckets() *storage.ListBucketsEntity
	PutObject() *storage.PutObjectEntity
	GetObject() *storage.GetObjectEntity
	DeleteObject() *storage.DeleteObjectEntity
	ListObjects() *storage.ListObjectsEntity
}

type StorageScanner struct {
	scanner *bufio.Scanner
}

func NewScanner(scanner *bufio.Scanner) *StorageScanner {
	return &StorageScanner{
		scanner: scanner,
	}
}

func (x *StorageScanner) CreateBucket() *storage.CreateBucketEntity {
	return &storage.CreateBucketEntity{}
}

func (x *StorageScanner) ListBuckets() *storage.ListBucketsEntity {
	return &storage.ListBucketsEntity{}
}

func (x *StorageScanner) PutObject() *storage.PutObjectEntity {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	fmt.Print("content >")
	x.scanner.Scan()
	content := x.scanner.Text()
	return &storage.PutObjectEntity{
		Key:     key,
		Content: content,
	}
}

func (x *StorageScanner) GetObject() *storage.GetObjectEntity {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &storage.GetObjectEntity{
		Key: key,
	}
}

func (x *StorageScanner) DeleteObject() *storage.DeleteObjectEntity {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &storage.DeleteObjectEntity{
		Key: key,
	}
}

func (x *StorageScanner) ListObjects() *storage.ListObjectsEntity {
	fmt.Print("prefix >")
	x.scanner.Scan()
	prefix := x.scanner.Text()
	return &storage.ListObjectsEntity{
		Prefix: prefix,
	}
}
