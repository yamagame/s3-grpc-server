package domain

import (
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type AwsCreateBucketClientData struct {
}

func (x *AwsCreateBucketClientData) Input(_ *CreateBucketEntity) *aws.CreateBucketRequest {
	return &aws.CreateBucketRequest{}
}

func (x *AwsCreateBucketClientData) Output(res *aws.CreateBucketResponse) *CreateBucketEntity {
	return &CreateBucketEntity{
		Result: AwsResult(res.Result),
	}
}

type AwsListBucketsClientData struct {
}

func (x *AwsListBucketsClientData) Input(_ *ListBucketsEntity) *aws.ListBucketsRequest {
	return &aws.ListBucketsRequest{}
}

func (x *AwsListBucketsClientData) Output(res *aws.ListBucketsResponse) *ListBucketsEntity {
	buckets := make([]Bucket, len(res.Buckets))
	for i, v := range res.Buckets {
		buckets[i].Name = v.Name
	}
	return &ListBucketsEntity{
		Result:  AwsResult(res.Result),
		Buckets: buckets,
	}
}

type AwsPutObjectClientData struct {
}

func (x *AwsPutObjectClientData) Input(req *PutObjectEntity) *aws.PutObjectRequest {
	return &aws.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *AwsPutObjectClientData) Output(res *aws.PutObjectResponse) *PutObjectEntity {
	return &PutObjectEntity{
		Result: AwsResult(res.Result),
	}
}

type AwsGetObjectClientData struct {
}

func (x *AwsGetObjectClientData) Input(req *GetObjectEntity) *aws.GetObjectRequest {
	return &aws.GetObjectRequest{
		Key: req.Key,
	}
}

func (x *AwsGetObjectClientData) Output(res *aws.GetObjectResponse) *GetObjectEntity {
	return &GetObjectEntity{
		Result:  AwsResult(res.Result),
		Content: res.Content,
	}
}

type AwsDeleteObjectClientData struct {
}

func (x *AwsDeleteObjectClientData) Input(req *DeleteObjectEntity) *aws.DeleteObjectRequest {
	return &aws.DeleteObjectRequest{
		Key: req.Key,
	}
}

func (x *AwsDeleteObjectClientData) Output(res *aws.DeleteObjectResponse) *DeleteObjectEntity {
	return &DeleteObjectEntity{
		Result: AwsResult(res.Result),
	}
}

type AwsListObjectsClientData struct {
}

func (x *AwsListObjectsClientData) Input(req *ListObjectsEntity) *aws.ListObjectsRequest {
	return &aws.ListObjectsRequest{
		Prefix: req.Prefix,
	}
}

func (x *AwsListObjectsClientData) Output(res *aws.ListObjectsResponse) *ListObjectsEntity {
	return &ListObjectsEntity{
		Result: AwsResult(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
