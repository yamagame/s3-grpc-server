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

type CreateBucketInput struct{}

func (x *AwsClientInput) CreateBucket(_ *domain.CreateBucketEntity) *aws.CreateBucketRequest {
	return &aws.CreateBucketRequest{}
}

type ListBucketsInput struct{}

func (x *AwsClientInput) ListBuckets(_ *domain.ListBucketsEntity) *aws.ListBucketsRequest {
	return &aws.ListBucketsRequest{}
}

type PutObjectInput struct {
	Key     string
	Content string
}

func (x *AwsClientInput) PutObject(req *domain.PutObjectEntity) *aws.PutObjectRequest {
	return &aws.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

type GetObjectInput struct {
	Key string
}

func (x *AwsClientInput) GetObject(req *domain.GetObjectEntity) *aws.GetObjectRequest {
	return &aws.GetObjectRequest{
		Key: req.Key,
	}
}

type DeleteObjectInput struct {
	Key string
}

func (x *AwsClientInput) DeleteObject(req *domain.DeleteObjectEntity) *aws.DeleteObjectRequest {
	return &aws.DeleteObjectRequest{
		Key: req.Key,
	}
}
