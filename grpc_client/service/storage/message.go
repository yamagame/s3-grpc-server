package storage

import (
	"sample/s3-grpc-server/infra/storage"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type StorageCreateBucketClientMessage struct {
}

func (x *StorageCreateBucketClientMessage) Input(_ *storage.CreateBucketEntity) *server.CreateBucketRequest {
	return &server.CreateBucketRequest{}
}

func (x *StorageCreateBucketClientMessage) Output(res *server.CreateBucketResponse) *storage.CreateBucketEntity {
	return &storage.CreateBucketEntity{
		Result: storage.StorageResult(res.Result),
	}
}

type StorageListBucketsClientMessage struct {
}

func (x *StorageListBucketsClientMessage) Input(_ *storage.ListBucketsEntity) *server.ListBucketsRequest {
	return &server.ListBucketsRequest{}
}

func (x *StorageListBucketsClientMessage) Output(res *server.ListBucketsResponse) *storage.ListBucketsEntity {
	buckets := make([]storage.Bucket, len(res.Buckets))
	for i, v := range res.Buckets {
		buckets[i].Name = v.Name
	}
	return &storage.ListBucketsEntity{
		Result:  storage.StorageResult(res.Result),
		Buckets: buckets,
	}
}

type StoragePutObjectClientMessage struct {
}

func (x *StoragePutObjectClientMessage) Input(req *storage.PutObjectEntity) *server.PutObjectRequest {
	return &server.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *StoragePutObjectClientMessage) Output(res *server.PutObjectResponse) *storage.PutObjectEntity {
	return &storage.PutObjectEntity{
		Result: storage.StorageResult(res.Result),
	}
}

type StorageGetObjectClientMessage struct {
}

func (x *StorageGetObjectClientMessage) Input(req *storage.GetObjectEntity) *server.GetObjectRequest {
	return &server.GetObjectRequest{
		Key: req.Key,
	}
}

func (x *StorageGetObjectClientMessage) Output(res *server.GetObjectResponse) *storage.GetObjectEntity {
	return &storage.GetObjectEntity{
		Result:  storage.StorageResult(res.Result),
		Content: res.Content,
	}
}

type StorageDeleteObjectClientMessage struct {
}

func (x *StorageDeleteObjectClientMessage) Input(req *storage.DeleteObjectEntity) *server.DeleteObjectRequest {
	return &server.DeleteObjectRequest{
		Key: req.Key,
	}
}

func (x *StorageDeleteObjectClientMessage) Output(res *server.DeleteObjectResponse) *storage.DeleteObjectEntity {
	return &storage.DeleteObjectEntity{
		Result: storage.StorageResult(res.Result),
	}
}

type StorageListObjectsClientMessage struct {
}

func (x *StorageListObjectsClientMessage) Input(req *storage.ListObjectsEntity) *server.ListObjectsRequest {
	return &server.ListObjectsRequest{
		Prefix: req.Prefix,
	}
}

func (x *StorageListObjectsClientMessage) Output(res *server.ListObjectsResponse) *storage.ListObjectsEntity {
	return &storage.ListObjectsEntity{
		Result: storage.StorageResult(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
