package storage

import "sample/s3-grpc-server/entitiy/storage/model"

type StorageScannerInterface interface {
	CreateBucket() *model.CreateBucket
	ListBuckets() *model.ListBuckets
	PutObject() *model.PutObject
	GetObject() *model.GetObject
	DeleteObject() *model.DeleteObject
	ListObjects() *model.ListObjects
}
