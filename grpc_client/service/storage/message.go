package storage

import (
	"sample/s3-grpc-server/infra/storage"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type StorageCreateBucketClientMessage struct {
}

func (x *StorageCreateBucketClientMessage) Input(_ *storage.CreateBucket) *server.CreateBucketRequest {
	return &server.CreateBucketRequest{}
}

func (x *StorageCreateBucketClientMessage) Output(res *server.CreateBucketResponse) *storage.CreateBucket {
	return &storage.CreateBucket{
		Result: storage.StorageResult(res.Result),
	}
}

type StorageListBucketsClientMessage struct {
}

func (x *StorageListBucketsClientMessage) Input(_ *storage.ListBuckets) *server.ListBucketsRequest {
	return &server.ListBucketsRequest{}
}

func (x *StorageListBucketsClientMessage) Output(res *server.ListBucketsResponse) *storage.ListBuckets {
	buckets := make([]storage.Bucket, len(res.Buckets))
	for i, v := range res.Buckets {
		buckets[i].Name = v.Name
	}
	return &storage.ListBuckets{
		Result:  storage.StorageResult(res.Result),
		Buckets: buckets,
	}
}

type StoragePutObjectClientMessage struct {
}

func (x *StoragePutObjectClientMessage) Input(req *storage.PutObject) *server.PutObjectRequest {
	return &server.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *StoragePutObjectClientMessage) Output(res *server.PutObjectResponse) *storage.PutObject {
	return &storage.PutObject{
		Result: storage.StorageResult(res.Result),
	}
}

type StorageGetObjectClientMessage struct {
}

func (x *StorageGetObjectClientMessage) Input(req *storage.GetObject) *server.GetObjectRequest {
	return &server.GetObjectRequest{
		Key: req.Key,
	}
}

func (x *StorageGetObjectClientMessage) Output(res *server.GetObjectResponse) *storage.GetObject {
	return &storage.GetObject{
		Result:  storage.StorageResult(res.Result),
		Content: res.Content,
	}
}

type StorageDeleteObjectClientMessage struct {
}

func (x *StorageDeleteObjectClientMessage) Input(req *storage.DeleteObject) *server.DeleteObjectRequest {
	return &server.DeleteObjectRequest{
		Key: req.Key,
	}
}

func (x *StorageDeleteObjectClientMessage) Output(res *server.DeleteObjectResponse) *storage.DeleteObject {
	return &storage.DeleteObject{
		Result: storage.StorageResult(res.Result),
	}
}

type StorageListObjectsClientMessage struct {
}

func (x *StorageListObjectsClientMessage) Input(req *storage.ListObjects) *server.ListObjectsRequest {
	return &server.ListObjectsRequest{
		Prefix: req.Prefix,
	}
}

func (x *StorageListObjectsClientMessage) Output(res *server.ListObjectsResponse) *storage.ListObjects {
	return &storage.ListObjects{
		Result: storage.StorageResult(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
