package storage

import (
	"context"
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/client/storage/gateway"
)

type clientGateway struct {
	createBucket gateway.CreateBucket
	listBuckets  gateway.ListBuckets
	putObject    gateway.PutObject
	getObject    gateway.GetObject
	deleteObject gateway.DeleteObject
	listobjects  gateway.ListObjects
}

type StorageClientRepository struct {
	clientGateway
	client grpc_server.StorageRepositoryClient
}

func NewStorageRepository(client grpc_server.StorageRepositoryClient) *StorageClientRepository {
	return &StorageClientRepository{
		client: client,
	}
}

func (x *StorageClientRepository) CreateBucket(ctx context.Context, object *model.CreateBucket) (*model.CreateBucket, error) {
	return x.createBucket.ToInfra(object, func(req *grpc_server.CreateBucketRequest) (*grpc_server.CreateBucketResponse, error) {
		return x.client.CreateBucket(ctx, req)
	})
}

func (x *StorageClientRepository) ListBuckets(ctx context.Context, object *model.ListBuckets) (*model.ListBuckets, error) {
	return x.listBuckets.ToInfra(object, func(req *grpc_server.ListBucketsRequest) (*grpc_server.ListBucketsResponse, error) {
		return x.client.ListBuckets(ctx, req)
	})
}

func (x *StorageClientRepository) PutObject(ctx context.Context, object *model.PutObject) (*model.PutObject, error) {
	return x.putObject.ToInfra(object, func(req *grpc_server.PutObjectRequest) (*grpc_server.PutObjectResponse, error) {
		return x.client.PutObject(ctx, req)
	})
}

func (x *StorageClientRepository) GetObject(ctx context.Context, object *model.GetObject) (*model.GetObject, error) {
	return x.getObject.ToInfra(object, func(req *grpc_server.GetObjectRequest) (*grpc_server.GetObjectResponse, error) {
		return x.client.GetObject(ctx, req)
	})
}

func (x *StorageClientRepository) DeleteObject(ctx context.Context, object *model.DeleteObject) (*model.DeleteObject, error) {
	return x.deleteObject.ToInfra(object, func(req *grpc_server.DeleteObjectRequest) (*grpc_server.DeleteObjectResponse, error) {
		return x.client.DeleteObject(ctx, req)
	})
}

func (x *StorageClientRepository) ListObjects(ctx context.Context, object *model.ListObjects) (*model.ListObjects, error) {
	return x.listobjects.ToInfra(object, func(req *grpc_server.ListObjectsRequest) (*grpc_server.ListObjectsResponse, error) {
		return x.client.ListObjects(ctx, req)
	})
}
