package storage

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/infra/storage"
)

type StorageScannerInterface interface {
	CreateBucket() *storage.CreateBucket
	ListBuckets() *storage.ListBuckets
	PutObject() *storage.PutObject
	GetObject() *storage.GetObject
	DeleteObject() *storage.DeleteObject
	ListObjects() *storage.ListObjects
}

type StorageScanner struct {
	scanner *bufio.Scanner
}

func NewScanner(scanner *bufio.Scanner) *StorageScanner {
	return &StorageScanner{
		scanner: scanner,
	}
}

func (x *StorageScanner) CreateBucket() *storage.CreateBucket {
	return &storage.CreateBucket{}
}

func (x *StorageScanner) ListBuckets() *storage.ListBuckets {
	return &storage.ListBuckets{}
}

func (x *StorageScanner) PutObject() *storage.PutObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	fmt.Print("content >")
	x.scanner.Scan()
	content := x.scanner.Text()
	return &storage.PutObject{
		Key:     key,
		Content: content,
	}
}

func (x *StorageScanner) GetObject() *storage.GetObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &storage.GetObject{
		Key: key,
	}
}

func (x *StorageScanner) DeleteObject() *storage.DeleteObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &storage.DeleteObject{
		Key: key,
	}
}

func (x *StorageScanner) ListObjects() *storage.ListObjects {
	fmt.Print("prefix >")
	x.scanner.Scan()
	prefix := x.scanner.Text()
	return &storage.ListObjects{
		Prefix: prefix,
	}
}
