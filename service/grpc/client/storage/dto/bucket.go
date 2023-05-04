package dto

import (
	"sample/s3-grpc-server/infra/storage/model"
	"sample/s3-grpc-server/libs/dto"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateBucket struct {
}

func (x *CreateBucket) ToInfra(req *model.CreateBucket, call func(req *grpc_server.CreateBucketRequest) (*grpc_server.CreateBucketResponse, error)) (*model.CreateBucket, error) {
	return dto.ToInfra[model.CreateBucket, grpc_server.CreateBucketRequest, grpc_server.CreateBucketResponse](x, req, call)
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

func (x *ListBuckets) ToInfra(req *model.ListBuckets, call func(req *grpc_server.ListBucketsRequest) (*grpc_server.ListBucketsResponse, error)) (*model.ListBuckets, error) {
	return dto.ToInfra[model.ListBuckets, grpc_server.ListBucketsRequest, grpc_server.ListBucketsResponse](x, req, call)
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
