package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	aws "sample/s3-grpc-server/grpc-server"
	"sample/s3-grpc-server/s3"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type awsServer struct {
	aws *s3.Client
	aws.UnimplementedAwsServer
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

// ListUsers implements awsServer.ListUsers
func (s *awsServer) ListUsers(ctx context.Context, in *aws.ListUsersRequest) (*aws.ListUsersResponse, error) {
	list, err := s.aws.ListUsers()
	if err != nil {
		fmt.Println(err)
		return &aws.ListUsersResponse{Result: aws.Result_ERR}, nil
	}
	result := make([]*aws.ListUsersResponseUser, len(list.Users))
	for i, user := range list.Users {
		result[i] = &aws.ListUsersResponseUser{
			Name: *user.Username,
		}
	}
	return &aws.ListUsersResponse{Result: aws.Result_OK, Users: result}, nil
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	zap.ReplaceGlobals(logger)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		// log.Fatalf("failed to listen: %v", err)
		sugar.Infof("failed to listen: %v", err)
	}
	bucket := os.Getenv("S3_BUCKET_NAME")
	s := grpc.NewServer()
	client, err := s3.NewClient(context.Background(), bucket)
	if err != nil {
		// log.Fatalf("failed to listen: %v", err)
		sugar.Infof("failed to listen: %v", err)
	}
	aws.RegisterAwsServer(s, &awsServer{aws: client})
	// log.Printf("server listening at %v", lis.Addr())
	sugar.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		sugar.Infof("failed to serve: %v", err)
	}
}
