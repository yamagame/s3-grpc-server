package storage

import (
	"context"
	"sample/s3-grpc-server/infra/storage/model"
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

type storageRepository struct {
	clientGateway
	client grpc_server.StorageRepositoryClient
}

func NewRepository(client grpc_server.StorageRepositoryClient) *storageRepository {
	return &storageRepository{
		client: client,
	}
}

func (x *storageRepository) CreateBucket(ctx context.Context, object *model.CreateBucket) (*model.CreateBucket, error) {
	return x.createBucket.ToInfra(object, func(req *grpc_server.CreateBucketRequest) (*grpc_server.CreateBucketResponse, error) {
		return x.client.CreateBucket(ctx, req)
	})
}

func (x *storageRepository) ListBuckets(ctx context.Context, object *model.ListBuckets) (*model.ListBuckets, error) {
	return x.listBuckets.ToInfra(object, func(req *grpc_server.ListBucketsRequest) (*grpc_server.ListBucketsResponse, error) {
		return x.client.ListBuckets(ctx, req)
	})
}

func (x *storageRepository) PutObject(ctx context.Context, object *model.PutObject) (*model.PutObject, error) {
	return x.putObject.ToInfra(object, func(req *grpc_server.PutObjectRequest) (*grpc_server.PutObjectResponse, error) {
		return x.client.PutObject(ctx, req)
	})
}

func (x *storageRepository) GetObject(ctx context.Context, object *model.GetObject) (*model.GetObject, error) {
	return x.getObject.ToInfra(object, func(req *grpc_server.GetObjectRequest) (*grpc_server.GetObjectResponse, error) {
		return x.client.GetObject(ctx, req)
	})
}

func (x *storageRepository) DeleteObject(ctx context.Context, object *model.DeleteObject) (*model.DeleteObject, error) {
	return x.deleteObject.ToInfra(object, func(req *grpc_server.DeleteObjectRequest) (*grpc_server.DeleteObjectResponse, error) {
		return x.client.DeleteObject(ctx, req)
	})
}

func (x *storageRepository) ListObjects(ctx context.Context, object *model.ListObjects) (*model.ListObjects, error) {
	return x.listobjects.ToInfra(object, func(req *grpc_server.ListObjectsRequest) (*grpc_server.ListObjectsResponse, error) {
		return x.client.ListObjects(ctx, req)
	})
}
