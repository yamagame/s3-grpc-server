package domain

import (
	server "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type StorageCreateBucketClientData struct {
}

func (x *StorageCreateBucketClientData) Input(_ *CreateBucketEntity) *server.CreateBucketRequest {
	return &server.CreateBucketRequest{}
}

func (x *StorageCreateBucketClientData) Output(res *server.CreateBucketResponse) *CreateBucketEntity {
	return &CreateBucketEntity{
		Result: StorageResult(res.Result),
	}
}

type StorageListBucketsClientData struct {
}

func (x *StorageListBucketsClientData) Input(_ *ListBucketsEntity) *server.ListBucketsRequest {
	return &server.ListBucketsRequest{}
}

func (x *StorageListBucketsClientData) Output(res *server.ListBucketsResponse) *ListBucketsEntity {
	buckets := make([]Bucket, len(res.Buckets))
	for i, v := range res.Buckets {
		buckets[i].Name = v.Name
	}
	return &ListBucketsEntity{
		Result:  StorageResult(res.Result),
		Buckets: buckets,
	}
}

type StoragePutObjectClientData struct {
}

func (x *StoragePutObjectClientData) Input(req *PutObjectEntity) *server.PutObjectRequest {
	return &server.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *StoragePutObjectClientData) Output(res *server.PutObjectResponse) *PutObjectEntity {
	return &PutObjectEntity{
		Result: StorageResult(res.Result),
	}
}

type StorageGetObjectClientData struct {
}

func (x *StorageGetObjectClientData) Input(req *GetObjectEntity) *server.GetObjectRequest {
	return &server.GetObjectRequest{
		Key: req.Key,
	}
}

func (x *StorageGetObjectClientData) Output(res *server.GetObjectResponse) *GetObjectEntity {
	return &GetObjectEntity{
		Result:  StorageResult(res.Result),
		Content: res.Content,
	}
}

type StorageDeleteObjectClientData struct {
}

func (x *StorageDeleteObjectClientData) Input(req *DeleteObjectEntity) *server.DeleteObjectRequest {
	return &server.DeleteObjectRequest{
		Key: req.Key,
	}
}

func (x *StorageDeleteObjectClientData) Output(res *server.DeleteObjectResponse) *DeleteObjectEntity {
	return &DeleteObjectEntity{
		Result: StorageResult(res.Result),
	}
}

type StorageListObjectsClientData struct {
}

func (x *StorageListObjectsClientData) Input(req *ListObjectsEntity) *server.ListObjectsRequest {
	return &server.ListObjectsRequest{
		Prefix: req.Prefix,
	}
}

func (x *StorageListObjectsClientData) Output(res *server.ListObjectsResponse) *ListObjectsEntity {
	return &ListObjectsEntity{
		Result: StorageResult(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
