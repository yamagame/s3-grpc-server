package domain

import (
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type AwsCreateBucketServerData struct {
}

func (x *AwsCreateBucketServerData) Input(req *aws.CreateBucketRequest) *CreateBucketEntity {
	return &CreateBucketEntity{}
}

func (x *AwsCreateBucketServerData) Output(res *CreateBucketEntity) *aws.CreateBucketResponse {
	return &aws.CreateBucketResponse{
		Result: aws.Result(res.Result),
	}
}

type AwsListBucketsServerData struct {
}

func (x *AwsListBucketsServerData) Input(req *aws.ListBucketsRequest) *ListBucketsEntity {
	return &ListBucketsEntity{}
}

func (x *AwsListBucketsServerData) Output(res *ListBucketsEntity) *aws.ListBucketsResponse {
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

type AwsPutObjectServerData struct {
}

func (x *AwsPutObjectServerData) Input(req *aws.PutObjectRequest) *PutObjectEntity {
	return &PutObjectEntity{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *AwsPutObjectServerData) Output(res *PutObjectEntity) *aws.PutObjectResponse {
	return &aws.PutObjectResponse{
		Result: aws.Result(res.Result),
		Key:    res.Key,
	}
}

type AwsGetObjectServerData struct {
}

func (x *AwsGetObjectServerData) Input(req *aws.GetObjectRequest) *GetObjectEntity {
	return &GetObjectEntity{
		Key: req.Key,
	}
}

func (x *AwsGetObjectServerData) Output(res *GetObjectEntity) *aws.GetObjectResponse {
	return &aws.GetObjectResponse{
		Result:  aws.Result(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type AwsDeleteObjectServerData struct {
}

func (x *AwsDeleteObjectServerData) Input(req *aws.DeleteObjectRequest) *DeleteObjectEntity {
	return &DeleteObjectEntity{
		Key: req.Key,
	}
}

func (x *AwsDeleteObjectServerData) Output(res *DeleteObjectEntity) *aws.DeleteObjectResponse {
	return &aws.DeleteObjectResponse{
		Result: aws.Result(res.Result),
		Key:    res.Key,
	}
}

type AwsListObjectsServerData struct {
}

func (x *AwsListObjectsServerData) Input(req *aws.ListObjectsRequest) *ListObjectsEntity {
	return &ListObjectsEntity{
		Prefix: req.Prefix,
	}
}

func (x *AwsListObjectsServerData) Output(res *ListObjectsEntity) *aws.ListObjectsResponse {
	return &aws.ListObjectsResponse{
		Result: aws.Result(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
