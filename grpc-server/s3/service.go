package s3

import (
	"context"
	"fmt"
	"sample/s3-grpc-server/domain"
	"sample/s3-grpc-server/domain/server"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

// server is used to implement helloworld.GreeterServer.
type awsServer struct {
	aws    *Client
	input  domain.AwsServerInputInterface
	output domain.AwsServerOutputInterface
	aws.UnimplementedAwsServer
}

func NewAWSServer(client *Client) *awsServer {
	return &awsServer{
		aws:    client,
		input:  server.NewAwsServerInput(),
		output: server.NewAwsServerOutput(),
	}
}

// CreateBucket implements awsServer.ListBuckets
func (s *awsServer) CreateBucket(ctx context.Context, in *aws.CreateBucketRequest) (*aws.CreateBucketResponse, error) {
	entity := s.input.CreateBucket(in)
	_, err := s.aws.CreateBucket()
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
	}
	return s.output.CreateBucket(entity), nil
}

// ListBuckets implements awsServer.ListBuckets
func (s *awsServer) ListBuckets(ctx context.Context, in *aws.ListBucketsRequest) (*aws.ListBucketsResponse, error) {
	entity := s.input.ListBuckets(in)
	list, err := s.aws.ListBuckets()
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
		entity.Buckets = make([]domain.Bucket, len(list.Buckets))
		for i, bucket := range list.Buckets {
			entity.Buckets[i] = domain.Bucket{
				Name: *bucket.Name,
			}
		}
	}
	return s.output.ListBuckets(entity), nil
}

// PutObject implements awsServer.PutObject
func (s *awsServer) PutObject(ctx context.Context, in *aws.PutObjectRequest) (*aws.PutObjectResponse, error) {
	entity := s.input.PutObject(in)
	_, err := s.aws.PutObjectWithString(entity.Key, entity.Content)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
	}
	return s.output.PutObject(entity), nil
}

// GetObject implements awsServer.GetObject
func (s *awsServer) GetObject(ctx context.Context, in *aws.GetObjectRequest) (*aws.GetObjectResponse, error) {
	entity := s.input.GetObject(in)
	ret, err := s.aws.GetObjectWithString(entity.Key)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
		entity.Content = ret
	}
	return s.output.GetObject(entity), nil
}

// DeleteObject implements awsServer.DeleteObject
func (s *awsServer) DeleteObject(ctx context.Context, in *aws.DeleteObjectRequest) (*aws.DeleteObjectResponse, error) {
	entity := s.input.DeleteObject(in)
	_, err := s.aws.DeleteObject(entity.Key)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
	}
	return s.output.DeleteObject(entity), nil
}
