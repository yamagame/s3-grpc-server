package storage

import (
	"context"
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/infra/storage"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/gateway"
	gw "sample/s3-grpc-server/service/grpc/server/storage/gateway"
)

type serverGateway struct {
	createBucket gw.CreateBucket
	listBuckets  gw.ListBuckets
	putObject    gw.PutObject
	getObject    gw.GetObject
	deleteObject gw.DeleteObject
	listObjects  gw.ListObjects
}

type StorageServerRepository struct {
	serverGateway
	service *storage.StorageRepository
	grpc_server.UnimplementedStorageRepositoryServer
}

func NewStorageServerRepository(service *storage.StorageRepository) *StorageServerRepository {
	return &StorageServerRepository{
		service: service,
	}
}

// CreateBucket implements storageRepositoryServer.CreateBucket
func (s *StorageServerRepository) CreateBucket(ctx context.Context, in *grpc_server.CreateBucketRequest) (*grpc_server.CreateBucketResponse, error) {
	return gateway.ToDomain(in, s.createBucket.Input, func(entity *model.CreateBucket) (*model.CreateBucket, error) {
		return s.service.CreateBucket(ctx, entity)
	}, s.createBucket.Output)
}

// ListBuckets implements storageRepositoryServer.ListBuckets
func (s *StorageServerRepository) ListBuckets(ctx context.Context, in *grpc_server.ListBucketsRequest) (*grpc_server.ListBucketsResponse, error) {
	return gateway.ToDomain(in, s.listBuckets.Input, func(entity *model.ListBuckets) (*model.ListBuckets, error) {
		return s.service.ListBuckets(ctx, entity)
	}, s.listBuckets.Output)
}

// PutObject implements storageRepositoryServer.PutObject
func (s *StorageServerRepository) PutObject(ctx context.Context, in *grpc_server.PutObjectRequest) (*grpc_server.PutObjectResponse, error) {
	return gateway.ToDomain(in, s.putObject.Input, func(entity *model.PutObject) (*model.PutObject, error) {
		return s.service.PutObject(ctx, entity)
	}, s.putObject.Output)
}

// GetObject implements storageRepositoryServer.GetObject
func (s *StorageServerRepository) GetObject(ctx context.Context, in *grpc_server.GetObjectRequest) (*grpc_server.GetObjectResponse, error) {
	return gateway.ToDomain(in, s.getObject.Input, func(entity *model.GetObject) (*model.GetObject, error) {
		return s.service.GetObject(ctx, entity)
	}, s.getObject.Output)
}

// DeleteObject implements storageRepositoryServer.DeleteObject
func (s *StorageServerRepository) DeleteObject(ctx context.Context, in *grpc_server.DeleteObjectRequest) (*grpc_server.DeleteObjectResponse, error) {
	return gateway.ToDomain(in, s.deleteObject.Input, func(entity *model.DeleteObject) (*model.DeleteObject, error) {
		return s.service.DeleteObject(ctx, entity)
	}, s.deleteObject.Output)
}

// ListObjects implements storageRepositoryServer.ListObjects
func (s *StorageServerRepository) ListObjects(ctx context.Context, in *grpc_server.ListObjectsRequest) (*grpc_server.ListObjectsResponse, error) {
	return gateway.ToDomain(in, s.listObjects.Input, func(entity *model.ListObjects) (*model.ListObjects, error) {
		return s.service.ListObjects(ctx, entity)
	}, s.listObjects.Output)
}
