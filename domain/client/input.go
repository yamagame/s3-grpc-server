package client

import (
	"sample/s3-grpc-server/domain"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type AwsClientInput struct {
}

func NewAwsClientInput() *AwsClientInput {
	return &AwsClientInput{}
}

func (x *AwsClientInput) CreateBucket(_ *domain.CreateBucketEntity) *aws.CreateBucketRequest {
	return &aws.CreateBucketRequest{}
}

func (x *AwsClientInput) ListBuckets(_ *domain.ListBucketsEntity) *aws.ListBucketsRequest {
	return &aws.ListBucketsRequest{}
}

func (x *AwsClientInput) PutObject(req *domain.PutObjectEntity) *aws.PutObjectRequest {
	return &aws.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *AwsClientInput) GetObject(req *domain.GetObjectEntity) *aws.GetObjectRequest {
	return &aws.GetObjectRequest{
		Key: req.Key,
	}
}

func (x *AwsClientInput) DeleteObject(req *domain.DeleteObjectEntity) *aws.DeleteObjectRequest {
	return &aws.DeleteObjectRequest{
		Key: req.Key,
	}
}
