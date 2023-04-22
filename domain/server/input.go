package server

import (
	"sample/s3-grpc-server/domain"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type AwsServerInput struct {
}

func NewAwsServerInput() *AwsServerInput {
	return &AwsServerInput{}
}

func (x *AwsServerInput) CreateBucket(req *aws.CreateBucketRequest) *domain.CreateBucketEntity {
	return &domain.CreateBucketEntity{}
}

func (x *AwsServerInput) ListBuckets(req *aws.ListBucketsRequest) *domain.ListBucketsEntity {
	return &domain.ListBucketsEntity{}
}

func (x *AwsServerInput) PutObject(req *aws.PutObjectRequest) *domain.PutObjectEntity {
	return &domain.PutObjectEntity{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *AwsServerInput) GetObject(req *aws.GetObjectRequest) *domain.GetObjectEntity {
	return &domain.GetObjectEntity{
		Key: req.Key,
	}
}

func (x *AwsServerInput) DeleteObject(req *aws.DeleteObjectRequest) *domain.DeleteObjectEntity {
	return &domain.DeleteObjectEntity{
		Key: req.Key,
	}
}
