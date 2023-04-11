package s3

import (
	"context"
	"fmt"
	aws "sample/s3-grpc-server/grpc-server"

	"go.uber.org/zap"
)

// server is used to implement helloworld.GreeterServer.
type awsServer struct {
	aws *Client
	aws.UnimplementedAwsServer
}

func NewAWSServer(client *Client) *awsServer {
	return &awsServer{aws: client}
}

// CreateBucket implements awsServer.ListBuckets
func (s *awsServer) CreateBucket(ctx context.Context, in *aws.CreateBucketRequest) (*aws.CreateBucketResponse, error) {
	_, err := s.aws.CreateBucket(in.Name)
	if err != nil {
		fmt.Println(err)
		return &aws.CreateBucketResponse{Result: aws.Result_ERR}, nil
	}
	return &aws.CreateBucketResponse{Result: aws.Result_OK}, nil
}

// ListBuckets implements awsServer.ListBuckets
func (s *awsServer) ListBuckets(ctx context.Context, in *aws.ListBucketsRequest) (*aws.ListBucketsResponse, error) {
	list, err := s.aws.ListBuckets()
	if err != nil {
		fmt.Println(err)
		return &aws.ListBucketsResponse{Result: aws.Result_ERR}, nil
	}
	result := make([]*aws.ListBucketsResponseBucket, len(list.Buckets))
	for i, bucket := range list.Buckets {
		result[i] = &aws.ListBucketsResponseBucket{
			Name: *bucket.Name,
		}
	}
	return &aws.ListBucketsResponse{Result: aws.Result_OK, Buckets: result}, nil
}

// PutObject implements awsServer.PutObject
func (s *awsServer) PutObject(ctx context.Context, in *aws.PutObjectRequest) (*aws.PutObjectResponse, error) {
	zap.S().Infof("Received: %v", in.GetKey())
	_, err := s.aws.PutObjectWithString(in.GetKey(), in.GetContent())
	if err != nil {
		fmt.Println(err)
		return &aws.PutObjectResponse{Result: aws.Result_ERR}, nil
	}
	return &aws.PutObjectResponse{Result: aws.Result_OK}, nil
}

// GetObject implements awsServer.GetObject
func (s *awsServer) GetObject(ctx context.Context, in *aws.GetObjectRequest) (*aws.GetObjectResponse, error) {
	zap.S().Infof("Received: %v", in.GetKey())
	ret, err := s.aws.GetObjectWithString(in.GetKey())
	if err != nil {
		fmt.Println(err)
		return &aws.GetObjectResponse{Result: aws.Result_ERR}, nil
	}
	return &aws.GetObjectResponse{Result: aws.Result_OK, Content: ret}, nil
}

// DeleteObject implements awsServer.DeleteObject
func (s *awsServer) DeleteObject(ctx context.Context, in *aws.DeleteObjectRequest) (*aws.DeleteObjectResponse, error) {
	zap.S().Infof("Received: %v", in.GetKey())
	_, err := s.aws.DeleteObject(in.GetKey())
	if err != nil {
		fmt.Println(err)
		return &aws.DeleteObjectResponse{Result: aws.Result_ERR}, nil
	}
	return &aws.DeleteObjectResponse{Result: aws.Result_OK}, nil
}
