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

type tableRepository struct {
	clientDTO
	client server.RepositoryClient
}

func NewTableRepository(client server.RepositoryClient) *tableRepository {
	return &tableRepository{
		client: client,
	}
}

func (x *tableRepository) Create(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.create.Domain(table, func(req *server.CreateTableRequest) (*server.CreateTableResponse, error) {
		return x.client.CreateTable(ctx, req)
	})
}

func (x *tableRepository) Read(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.read.Domain(table, func(req *server.ReadTableRequest) (*server.ReadTableResponse, error) {
		return x.client.ReadTable(ctx, req)
	})
}

func (x *tableRepository) Update(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.update.Domain(table, func(req *server.UpdateTableRequest) (*server.UpdateTableResponse, error) {
		return x.client.UpdateTable(ctx, req)
	})
}

func (x *tableRepository) Delete(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.delete.Domain(table, func(req *server.DeleteTableRequest) (*server.DeleteTableResponse, error) {
		return x.client.DeleteTable(ctx, req)
	})
}
