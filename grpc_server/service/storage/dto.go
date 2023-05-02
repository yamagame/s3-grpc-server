package storage

import (
	"sample/s3-grpc-server/infra/storage/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateBucketDTO struct {
}

func (x *CreateBucketDTO) Input(req *server.CreateBucketRequest) *model.CreateBucket {
	return &model.CreateBucket{}
}

func (x *CreateBucketDTO) Output(res *model.CreateBucket) *server.CreateBucketResponse {
	return &server.CreateBucketResponse{
		Result: server.Result(res.Result),
	}
}

type ListBucketsDTO struct {
}

func (x *ListBucketsDTO) Input(req *server.ListBucketsRequest) *model.ListBuckets {
	return &model.ListBuckets{}
}

func (x *ListBucketsDTO) Output(res *model.ListBuckets) *server.ListBucketsResponse {
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

type PutObjecDTO struct {
}

func (x *PutObjecDTO) Input(req *server.PutObjectRequest) *model.PutObject {
	return &model.PutObject{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *PutObjecDTO) Output(res *model.PutObject) *server.PutObjectResponse {
	return &server.PutObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type GetObjectDTO struct {
}

func (x *GetObjectDTO) Input(req *server.GetObjectRequest) *model.GetObject {
	return &model.GetObject{
		Key: req.Key,
	}
}

func (x *GetObjectDTO) Output(res *model.GetObject) *server.GetObjectResponse {
	return &server.GetObjectResponse{
		Result:  server.Result(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type DeleteObjectDTO struct {
}

func (x *DeleteObjectDTO) Input(req *server.DeleteObjectRequest) *model.DeleteObject {
	return &model.DeleteObject{
		Key: req.Key,
	}
}

func (x *DeleteObjectDTO) Output(res *model.DeleteObject) *server.DeleteObjectResponse {
	return &server.DeleteObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type ListObjectsDTO struct {
}

func (x *ListObjectsDTO) Input(req *server.ListObjectsRequest) *model.ListObjects {
	return &model.ListObjects{
		Prefix: req.Prefix,
	}
}

func (x *ListObjectsDTO) Output(res *model.ListObjects) *server.ListObjectsResponse {
	return &server.ListObjectsResponse{
		Result: server.Result(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
