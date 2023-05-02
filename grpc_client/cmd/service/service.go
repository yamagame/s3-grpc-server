package service

import (
	"context"
	repository_service "sample/s3-grpc-server/grpc_client/service/repository"
	storage_service "sample/s3-grpc-server/grpc_client/service/storage"
	repository_model "sample/s3-grpc-server/infra/repository/model"
	storage_model "sample/s3-grpc-server/infra/storage/model"
)

type ClientService struct {
	storageClient    *storage_service.StorageClient
	repositoryClient *repository_service.RepositoryClient
}

func NewClientService(storageClient *storage_service.StorageClient, repositoryClient *repository_service.RepositoryClient) *ClientService {
	return &ClientService{
		storageClient:    storageClient,
		repositoryClient: repositoryClient,
	}
}

func (x *ClientService) CreateFileInfo(ctx context.Context) (*repository_model.FileInfo, error) {
	return x.repositoryClient.CreateFileInfo(ctx)
}

func (x *ClientService) ReadFileInfo(ctx context.Context) (*repository_model.FileInfo, error) {
	return x.repositoryClient.ReadFileInfo(ctx)
}

func (x *ClientService) UpdateFileInfo(ctx context.Context) (*repository_model.FileInfo, error) {
	return x.repositoryClient.UpdateFileInfo(ctx)
}

func (x *ClientService) DeleteFileInfo(ctx context.Context) (*repository_model.FileInfo, error) {
	return x.repositoryClient.DeleteFileInfo(ctx)
}

func (x *ClientService) CreateBucket(ctx context.Context) (*storage_model.CreateBucket, error) {
	return x.storageClient.CreateBucket(ctx)
}

func (x *ClientService) ListBuckets(ctx context.Context) (*storage_model.ListBuckets, error) {
	return x.storageClient.ListBuckets(ctx)
}

func (x *ClientService) PutObject(ctx context.Context) (*storage_model.PutObject, error) {
	return x.storageClient.PutObject(ctx)
}

func (x *ClientService) GetObject(ctx context.Context) (*storage_model.GetObject, error) {
	return x.storageClient.GetObject(ctx)
}

func (x *ClientService) DeleteObject(ctx context.Context) (*storage_model.DeleteObject, error) {
	return x.storageClient.DeleteObject(ctx)
}

func (x *ClientService) ListObjects(ctx context.Context) (*storage_model.ListObjects, error) {
	return x.storageClient.ListObjects(ctx)
}
