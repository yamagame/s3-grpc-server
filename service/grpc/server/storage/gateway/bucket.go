package gateway

import (
	"sample/s3-grpc-server/infra/storage/model"
	"sample/s3-grpc-server/libs/gateway"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateBucket struct {
}

func (x *CreateBucket) ToDomain(req *grpc_server.CreateBucketRequest, call func(table *model.CreateBucket) (*model.CreateBucket, error)) (*grpc_server.CreateBucketResponse, error) {
	return gateway.ToDomain[model.CreateBucket, grpc_server.CreateBucketRequest, grpc_server.CreateBucketResponse](x, req, call)
}

func (x *CreateBucket) Input(req *grpc_server.CreateBucketRequest) *model.CreateBucket {
	return &model.CreateBucket{}
}

func (x *CreateBucket) Output(res *model.CreateBucket) *grpc_server.CreateBucketResponse {
	return &grpc_server.CreateBucketResponse{
		Result: grpc_server.Result(res.Result),
	}
}

type ListBuckets struct {
}

func (x *ListBuckets) ToDomain(req *grpc_server.ListBucketsRequest, call func(table *model.ListBuckets) (*model.ListBuckets, error)) (*grpc_server.ListBucketsResponse, error) {
	return gateway.ToDomain[model.ListBuckets, grpc_server.ListBucketsRequest, grpc_server.ListBucketsResponse](x, req, call)
}

func (x *ListBuckets) Input(req *grpc_server.ListBucketsRequest) *model.ListBuckets {
	return &model.ListBuckets{}
}

func (x *ListBuckets) Output(res *model.ListBuckets) *grpc_server.ListBucketsResponse {
	bueckts := make([]*grpc_server.ListBucketsResponseBucket, len(res.Buckets))
	for i, bucket := range res.Buckets {
		bueckts[i] = &grpc_server.ListBucketsResponseBucket{
			Name: bucket.Name,
		}
	}
	return &grpc_server.ListBucketsResponse{
		Result:  grpc_server.Result(res.Result),
		Buckets: bueckts,
	}
}
