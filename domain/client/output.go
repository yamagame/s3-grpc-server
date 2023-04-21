package client

import (
	"sample/s3-grpc-server/domain"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type AwsClientOutput struct {
}

func NewAwsClientOutput() *AwsClientOutput {
	return &AwsClientOutput{}
}

type CreateBucketOutput struct{}

func (x *AwsClientOutput) CreateBucket(res *aws.CreateBucketResponse) *domain.CreateBucketEntity {
	return &domain.CreateBucketEntity{
		Result: domain.AwsResult(res.Result),
	}
}

type ListBucketsOutput struct{}

func (x *AwsClientOutput) ListBuckets(res *aws.ListBucketsResponse) *domain.ListBucketsEntity {
	buckets := make([]domain.Bucket, len(res.Buckets))
	for i, v := range res.Buckets {
		buckets[i].Name = v.Name
	}
	return &domain.ListBucketsEntity{
		Result:  domain.AwsResult(res.Result),
		Buckets: buckets,
	}
}

type PutObjectOutput struct {
	Key     string
	Content string
}

func (x *AwsClientOutput) PutObject(res *aws.PutObjectResponse) *domain.PutObjectEntity {
	return &domain.PutObjectEntity{
		Result: domain.AwsResult(res.Result),
	}
}

type GetObjectOutput struct {
	Key string
}

func (x *AwsClientOutput) GetObject(res *aws.GetObjectResponse) *domain.GetObjectEntity {
	return &domain.GetObjectEntity{
		Result:  domain.AwsResult(res.Result),
		Content: res.Content,
	}
}

type DeleteObjectOutput struct {
	Key string
}

func (x *AwsClientOutput) DeleteObject(res *aws.DeleteObjectResponse) *domain.DeleteObjectEntity {
	return &domain.DeleteObjectEntity{
		Result: domain.AwsResult(res.Result),
	}
}