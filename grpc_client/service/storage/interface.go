package storage

import "sample/s3-grpc-server/infra/storage/model"

type StorageScannerInterface interface {
	CreateBucket() *model.CreateBucket
	ListBuckets() *model.ListBuckets
	PutObject() *model.PutObject
	GetObject() *model.GetObject
	DeleteObject() *model.DeleteObject
	ListObjects() *model.ListObjects
}
