package dto

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
