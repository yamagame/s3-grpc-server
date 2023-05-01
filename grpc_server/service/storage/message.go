package storage

import (
	"sample/s3-grpc-server/infra/storage"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type StorageCreateBucketServerMessage struct {
}

func (x *StorageCreateBucketServerMessage) Input(req *server.CreateBucketRequest) *storage.CreateBucket {
	return &storage.CreateBucket{}
}

func (x *StorageCreateBucketServerMessage) Output(res *storage.CreateBucket) *server.CreateBucketResponse {
	return &server.CreateBucketResponse{
		Result: server.Result(res.Result),
	}
}

type StorageListBucketsServerMessage struct {
}

func (x *StorageListBucketsServerMessage) Input(req *server.ListBucketsRequest) *storage.ListBuckets {
	return &storage.ListBuckets{}
}

func (x *StorageListBucketsServerMessage) Output(res *storage.ListBuckets) *server.ListBucketsResponse {
	bueckts := make([]*server.ListBucketsResponseBucket, len(res.Buckets))
	for i, bucket := range res.Buckets {
		bueckts[i] = &server.ListBucketsResponseBucket{
			Name: bucket.Name,
		}
	}
	return &server.ListBucketsResponse{
		Result:  server.Result(res.Result),
		Buckets: bueckts,
	}
}

type StoragePutObjectServerMessage struct {
}

func (x *StoragePutObjectServerMessage) Input(req *server.PutObjectRequest) *storage.PutObject {
	return &storage.PutObject{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *StoragePutObjectServerMessage) Output(res *storage.PutObject) *server.PutObjectResponse {
	return &server.PutObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type StorageGetObjectServerMessage struct {
}

func (x *StorageGetObjectServerMessage) Input(req *server.GetObjectRequest) *storage.GetObject {
	return &storage.GetObject{
		Key: req.Key,
	}
}

func (x *StorageGetObjectServerMessage) Output(res *storage.GetObject) *server.GetObjectResponse {
	return &server.GetObjectResponse{
		Result:  server.Result(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type StorageDeleteObjectServerMessage struct {
}

func (x *StorageDeleteObjectServerMessage) Input(req *server.DeleteObjectRequest) *storage.DeleteObject {
	return &storage.DeleteObject{
		Key: req.Key,
	}
}

func (x *StorageDeleteObjectServerMessage) Output(res *storage.DeleteObject) *server.DeleteObjectResponse {
	return &server.DeleteObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type StorageListObjectsServerMessage struct {
}

func (x *StorageListObjectsServerMessage) Input(req *server.ListObjectsRequest) *storage.ListObjects {
	return &storage.ListObjects{
		Prefix: req.Prefix,
	}
}

func (x *StorageListObjectsServerMessage) Output(res *storage.ListObjects) *server.ListObjectsResponse {
	return &server.ListObjectsResponse{
		Result: server.Result(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
