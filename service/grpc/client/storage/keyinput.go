package storage

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/entitiy/storage/model"
)

type KeyInput struct {
	scanner *bufio.Scanner
}

func NewKeyInput(scanner *bufio.Scanner) *KeyInput {
	return &KeyInput{
		scanner: scanner,
	}
}

func (x *KeyInput) CreateBucket() *model.CreateBucket {
	return &model.CreateBucket{}
}

func (x *KeyInput) ListBuckets() *model.ListBuckets {
	return &model.ListBuckets{}
}

func (x *KeyInput) PutObject() *model.PutObject {
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

func (x *KeyInput) GetObject() *model.GetObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &model.GetObject{
		Key: key,
	}
}

func (x *KeyInput) DeleteObject() *model.DeleteObject {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &model.DeleteObject{
		Key: key,
	}
}

func (x *KeyInput) ListObjects() *model.ListObjects {
	fmt.Print("prefix >")
	x.scanner.Scan()
	prefix := x.scanner.Text()
	return &model.ListObjects{
		Prefix: prefix,
	}
}
