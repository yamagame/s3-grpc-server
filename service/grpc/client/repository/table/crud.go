package table

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type clientDTO struct {
	create CreateTable
	read   ReadTable
	update UpdateTable
	delete DeleteTable
}

type TableRepository struct {
	clientDTO
	client server.TableRepositoryClient
}

func NewTableRepository(client server.TableRepositoryClient) *TableRepository {
	return &TableRepository{
		client: client,
	}
}

// Create implements tableRepository.Create
func (x *TableRepository) Create(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.create.Domain(table, func(req *server.CreateTableRequest) (*server.CreateTableResponse, error) {
		return x.client.Create(ctx, req)
	})
}

// Read implements tableRepository.Read
func (x *TableRepository) Read(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.read.Domain(table, func(req *server.ReadTableRequest) (*server.ReadTableResponse, error) {
		return x.client.Read(ctx, req)
	})
}

// Update implements tableRepository.Update
func (x *TableRepository) Update(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.update.Domain(table, func(req *server.UpdateTableRequest) (*server.UpdateTableResponse, error) {
		return x.client.Update(ctx, req)
	})
}

// Delete implements tableRepository.Delete
func (x *TableRepository) Delete(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.delete.Domain(table, func(req *server.DeleteTableRequest) (*server.DeleteTableResponse, error) {
		return x.client.Delete(ctx, req)
	})
}
