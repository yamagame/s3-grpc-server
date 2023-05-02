package dto

import (
	"sample/s3-grpc-server/infra/storage/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateBucket struct {
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
