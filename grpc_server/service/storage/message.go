package storage

import (
	"sample/s3-grpc-server/infra/storage"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type StorageCreateBucketServerMessage struct {
}

func (x *StorageCreateBucketServerMessage) Input(req *server.CreateBucketRequest) *storage.CreateBucketEntity {
	return &storage.CreateBucketEntity{}
}

func (x *StorageCreateBucketServerMessage) Output(res *storage.CreateBucketEntity) *server.CreateBucketResponse {
	return &server.CreateBucketResponse{
		Result: server.Result(res.Result),
	}
}

type StorageListBucketsServerMessage struct {
}

func (x *StorageListBucketsServerMessage) Input(req *server.ListBucketsRequest) *storage.ListBucketsEntity {
	return &storage.ListBucketsEntity{}
}

func (x *StorageListBucketsServerMessage) Output(res *storage.ListBucketsEntity) *server.ListBucketsResponse {
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

func (x *StoragePutObjectServerMessage) Input(req *server.PutObjectRequest) *storage.PutObjectEntity {
	return &storage.PutObjectEntity{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *StoragePutObjectServerMessage) Output(res *storage.PutObjectEntity) *server.PutObjectResponse {
	return &server.PutObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type StorageGetObjectServerMessage struct {
}

func (x *StorageGetObjectServerMessage) Input(req *server.GetObjectRequest) *storage.GetObjectEntity {
	return &storage.GetObjectEntity{
		Key: req.Key,
	}
}

func (x *StorageGetObjectServerMessage) Output(res *storage.GetObjectEntity) *server.GetObjectResponse {
	return &server.GetObjectResponse{
		Result:  server.Result(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type StorageDeleteObjectServerMessage struct {
}

func (x *StorageDeleteObjectServerMessage) Input(req *server.DeleteObjectRequest) *storage.DeleteObjectEntity {
	return &storage.DeleteObjectEntity{
		Key: req.Key,
	}
}

func (x *StorageDeleteObjectServerMessage) Output(res *storage.DeleteObjectEntity) *server.DeleteObjectResponse {
	return &server.DeleteObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type StorageListObjectsServerMessage struct {
}

func (x *StorageListObjectsServerMessage) Input(req *server.ListObjectsRequest) *storage.ListObjectsEntity {
	return &storage.ListObjectsEntity{
		Prefix: req.Prefix,
	}
}

func (x *StorageListObjectsServerMessage) Output(res *storage.ListObjectsEntity) *server.ListObjectsResponse {
	return &server.ListObjectsResponse{
		Result: server.Result(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
