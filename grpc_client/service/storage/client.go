package storage

import (
	"context"
	"sample/s3-grpc-server/infra/storage/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type storageClientMessage struct {
	createBucket StorageCreateBucketClientMessage
	listBuckets  StorageListBucketsClientMessage
	putObject    StoragePutObjectClientMessage
	getObject    StorageGetObjectClientMessage
	deleteObject StorageDeleteObjectClientMessage
	listobjects  StorageListObjectsClientMessage
}

type StorageClient struct {
	storageClientMessage
	scanner StorageScannerInterface
	client  server.StorageClient
}

func NewStorageClient(client server.StorageClient, scanner StorageScannerInterface) *StorageClient {
	return &StorageClient{
		scanner: scanner,
		client:  client,
	}
}

func (x *StorageClient) CreateBucket(ctx context.Context) (*model.CreateBucket, error) {
	req := &server.CreateBucketRequest{}
	res, err := x.client.CreateBucket(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createBucket.Output(res), nil
}

func (x *StorageClient) ListBuckets(ctx context.Context) (*model.ListBuckets, error) {
	req := &server.ListBucketsRequest{}
	res, err := x.client.ListBuckets(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.listBuckets.Output(res), nil
}

func (x *StorageClient) PutObject(ctx context.Context) (*model.PutObject, error) {
	req := x.putObject.Input(x.scanner.PutObject())
	res, err := x.client.PutObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.putObject.Output(res), nil
}

func (x *StorageClient) GetObject(ctx context.Context) (*model.GetObject, error) {
	req := x.getObject.Input(x.scanner.GetObject())
	res, err := x.client.GetObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.getObject.Output(res), nil
}

func (x *StorageClient) DeleteObject(ctx context.Context) (*model.DeleteObject, error) {
	req := x.deleteObject.Input(x.scanner.DeleteObject())
	res, err := x.client.DeleteObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteObject.Output(res), nil
}

func (x *StorageClient) ListObjects(ctx context.Context) (*model.ListObjects, error) {
	req := x.listobjects.Input(x.scanner.ListObjects())
	res, err := x.client.ListObjects(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.listobjects.Output(res), nil
}
