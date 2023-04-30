package internal

import (
	"context"
	"sample/s3-grpc-server/infra/storage"
	server "sample/s3-grpc-server/proto/grpc-server"
)

type storageClientDomain struct {
	createBucket storage.StorageCreateBucketClientData
	listBuckets  storage.StorageListBucketsClientData
	putObject    storage.StoragePutObjectClientData
	getObject    storage.StorageGetObjectClientData
	deleteObject storage.StorageDeleteObjectClientData
	listobjects  storage.StorageListObjectsClientData
}

type StorageClient struct {
	scanner StorageScannerInterface
	domain  storageClientDomain
	client  server.StorageClient
}

func NewStorageClient(aws server.StorageClient, scanner StorageScannerInterface) *StorageClient {
	return &StorageClient{
		scanner: scanner,
		domain:  storageClientDomain{},
		client:  aws,
	}
}

func (x *StorageClient) CreateBucket(ctx context.Context) (*storage.CreateBucketEntity, error) {
	req := &server.CreateBucketRequest{}
	res, err := x.client.CreateBucket(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.createBucket.Output(res), nil
}

func (x *StorageClient) ListBuckets(ctx context.Context) (*storage.ListBucketsEntity, error) {
	req := &server.ListBucketsRequest{}
	res, err := x.client.ListBuckets(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.listBuckets.Output(res), nil
}

func (x *StorageClient) PutObject(ctx context.Context) (*storage.PutObjectEntity, error) {
	req := x.domain.putObject.Input(x.scanner.PutObject())
	res, err := x.client.PutObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.putObject.Output(res), nil
}

func (x *StorageClient) GetObject(ctx context.Context) (*storage.GetObjectEntity, error) {
	req := x.domain.getObject.Input(x.scanner.GetObject())
	res, err := x.client.GetObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.getObject.Output(res), nil
}

func (x *StorageClient) DeleteObject(ctx context.Context) (*storage.DeleteObjectEntity, error) {
	req := x.domain.deleteObject.Input(x.scanner.DeleteObject())
	res, err := x.client.DeleteObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.deleteObject.Output(res), nil
}

func (x *StorageClient) ListObjects(ctx context.Context) (*storage.ListObjectsEntity, error) {
	req := x.domain.listobjects.Input(x.scanner.ListObjects())
	res, err := x.client.ListObjects(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.listobjects.Output(res), nil
}
