package storage

import (
	"sample/s3-grpc-server/infra/storage/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type StorageCreateBucketClientMessage struct {
}

func (x *StorageCreateBucketClientMessage) Input(_ *model.CreateBucket) *server.CreateBucketRequest {
	return &server.CreateBucketRequest{}
}

func (x *StorageCreateBucketClientMessage) Output(res *server.CreateBucketResponse) *model.CreateBucket {
	return &model.CreateBucket{
		Result: model.StorageResult(res.Result),
	}
}

type StorageListBucketsClientMessage struct {
}

func (x *StorageListBucketsClientMessage) Input(_ *model.ListBuckets) *server.ListBucketsRequest {
	return &server.ListBucketsRequest{}
}

func (x *StorageListBucketsClientMessage) Output(res *server.ListBucketsResponse) *model.ListBuckets {
	buckets := make([]model.Bucket, len(res.Buckets))
	for i, v := range res.Buckets {
		buckets[i].Name = v.Name
	}
	return &model.ListBuckets{
		Result:  model.StorageResult(res.Result),
		Buckets: buckets,
	}
}

type StoragePutObjectClientMessage struct {
}

func (x *StoragePutObjectClientMessage) Input(req *model.PutObject) *server.PutObjectRequest {
	return &server.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *StoragePutObjectClientMessage) Output(res *server.PutObjectResponse) *model.PutObject {
	return &model.PutObject{
		Result: model.StorageResult(res.Result),
	}
}

type StorageGetObjectClientMessage struct {
}

func (x *StorageGetObjectClientMessage) Input(req *model.GetObject) *server.GetObjectRequest {
	return &server.GetObjectRequest{
		Key: req.Key,
	}
}

func (x *StorageGetObjectClientMessage) Output(res *server.GetObjectResponse) *model.GetObject {
	return &model.GetObject{
		Result:  model.StorageResult(res.Result),
		Content: res.Content,
	}
}

type StorageDeleteObjectClientMessage struct {
}

func (x *StorageDeleteObjectClientMessage) Input(req *model.DeleteObject) *server.DeleteObjectRequest {
	return &server.DeleteObjectRequest{
		Key: req.Key,
	}
}

func (x *StorageDeleteObjectClientMessage) Output(res *server.DeleteObjectResponse) *model.DeleteObject {
	return &model.DeleteObject{
		Result: model.StorageResult(res.Result),
	}
}

type StorageListObjectsClientMessage struct {
}

func (x *StorageListObjectsClientMessage) Input(req *model.ListObjects) *server.ListObjectsRequest {
	return &server.ListObjectsRequest{
		Prefix: req.Prefix,
	}
}

func (x *StorageListObjectsClientMessage) Output(res *server.ListObjectsResponse) *model.ListObjects {
	return &model.ListObjects{
		Result: model.StorageResult(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
