package service

import (
	"context"
	"sample/s3-grpc-server/domain"
	server "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type storageClientDomain struct {
	createBucket domain.StorageCreateBucketClientData
	listBuckets  domain.StorageListBucketsClientData
	putObject    domain.StoragePutObjectClientData
	getObject    domain.StorageGetObjectClientData
	deleteObject domain.StorageDeleteObjectClientData
	listobjects  domain.StorageListObjectsClientData
}

type StorageService struct {
	scanner StorageScannerInterface
	domain  storageClientDomain
	client  server.StorageClient
}

func NewStorageService(aws server.StorageClient, scanner StorageScannerInterface) *StorageService {
	return &StorageService{
		scanner: scanner,
		domain:  storageClientDomain{},
		client:  aws,
	}
}

func (x *StorageService) CreateBucket(ctx context.Context) (*domain.CreateBucketEntity, error) {
	req := &server.CreateBucketRequest{}
	res, err := x.client.CreateBucket(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.createBucket.Output(res), nil
}

func (x *StorageService) ListBuckets(ctx context.Context) (*domain.ListBucketsEntity, error) {
	req := &server.ListBucketsRequest{}
	res, err := x.client.ListBuckets(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.listBuckets.Output(res), nil
}

func (x *StorageService) PutObject(ctx context.Context) (*domain.PutObjectEntity, error) {
	req := x.domain.putObject.Input(x.scanner.PutObject())
	res, err := x.client.PutObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.putObject.Output(res), nil
}

func (x *StorageService) GetObject(ctx context.Context) (*domain.GetObjectEntity, error) {
	req := x.domain.getObject.Input(x.scanner.GetObject())
	res, err := x.client.GetObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.getObject.Output(res), nil
}

func (x *StorageService) DeleteObject(ctx context.Context) (*domain.DeleteObjectEntity, error) {
	req := x.domain.deleteObject.Input(x.scanner.DeleteObject())
	res, err := x.client.DeleteObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.deleteObject.Output(res), nil
}

func (x *StorageService) ListObjects(ctx context.Context) (*domain.ListObjectsEntity, error) {
	req := x.domain.listobjects.Input(x.scanner.ListObjects())
	res, err := x.client.ListObjects(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.listobjects.Output(res), nil
}
