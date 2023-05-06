package gateway

import (
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateBucket struct {
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
