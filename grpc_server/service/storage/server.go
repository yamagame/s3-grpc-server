package storage

import (
	"context"
	"sample/s3-grpc-server/infra/storage"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type storageServerDomain struct {
	createBucket storage.StorageCreateBucketServerData
	listBuckets  storage.StorageListBucketsServerData
	putObject    storage.StoragePutObjectServerData
	getObject    storage.StorageGetObjectServerData
	deleteObject storage.StorageDeleteObjectServerData
	listObjects  storage.StorageListObjectsServerData
}

type storageServer struct {
	service *storageService
	domain  *storageServerDomain
	server.UnimplementedStorageServer
}

func NewStorageServer(client storage.ClientInterface) *storageServer {
	return &storageServer{
		service: NewStorageService(client),
		domain:  &storageServerDomain{},
	}
}

// CreateBucket implements awsServer.CreateBucket
func (s *storageServer) CreateBucket(ctx context.Context, in *server.CreateBucketRequest) (*server.CreateBucketResponse, error) {
	entity := s.domain.createBucket.Input(in)
	entity, err := s.service.CreateBucket(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.createBucket.Output(entity), nil
}

// ListBuckets implements awsServer.ListBuckets
func (s *storageServer) ListBuckets(ctx context.Context, in *server.ListBucketsRequest) (*server.ListBucketsResponse, error) {
	entity := s.domain.listBuckets.Input(in)
	entity, err := s.service.ListBuckets(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.listBuckets.Output(entity), nil
}

// PutObject implements awsServer.PutObject
func (s *storageServer) PutObject(ctx context.Context, in *server.PutObjectRequest) (*server.PutObjectResponse, error) {
	entity := s.domain.putObject.Input(in)
	entity, err := s.service.PutObject(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.putObject.Output(entity), nil
}

// GetObject implements awsServer.GetObject
func (s *storageServer) GetObject(ctx context.Context, in *server.GetObjectRequest) (*server.GetObjectResponse, error) {
	entity := s.domain.getObject.Input(in)
	entity, err := s.service.GetObject(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.getObject.Output(entity), nil
}

// DeleteObject implements awsServer.DeleteObject
func (s *storageServer) DeleteObject(ctx context.Context, in *server.DeleteObjectRequest) (*server.DeleteObjectResponse, error) {
	entity := s.domain.deleteObject.Input(in)
	entity, err := s.service.DeleteObject(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.deleteObject.Output(entity), nil
}

// ListObjects implements awsServer.ListObjects
func (s *storageServer) ListObjects(ctx context.Context, in *server.ListObjectsRequest) (*server.ListObjectsResponse, error) {
	entity := s.domain.listObjects.Input(in)
	entity, err := s.service.ListObjects(entity)
	if err != nil {
		return nil, err
	}
	return s.domain.listObjects.Output(entity), nil
}

// CreateFileInfo implements awsServer.CreateFileInfo
func (s *storageServer) CreateFileInfo(ctx context.Context, in *server.CreateFileInfoRequest) (*server.CreateFileInfoResponse, error) {
	return &server.CreateFileInfoResponse{}, nil
}

// ReadFileInfo implements awsServer.ReadFileInfo
func (s *storageServer) ReadFileInfo(ctx context.Context, in *server.ReadFileInfoRequest) (*server.ReadFileInfoResponse, error) {
	return &server.ReadFileInfoResponse{}, nil
}

// UpdateFileInfo implements awsServer.UpdateFileInfo
func (s *storageServer) UpdateFileInfo(ctx context.Context, in *server.UpdateFileInfoRequest) (*server.UpdateFileInfoResponse, error) {
	return &server.UpdateFileInfoResponse{}, nil
}

// DeleteFileInfo implements awsServer.DeleteFileInfo
func (s *storageServer) DeleteFileInfo(ctx context.Context, in *server.DeleteFileInfoRequest) (*server.DeleteFileInfoResponse, error) {
	return &server.DeleteFileInfoResponse{}, nil
}
