package storage

import (
	"sample/s3-grpc-server/infra/storage/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateBucketDTO struct {
}

func (x *CreateBucketDTO) Input(_ *model.CreateBucket) *server.CreateBucketRequest {
	return &server.CreateBucketRequest{}
}

func (x *CreateBucketDTO) Output(res *server.CreateBucketResponse) *model.CreateBucket {
	return &model.CreateBucket{
		Result: model.StorageResult(res.Result),
	}
}

type ListBucketsDTO struct {
}

func (x *ListBucketsDTO) Input(_ *model.ListBuckets) *server.ListBucketsRequest {
	return &server.ListBucketsRequest{}
}

func (x *ListBucketsDTO) Output(res *server.ListBucketsResponse) *model.ListBuckets {
	buckets := make([]model.Bucket, len(res.Buckets))
	for i, v := range res.Buckets {
		buckets[i].Name = v.Name
	}
	return &model.ListBuckets{
		Result:  model.StorageResult(res.Result),
		Buckets: buckets,
	}
}

type PutObjectDTO struct {
}

func (x *PutObjectDTO) Input(req *model.PutObject) *server.PutObjectRequest {
	return &server.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *PutObjectDTO) Output(res *server.PutObjectResponse) *model.PutObject {
	return &model.PutObject{
		Result: model.StorageResult(res.Result),
	}
}

type GetObjectDTO struct {
}

func (x *GetObjectDTO) Input(req *model.GetObject) *server.GetObjectRequest {
	return &server.GetObjectRequest{
		Key: req.Key,
	}
}

func (x *GetObjectDTO) Output(res *server.GetObjectResponse) *model.GetObject {
	return &model.GetObject{
		Result:  model.StorageResult(res.Result),
		Content: res.Content,
	}
}

type DeleteObjectDTO struct {
}

func (x *DeleteObjectDTO) Input(req *model.DeleteObject) *server.DeleteObjectRequest {
	return &server.DeleteObjectRequest{
		Key: req.Key,
	}
}

func (x *DeleteObjectDTO) Output(res *server.DeleteObjectResponse) *model.DeleteObject {
	return &model.DeleteObject{
		Result: model.StorageResult(res.Result),
	}
}

type ListObjectsDTO struct {
}

func (x *ListObjectsDTO) Input(req *model.ListObjects) *server.ListObjectsRequest {
	return &server.ListObjectsRequest{
		Prefix: req.Prefix,
	}
}

func (x *ListObjectsDTO) Output(res *server.ListObjectsResponse) *model.ListObjects {
	return &model.ListObjects{
		Result: model.StorageResult(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
