package storage

import (
	server "sample/s3-grpc-server/proto/grpc-server"
)

type StorageCreateBucketServerData struct {
}

func (x *StorageCreateBucketServerData) Input(req *server.CreateBucketRequest) *CreateBucketEntity {
	return &CreateBucketEntity{}
}

func (x *StorageCreateBucketServerData) Output(res *CreateBucketEntity) *server.CreateBucketResponse {
	return &server.CreateBucketResponse{
		Result: server.Result(res.Result),
	}
}

type StorageListBucketsServerData struct {
}

func (x *StorageListBucketsServerData) Input(req *server.ListBucketsRequest) *ListBucketsEntity {
	return &ListBucketsEntity{}
}

func (x *StorageListBucketsServerData) Output(res *ListBucketsEntity) *server.ListBucketsResponse {
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

type StoragePutObjectServerData struct {
}

func (x *StoragePutObjectServerData) Input(req *server.PutObjectRequest) *PutObjectEntity {
	return &PutObjectEntity{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *StoragePutObjectServerData) Output(res *PutObjectEntity) *server.PutObjectResponse {
	return &server.PutObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type StorageGetObjectServerData struct {
}

func (x *StorageGetObjectServerData) Input(req *server.GetObjectRequest) *GetObjectEntity {
	return &GetObjectEntity{
		Key: req.Key,
	}
}

func (x *StorageGetObjectServerData) Output(res *GetObjectEntity) *server.GetObjectResponse {
	return &server.GetObjectResponse{
		Result:  server.Result(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type StorageDeleteObjectServerData struct {
}

func (x *StorageDeleteObjectServerData) Input(req *server.DeleteObjectRequest) *DeleteObjectEntity {
	return &DeleteObjectEntity{
		Key: req.Key,
	}
}

func (x *StorageDeleteObjectServerData) Output(res *DeleteObjectEntity) *server.DeleteObjectResponse {
	return &server.DeleteObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type StorageListObjectsServerData struct {
}

func (x *StorageListObjectsServerData) Input(req *server.ListObjectsRequest) *ListObjectsEntity {
	return &ListObjectsEntity{
		Prefix: req.Prefix,
	}
}

func (x *StorageListObjectsServerData) Output(res *ListObjectsEntity) *server.ListObjectsResponse {
	return &server.ListObjectsResponse{
		Result: server.Result(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
