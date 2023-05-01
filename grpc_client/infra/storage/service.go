package storage

import (
	"context"
	"sample/s3-grpc-server/infra/storage"
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
	scanner StorageScannerInterface
	message storageClientMessage
	client  server.StorageClient
}

func NewStorageClient(client server.StorageClient, scanner StorageScannerInterface) *StorageClient {
	return &StorageClient{
		scanner: scanner,
		message: storageClientMessage{},
		client:  client,
	}
}

func (x *StorageClient) CreateBucket(ctx context.Context) (*storage.CreateBucketEntity, error) {
	req := &server.CreateBucketRequest{}
	res, err := x.client.CreateBucket(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.message.createBucket.Output(res), nil
}

func (x *StorageClient) ListBuckets(ctx context.Context) (*storage.ListBucketsEntity, error) {
	req := &server.ListBucketsRequest{}
	res, err := x.client.ListBuckets(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.message.listBuckets.Output(res), nil
}

func (x *StorageClient) PutObject(ctx context.Context) (*storage.PutObjectEntity, error) {
	req := x.message.putObject.Input(x.scanner.PutObject())
	res, err := x.client.PutObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.message.putObject.Output(res), nil
}

func (x *StorageClient) GetObject(ctx context.Context) (*storage.GetObjectEntity, error) {
	req := x.message.getObject.Input(x.scanner.GetObject())
	res, err := x.client.GetObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.message.getObject.Output(res), nil
}

func (x *StorageClient) DeleteObject(ctx context.Context) (*storage.DeleteObjectEntity, error) {
	req := x.message.deleteObject.Input(x.scanner.DeleteObject())
	res, err := x.client.DeleteObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.message.deleteObject.Output(res), nil
}

func (x *StorageClient) ListObjects(ctx context.Context) (*storage.ListObjectsEntity, error) {
	req := x.message.listobjects.Input(x.scanner.ListObjects())
	res, err := x.client.ListObjects(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.message.listobjects.Output(res), nil
}
