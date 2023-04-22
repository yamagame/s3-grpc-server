package storage

import (
	"context"
	"sample/s3-grpc-server/domain"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type awsServerDomain struct {
	createBucket domain.AwsCreateBucketServerData
	listBuckets  domain.AwsListBucketsServerData
	putObject    domain.AwsPutObjectServerData
	getObject    domain.AwsGetObjectServerData
	deleteObject domain.AwsDeleteObjectServerData
	listObjects  domain.AwsListObjectsServerData
}

type awsServer struct {
	service *awsService
	domain  *awsServerDomain
	aws.UnimplementedAwsServer
}

func NewAWSServer(client *Client) *awsServer {
	return &awsServer{
		service: NewAWSService(client),
		domain:  &awsServerDomain{},
	}
}

// CreateBucket implements awsServer.ListBuckets
func (s *awsServer) CreateBucket(ctx context.Context, in *aws.CreateBucketRequest) (*aws.CreateBucketResponse, error) {
	entity := s.domain.createBucket.Input(in)
	entity, err := s.service.CreateBucket(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.createBucket.Output(entity), nil
}

// ListBuckets implements awsServer.ListBuckets
func (s *awsServer) ListBuckets(ctx context.Context, in *aws.ListBucketsRequest) (*aws.ListBucketsResponse, error) {
	entity := s.domain.listBuckets.Input(in)
	entity, err := s.service.ListBuckets(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.listBuckets.Output(entity), nil
}

// PutObject implements awsServer.PutObject
func (s *awsServer) PutObject(ctx context.Context, in *aws.PutObjectRequest) (*aws.PutObjectResponse, error) {
	entity := s.domain.putObject.Input(in)
	entity, err := s.service.PutObject(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.putObject.Output(entity), nil
}

// GetObject implements awsServer.GetObject
func (s *awsServer) GetObject(ctx context.Context, in *aws.GetObjectRequest) (*aws.GetObjectResponse, error) {
	entity := s.domain.getObject.Input(in)
	entity, err := s.service.GetObject(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.getObject.Output(entity), nil
}

// DeleteObject implements awsServer.DeleteObject
func (s *awsServer) DeleteObject(ctx context.Context, in *aws.DeleteObjectRequest) (*aws.DeleteObjectResponse, error) {
	entity := s.domain.deleteObject.Input(in)
	entity, err := s.service.DeleteObject(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.deleteObject.Output(entity), nil
}

// ListObjects implements awsServer.ListObjects
func (s *awsServer) ListObjects(ctx context.Context, in *aws.ListObjectsRequest) (*aws.ListObjectsResponse, error) {
	entity := s.domain.listObjects.Input(in)
	entity, err := s.service.ListObjects(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.listObjects.Output(entity), nil
}
