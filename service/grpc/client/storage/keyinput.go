package storage

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/entitiy/storage/model"
)

type StorageKeyInput struct {
	scanner *bufio.Scanner
}

func NewKeyInput(scanner *bufio.Scanner) *StorageKeyInput {
	return &StorageKeyInput{
		scanner: scanner,
	}
}

func (x *StorageKeyInput) CreateBucket() *model.CreateBucket {
	return &model.CreateBucket{}
}

func (x *StorageKeyInput) ListBuckets() *model.ListBuckets {
	return &model.ListBuckets{}
}

func (x *StorageKeyInput) PutObject() *model.PutObject {
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

func (x *StorageKeyInput) GetObject() *model.GetObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &model.GetObject{
		Key: key,
	}
}

func (x *StorageKeyInput) DeleteObject() *model.DeleteObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &model.DeleteObject{
		Key: key,
	}
}

func (x *StorageKeyInput) ListObjects() *model.ListObjects {
	fmt.Print("prefix >")
	x.scanner.Scan()
	prefix := x.scanner.Text()
	return &model.ListObjects{
		Prefix: prefix,
	}
}
