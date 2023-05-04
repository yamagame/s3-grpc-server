package storage

import (
	"context"
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/infra/storage"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/server/storage/gateway"
)

type serverGateway struct {
	createBucket gateway.CreateBucket
	listBuckets  gateway.ListBuckets
	putObject    gateway.PutObject
	getObject    gateway.GetObject
	deleteObject gateway.DeleteObject
	listObjects  gateway.ListObjects
}

type storageRepositoryServer struct {
	serverGateway
	service *storage.StorageRepository
	grpc_server.UnimplementedStorageRepositoryServer
}

func NewStorageRepositoryServer(service *storage.StorageRepository) *storageRepositoryServer {
	return &storageRepositoryServer{
		service: service,
	}
}

// CreateBucket implements storageRepositoryServer.CreateBucket
func (s *storageRepositoryServer) CreateBucket(ctx context.Context, in *grpc_server.CreateBucketRequest) (*grpc_server.CreateBucketResponse, error) {
	return s.createBucket.ToDomain(in, func(entity *model.CreateBucket) (*model.CreateBucket, error) {
		return s.service.CreateBucket(ctx, entity)
	})
}

// ListBuckets implements storageRepositoryServer.ListBuckets
func (s *storageRepositoryServer) ListBuckets(ctx context.Context, in *grpc_server.ListBucketsRequest) (*grpc_server.ListBucketsResponse, error) {
	return s.listBuckets.ToDomain(in, func(entity *model.ListBuckets) (*model.ListBuckets, error) {
		return s.service.ListBuckets(ctx, entity)
	})
}

// PutObject implements storageRepositoryServer.PutObject
func (s *storageRepositoryServer) PutObject(ctx context.Context, in *grpc_server.PutObjectRequest) (*grpc_server.PutObjectResponse, error) {
	return s.putObject.ToDomain(in, func(entity *model.PutObject) (*model.PutObject, error) {
		return s.service.PutObject(ctx, entity)
	})
}

// GetObject implements storageRepositoryServer.GetObject
func (s *storageRepositoryServer) GetObject(ctx context.Context, in *grpc_server.GetObjectRequest) (*grpc_server.GetObjectResponse, error) {
	return s.getObject.ToDomain(in, func(entity *model.GetObject) (*model.GetObject, error) {
		return s.service.GetObject(ctx, entity)
	})
}

// DeleteObject implements storageRepositoryServer.DeleteObject
func (s *storageRepositoryServer) DeleteObject(ctx context.Context, in *grpc_server.DeleteObjectRequest) (*grpc_server.DeleteObjectResponse, error) {
	return s.deleteObject.ToDomain(in, func(entity *model.DeleteObject) (*model.DeleteObject, error) {
		return s.service.DeleteObject(ctx, entity)
	})
}

// ListObjects implements storageRepositoryServer.ListObjects
func (s *storageRepositoryServer) ListObjects(ctx context.Context, in *grpc_server.ListObjectsRequest) (*grpc_server.ListObjectsResponse, error) {
	return s.listObjects.ToDomain(in, func(entity *model.ListObjects) (*model.ListObjects, error) {
		return s.service.ListObjects(ctx, entity)
	})
}
