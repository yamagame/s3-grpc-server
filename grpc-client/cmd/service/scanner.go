package service

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/domain"
)

type AwsScannerInterface interface {
	CreateBucket() *domain.CreateBucketEntity
	ListBuckets() *domain.ListBucketsEntity
	PutObject() *domain.PutObjectEntity
	GetObject() *domain.GetObjectEntity
	DeleteObject() *domain.DeleteObjectEntity
}

type AwsScanner struct {
	scanner *bufio.Scanner
}

func NewAwsScanner(scanner *bufio.Scanner) *AwsScanner {
	return &AwsScanner{
		scanner: scanner,
	}
}

func (x *AwsScanner) CreateBucket() *domain.CreateBucketEntity {
	return &domain.CreateBucketEntity{}
}

func (x *AwsScanner) ListBuckets() *domain.ListBucketsEntity {
	return &domain.ListBucketsEntity{}
}

func (x *AwsScanner) PutObject() *domain.PutObjectEntity {
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

func (x *AwsScanner) GetObject() *domain.GetObjectEntity {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &domain.GetObjectEntity{
		Key: key,
	}
}

func (x *AwsScanner) DeleteObject() *domain.DeleteObjectEntity {
	fmt.Print("key >")
	x.scanner.Scan()
	key := x.scanner.Text()
	return &domain.DeleteObjectEntity{
		Key: key,
	}
}
