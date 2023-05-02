package storage

import (
	"context"
	"sample/s3-grpc-server/grpc_server/service/storage/dto"
	"sample/s3-grpc-server/infra/storage"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	createBucket dto.CreateBucket
	listBuckets  dto.ListBuckets
	putObject    dto.PutObject
	getObject    dto.GetObject
	deleteObject dto.DeleteObject
	listObjects  dto.ListObjects
}

type storageServer struct {
	serverDTO
	service *storage.StorageService
	server.UnimplementedStorageServer
}

func NewStorageServer(service *storage.StorageService) *storageServer {
	return &storageServer{
		service: service,
	}
}

// CreateBucket implements awsServer.CreateBucket
func (s *storageServer) CreateBucket(ctx context.Context, in *server.CreateBucketRequest) (*server.CreateBucketResponse, error) {
	entity := s.createBucket.Input(in)
	entity, err := s.service.CreateBucket(entity)
	if err != nil {
		return nil, err
	}
	return s.createBucket.Output(entity), nil
}

// ListBuckets implements awsServer.ListBuckets
func (s *storageServer) ListBuckets(ctx context.Context, in *server.ListBucketsRequest) (*server.ListBucketsResponse, error) {
	entity := s.listBuckets.Input(in)
	entity, err := s.service.ListBuckets(entity)
	if err != nil {
		return nil, err
	}
	return s.listBuckets.Output(entity), nil
}

// PutObject implements awsServer.PutObject
func (s *storageServer) PutObject(ctx context.Context, in *server.PutObjectRequest) (*server.PutObjectResponse, error) {
	entity := s.putObject.Input(in)
	entity, err := s.service.PutObject(entity)
	if err != nil {
		return nil, err
	}
	return s.putObject.Output(entity), nil
}

// GetObject implements awsServer.GetObject
func (s *storageServer) GetObject(ctx context.Context, in *server.GetObjectRequest) (*server.GetObjectResponse, error) {
	entity := s.getObject.Input(in)
	entity, err := s.service.GetObject(entity)
	if err != nil {
		return nil, err
	}
	return s.getObject.Output(entity), nil
}

// DeleteObject implements awsServer.DeleteObject
func (s *storageServer) DeleteObject(ctx context.Context, in *server.DeleteObjectRequest) (*server.DeleteObjectResponse, error) {
	entity := s.deleteObject.Input(in)
	entity, err := s.service.DeleteObject(entity)
	if err != nil {
		return nil, err
	}
	return s.deleteObject.Output(entity), nil
}

// ListObjects implements awsServer.ListObjects
func (s *storageServer) ListObjects(ctx context.Context, in *server.ListObjectsRequest) (*server.ListObjectsResponse, error) {
	entity := s.listObjects.Input(in)
	entity, err := s.service.ListObjects(entity)
	if err != nil {
		return nil, err
	}
	return s.listObjects.Output(entity), nil
}
