package storage

import (
	"context"
	"sample/s3-grpc-server/grpc_client/service/storage/dto"
	"sample/s3-grpc-server/infra/storage/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type clientDTO struct {
	createBucket dto.CreateBucketDTO
	listBuckets  dto.ListBucketsDTO
	putObject    dto.PutObjectDTO
	getObject    dto.GetObjectDTO
	deleteObject dto.DeleteObjectDTO
	listobjects  dto.ListObjectsDTO
}

type storageRepository struct {
	clientDTO
	client server.StorageClient
}

func NewStorageRepository(client server.StorageClient) *storageRepository {
	return &storageRepository{
		client: client,
	}
}

func (x *storageRepository) CreateBucket(ctx context.Context, object *model.CreateBucket) (*model.CreateBucket, error) {
	req := x.createBucket.Input(object)
	res, err := x.client.CreateBucket(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createBucket.Output(res), nil
}

func (x *storageRepository) ListBuckets(ctx context.Context, object *model.ListBuckets) (*model.ListBuckets, error) {
	req := x.listBuckets.Input(object)
	res, err := x.client.ListBuckets(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.listBuckets.Output(res), nil
}

func (x *storageRepository) PutObject(ctx context.Context, object *model.PutObject) (*model.PutObject, error) {
	req := x.putObject.Input(object)
	res, err := x.client.PutObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.putObject.Output(res), nil
}

func (x *storageRepository) GetObject(ctx context.Context, object *model.GetObject) (*model.GetObject, error) {
	req := x.getObject.Input(object)
	res, err := x.client.GetObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.getObject.Output(res), nil
}

func (x *storageRepository) DeleteObject(ctx context.Context, object *model.DeleteObject) (*model.DeleteObject, error) {
	req := x.deleteObject.Input(object)
	res, err := x.client.DeleteObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteObject.Output(res), nil
}

func (x *storageRepository) ListObjects(ctx context.Context, object *model.ListObjects) (*model.ListObjects, error) {
	req := x.listobjects.Input(object)
	res, err := x.client.ListObjects(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.listobjects.Output(res), nil
}
