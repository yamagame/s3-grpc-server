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

type client struct {
	storageClient *storage.StorageClient
	file          *file.CRUDClient
	table         *table.CRUDClient
	cell          *cell.CRUDClient
}

func NewClient(
	storageClient *storage.StorageClient,
	file *file.CRUDClient,
	table *table.CRUDClient,
	cell *cell.CRUDClient,
) *client {
	return &client{
		storageClient: storageClient,
		file:          file,
		table:         table,
		cell:          cell,
	}
}

// File ----------------------------------------------------------------------

func (x *client) CreateFile(ctx context.Context) (*repository_model.File, error) {
	return x.file.Create(ctx)
}

func (x *client) ReadFile(ctx context.Context) (*repository_model.File, error) {
	return x.file.Read(ctx)
}

func (x *client) UpdateFile(ctx context.Context) (*repository_model.File, error) {
	return x.file.Update(ctx)
}

func (x *client) DeleteFile(ctx context.Context) (*repository_model.File, error) {
	return x.file.Delete(ctx)
}

// Table ----------------------------------------------------------------------

func (x *client) CreateTable(ctx context.Context) (*repository_model.Table, error) {
	return x.table.Create(ctx)
}

func (x *client) ReadTable(ctx context.Context) (*repository_model.Table, error) {
	return x.table.Read(ctx)
}

func (x *client) UpdateTable(ctx context.Context) (*repository_model.Table, error) {
	return x.table.Update(ctx)
}

func (x *client) DeleteTable(ctx context.Context) (*repository_model.Table, error) {
	return x.table.Delete(ctx)
}

// Cell ----------------------------------------------------------------------

func (x *client) CreateCell(ctx context.Context) (*repository_model.Cell, error) {
	return x.cell.Create(ctx)
}

func (x *client) ReadCell(ctx context.Context) (*repository_model.Cell, error) {
	return x.cell.Read(ctx)
}

func (x *client) UpdateCell(ctx context.Context) (*repository_model.Cell, error) {
	return x.cell.Update(ctx)
}

func (x *client) DeleteCell(ctx context.Context) (*repository_model.Cell, error) {
	return x.cell.Delete(ctx)
}

// Storage -----------------------------------------------------------------------

func (x *client) CreateBucket(ctx context.Context) (*storage_model.CreateBucket, error) {
	return x.storageClient.CreateBucket(ctx)
}

func (x *client) ListBuckets(ctx context.Context) (*storage_model.ListBuckets, error) {
	return x.storageClient.ListBuckets(ctx)
}

func (x *client) PutObject(ctx context.Context) (*storage_model.PutObject, error) {
	return x.storageClient.PutObject(ctx)
}

func (x *client) GetObject(ctx context.Context) (*storage_model.GetObject, error) {
	return x.storageClient.GetObject(ctx)
}

func (x *client) DeleteObject(ctx context.Context) (*storage_model.DeleteObject, error) {
	return x.storageClient.DeleteObject(ctx)
}

func (x *client) ListObjects(ctx context.Context) (*storage_model.ListObjects, error) {
	return x.storageClient.ListObjects(ctx)
}
