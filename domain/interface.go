package domain

import aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"

type AwsClientInputInterface interface {
	CreateBucket(req *CreateBucketEntity) *aws.CreateBucketRequest
	ListBuckets(req *ListBucketsEntity) *aws.ListBucketsRequest
	PutObject(req *PutObjectEntity) *aws.PutObjectRequest
	GetObject(req *GetObjectEntity) *aws.GetObjectRequest
	DeleteObject(req *DeleteObjectEntity) *aws.DeleteObjectRequest
}

type AwsClientOutputInterface interface {
	CreateBucket(res *aws.CreateBucketResponse) *CreateBucketEntity
	ListBuckets(res *aws.ListBucketsResponse) *ListBucketsEntity
	PutObject(res *aws.PutObjectResponse) *PutObjectEntity
	GetObject(res *aws.GetObjectResponse) *GetObjectEntity
	DeleteObject(res *aws.DeleteObjectResponse) *DeleteObjectEntity
}

type AwsServerInputInterface interface {
	CreateBucket(res *aws.CreateBucketRequest) *CreateBucketEntity
	ListBuckets(res *aws.ListBucketsRequest) *ListBucketsEntity
	PutObject(res *aws.PutObjectRequest) *PutObjectEntity
	GetObject(res *aws.GetObjectRequest) *GetObjectEntity
	DeleteObject(res *aws.DeleteObjectRequest) *DeleteObjectEntity
}

type AwsServerOutputInterface interface {
	CreateBucket(req *CreateBucketEntity) *aws.CreateBucketResponse
	ListBuckets(req *ListBucketsEntity) *aws.ListBucketsResponse
	PutObject(req *PutObjectEntity) *aws.PutObjectResponse
	GetObject(req *GetObjectEntity) *aws.GetObjectResponse
	DeleteObject(req *DeleteObjectEntity) *aws.DeleteObjectResponse
}
