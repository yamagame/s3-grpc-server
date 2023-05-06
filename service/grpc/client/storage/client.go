package storage

import (
	"context"
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/proto/grpc_server"
	gw "sample/s3-grpc-server/service/grpc/client/storage/gateway"
	"sample/s3-grpc-server/service/grpc/gateway"
)

type clientGateway struct {
	createBucket gw.CreateBucket
	listBuckets  gw.ListBuckets
	putObject    gw.PutObject
	getObject    gw.GetObject
	deleteObject gw.DeleteObject
	listobjects  gw.ListObjects
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
	return gateway.ToInfra(object, x.createBucket.Input, func(req *grpc_server.CreateBucketRequest) (*grpc_server.CreateBucketResponse, error) {
		return x.client.CreateBucket(ctx, req)
	}, x.createBucket.Output)
}

func (x *StorageClientRepository) ListBuckets(ctx context.Context, object *model.ListBuckets) (*model.ListBuckets, error) {
	return gateway.ToInfra(object, x.listBuckets.Input, func(req *grpc_server.ListBucketsRequest) (*grpc_server.ListBucketsResponse, error) {
		return x.client.ListBuckets(ctx, req)
	}, x.listBuckets.Output)
}

func (x *StorageClientRepository) PutObject(ctx context.Context, object *model.PutObject) (*model.PutObject, error) {
	return gateway.ToInfra(object, x.putObject.Input, func(req *grpc_server.PutObjectRequest) (*grpc_server.PutObjectResponse, error) {
		return x.client.PutObject(ctx, req)
	}, x.putObject.Output)
}

func (x *StorageClientRepository) GetObject(ctx context.Context, object *model.GetObject) (*model.GetObject, error) {
	return gateway.ToInfra(object, x.getObject.Input, func(req *grpc_server.GetObjectRequest) (*grpc_server.GetObjectResponse, error) {
		return x.client.GetObject(ctx, req)
	}, x.getObject.Output)
}

func (x *StorageClientRepository) DeleteObject(ctx context.Context, object *model.DeleteObject) (*model.DeleteObject, error) {
	return gateway.ToInfra(object, x.deleteObject.Input, func(req *grpc_server.DeleteObjectRequest) (*grpc_server.DeleteObjectResponse, error) {
		return x.client.DeleteObject(ctx, req)
	}, x.deleteObject.Output)
}

func (x *StorageClientRepository) ListObjects(ctx context.Context, object *model.ListObjects) (*model.ListObjects, error) {
	return gateway.ToInfra(object, x.listobjects.Input, func(req *grpc_server.ListObjectsRequest) (*grpc_server.ListObjectsResponse, error) {
		return x.client.ListObjects(ctx, req)
	}, x.listobjects.Output)
}
