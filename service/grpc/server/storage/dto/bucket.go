package dto

import (
	"sample/s3-grpc-server/infra/storage/model"
	"sample/s3-grpc-server/libs/dto"
	"sample/s3-grpc-server/proto/grpc_server"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateBucket struct {
}

func (x *CreateBucket) Domain(req *grpc_server.CreateBucketRequest, call func(table *model.CreateBucket) (*model.CreateBucket, error)) (*grpc_server.CreateBucketResponse, error) {
	return dto.ToDomain[model.CreateBucket, grpc_server.CreateBucketRequest, grpc_server.CreateBucketResponse](x, req, call)
}

func (x *CreateBucket) Input(req *server.CreateBucketRequest) *model.CreateBucket {
	return &model.CreateBucket{}
}

func (x *CreateBucket) Output(res *model.CreateBucket) *server.CreateBucketResponse {
	return &server.CreateBucketResponse{
		Result: server.Result(res.Result),
	}
}

type ListBuckets struct {
}

func (x *ListBuckets) Domain(req *grpc_server.ListBucketsRequest, call func(table *model.ListBuckets) (*model.ListBuckets, error)) (*grpc_server.ListBucketsResponse, error) {
	return dto.ToDomain[model.ListBuckets, grpc_server.ListBucketsRequest, grpc_server.ListBucketsResponse](x, req, call)
}

func (x *ListBuckets) Input(req *server.ListBucketsRequest) *model.ListBuckets {
	return &model.ListBuckets{}
}

func (x *ListBuckets) Output(res *model.ListBuckets) *server.ListBucketsResponse {
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
