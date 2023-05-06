package gateway

import (
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateBucket struct {
}

func (x *CreateBucket) Input(_ *model.CreateBucket) *grpc_server.CreateBucketRequest {
	return &grpc_server.CreateBucketRequest{}
}

func (x *CreateBucket) Output(res *grpc_server.CreateBucketResponse) *model.CreateBucket {
	return &model.CreateBucket{
		Result: model.StorageResult(res.Result),
	}
}

type ListBuckets struct {
}

func (x *ListBuckets) Input(_ *model.ListBuckets) *grpc_server.ListBucketsRequest {
	return &grpc_server.ListBucketsRequest{}
}

func (x *ListBuckets) Output(res *grpc_server.ListBucketsResponse) *model.ListBuckets {
	buckets := make([]model.Bucket, len(res.Buckets))
	for i, v := range res.Buckets {
		buckets[i].Name = v.Name
	}
	return &model.ListBuckets{
		Result:  model.StorageResult(res.Result),
		Buckets: buckets,
	}
}
