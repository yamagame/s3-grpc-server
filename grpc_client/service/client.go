package service

import (
	"context"
	"sample/s3-grpc-server/grpc_client/service/repository/cell"
	"sample/s3-grpc-server/grpc_client/service/repository/file"
	"sample/s3-grpc-server/grpc_client/service/repository/table"
	"sample/s3-grpc-server/grpc_client/service/storage"
	repository_model "sample/s3-grpc-server/infra/repository/model"
	storage_model "sample/s3-grpc-server/infra/storage/model"
)

type ClientService struct {
	storageClient *storage.StorageClient
	fileClient    *file.FileClient
	tableClient   *table.TableClient
	cellClient    *cell.CellClient
}

func NewClientService(
	storageClient *storage.StorageClient,
	fileClient *file.FileClient,
	tableClient *table.TableClient,
	cellClient *cell.CellClient,
) *ClientService {
	return &ClientService{
		storageClient: storageClient,
		fileClient:    fileClient,
		tableClient:   tableClient,
		cellClient:    cellClient,
	}
}

// File ----------------------------------------------------------------------

func (x *ClientService) CreateFile(ctx context.Context) (*repository_model.File, error) {
	return x.fileClient.Create(ctx)
}

func (x *ClientService) ReadFile(ctx context.Context) (*repository_model.File, error) {
	return x.fileClient.Read(ctx)
}

func (x *ClientService) UpdateFile(ctx context.Context) (*repository_model.File, error) {
	return x.fileClient.Update(ctx)
}

func (x *ClientService) DeleteFile(ctx context.Context) (*repository_model.File, error) {
	return x.fileClient.Delete(ctx)
}

// Table ----------------------------------------------------------------------

func (x *ClientService) CreateTable(ctx context.Context) (*repository_model.Table, error) {
	return x.tableClient.Create(ctx)
}

func (x *ClientService) ReadTable(ctx context.Context) (*repository_model.Table, error) {
	return x.tableClient.Read(ctx)
}

func (x *ClientService) UpdateTable(ctx context.Context) (*repository_model.Table, error) {
	return x.tableClient.Update(ctx)
}

func (x *ClientService) DeleteTable(ctx context.Context) (*repository_model.Table, error) {
	return x.tableClient.Delete(ctx)
}

// Cell ----------------------------------------------------------------------

func (x *ClientService) CreateCell(ctx context.Context) (*repository_model.Cell, error) {
	return x.cellClient.Create(ctx)
}

func (x *ClientService) ReadCell(ctx context.Context) (*repository_model.Cell, error) {
	return x.cellClient.Read(ctx)
}

func (x *ClientService) UpdateCell(ctx context.Context) (*repository_model.Cell, error) {
	return x.cellClient.Update(ctx)
}

func (x *ClientService) DeleteCell(ctx context.Context) (*repository_model.Cell, error) {
	return x.cellClient.Delete(ctx)
}

// Storage -----------------------------------------------------------------------

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
