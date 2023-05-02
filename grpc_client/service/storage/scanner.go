package storage

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/infra/storage/model"
)

type StorageScannerInterface interface {
	CreateBucket() *model.CreateBucket
	ListBuckets() *model.ListBuckets
	PutObject() *model.PutObject
	GetObject() *model.GetObject
	DeleteObject() *model.DeleteObject
	ListObjects() *model.ListObjects
}

type StorageScanner struct {
	scanner *bufio.Scanner
}

func NewScanner(scanner *bufio.Scanner) *StorageScanner {
	return &StorageScanner{
		scanner: scanner,
	}
}

func (x *StorageScanner) CreateBucket() *model.CreateBucket {
	return &model.CreateBucket{}
}

func (x *StorageScanner) ListBuckets() *model.ListBuckets {
	return &model.ListBuckets{}
}

func (x *StorageScanner) PutObject() *model.PutObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	fmt.Print("content >")
	x.scanner.Scan()
	content := x.scanner.Text()
	return &model.PutObject{
		Key:     key,
		Content: content,
	}
}

func (x *StorageScanner) GetObject() *model.GetObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &model.GetObject{
		Key: key,
	}
}

func (x *StorageScanner) DeleteObject() *model.DeleteObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &model.DeleteObject{
		Key: key,
	}
}

func (x *StorageScanner) ListObjects() *model.ListObjects {
	fmt.Print("prefix >")
	x.scanner.Scan()
	prefix := x.scanner.Text()
	return &model.ListObjects{
		Prefix: prefix,
	}
}
