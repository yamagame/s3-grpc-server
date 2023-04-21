package server

import (
	"sample/s3-grpc-server/domain"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type AwsServerOutput struct {
}

func NewAwsServerOutput() *AwsServerOutput {
	return &AwsServerOutput{}
}

type CreateBucketInput struct{}

func (x *AwsServerOutput) CreateBucket(res *domain.CreateBucketEntity) *aws.CreateBucketResponse {
	return &aws.CreateBucketResponse{
		Result: aws.Result(res.Result),
	}
}

type ListBucketsInput struct{}

func (x *AwsServerOutput) ListBuckets(res *domain.ListBucketsEntity) *aws.ListBucketsResponse {
	bueckts := make([]*aws.ListBucketsResponseBucket, len(res.Buckets))
	for i, bucket := range res.Buckets {
		bueckts[i] = &aws.ListBucketsResponseBucket{
			Name: bucket.Name,
		}
	}
	return &aws.ListBucketsResponse{
		Result:  aws.Result(res.Result),
		Buckets: bueckts,
	}
}

type PutObjectInput struct {
	Key     string
	Content string
}

func (x *AwsServerOutput) PutObject(res *domain.PutObjectEntity) *aws.PutObjectResponse {
	return &aws.PutObjectResponse{
		Result: aws.Result(res.Result),
		Key:    res.Key,
	}
}

type GetObjectInput struct {
	Key string
}

func (x *AwsServerOutput) GetObject(res *domain.GetObjectEntity) *aws.GetObjectResponse {
	return &aws.GetObjectResponse{
		Result:  aws.Result(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type DeleteObjectInput struct {
	Key string
}

func (x *AwsServerOutput) DeleteObject(res *domain.DeleteObjectEntity) *aws.DeleteObjectResponse {
	return &aws.DeleteObjectResponse{
		Result: aws.Result(res.Result),
		Key:    res.Key,
	}
}