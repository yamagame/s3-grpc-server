package storage

import (
	"context"
	"sample/s3-grpc-server/infra/storage"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type storageServerMessage struct {
	createBucket StorageCreateBucketServerMessage
	listBuckets  StorageListBucketsServerMessage
	putObject    StoragePutObjectServerMessage
	getObject    StorageGetObjectServerMessage
	deleteObject StorageDeleteObjectServerMessage
	listObjects  StorageListObjectsServerMessage
}

type storageServer struct {
	service *storage.StorageService
	message *storageServerMessage
	server.UnimplementedStorageServer
}

func NewStorageServer(service *storage.StorageService) *storageServer {
	return &storageServer{
		service: service,
		message: &storageServerMessage{},
	}
}

// CreateBucket implements awsServer.CreateBucket
func (s *storageServer) CreateBucket(ctx context.Context, in *server.CreateBucketRequest) (*server.CreateBucketResponse, error) {
	entity := s.message.createBucket.Input(in)
	entity, err := s.service.CreateBucket(entity)
	if err != nil {
		return nil, err
	}
	return s.message.createBucket.Output(entity), nil
}

// ListBuckets implements awsServer.ListBuckets
func (s *storageServer) ListBuckets(ctx context.Context, in *server.ListBucketsRequest) (*server.ListBucketsResponse, error) {
	entity := s.message.listBuckets.Input(in)
	entity, err := s.service.ListBuckets(entity)
	if err != nil {
		return nil, err
	}
	return s.message.listBuckets.Output(entity), nil
}

// PutObject implements awsServer.PutObject
func (s *storageServer) PutObject(ctx context.Context, in *server.PutObjectRequest) (*server.PutObjectResponse, error) {
	entity := s.message.putObject.Input(in)
	entity, err := s.service.PutObject(entity)
	if err != nil {
		return nil, err
	}
	return s.message.putObject.Output(entity), nil
}

// GetObject implements awsServer.GetObject
func (s *storageServer) GetObject(ctx context.Context, in *server.GetObjectRequest) (*server.GetObjectResponse, error) {
	entity := s.message.getObject.Input(in)
	entity, err := s.service.GetObject(entity)
	if err != nil {
		return nil, err
	}
	return s.message.getObject.Output(entity), nil
}

// DeleteObject implements awsServer.DeleteObject
func (s *storageServer) DeleteObject(ctx context.Context, in *server.DeleteObjectRequest) (*server.DeleteObjectResponse, error) {
	entity := s.message.deleteObject.Input(in)
	entity, err := s.service.DeleteObject(entity)
	if err != nil {
		return nil, err
	}
	return s.message.deleteObject.Output(entity), nil
}

// ListObjects implements awsServer.ListObjects
func (s *storageServer) ListObjects(ctx context.Context, in *server.ListObjectsRequest) (*server.ListObjectsResponse, error) {
	entity := s.message.listObjects.Input(in)
	entity, err := s.service.ListObjects(entity)
	if err != nil {
		return nil, err
	}
	return s.message.listObjects.Output(entity), nil
}
