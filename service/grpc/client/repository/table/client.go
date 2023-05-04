package table

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type clientGateway struct {
	create CreateTable
	read   ReadTable
	update UpdateTable
	delete DeleteTable
}

type TableRepository struct {
	clientGateway
	client grpc_server.TableRepositoryClient
}

func NewTableRepository(client grpc_server.TableRepositoryClient) *TableRepository {
	return &TableRepository{
		client: client,
	}
}

// Create implements tableRepository.Create
func (x *TableRepository) Create(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.create.ToInfra(table, func(req *grpc_server.CreateTableRequest) (*grpc_server.CreateTableResponse, error) {
		return x.client.Create(ctx, req)
	})
}

// Read implements tableRepository.Read
func (x *TableRepository) Read(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.read.ToInfra(table, func(req *grpc_server.ReadTableRequest) (*grpc_server.ReadTableResponse, error) {
		return x.client.Read(ctx, req)
	})
}

// Update implements tableRepository.Update
func (x *TableRepository) Update(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.update.ToInfra(table, func(req *grpc_server.UpdateTableRequest) (*grpc_server.UpdateTableResponse, error) {
		return x.client.Update(ctx, req)
	})
}

// Delete implements tableRepository.Delete
func (x *TableRepository) Delete(ctx context.Context, table *model.Table) (*model.Table, error) {
	return x.delete.ToInfra(table, func(req *grpc_server.DeleteTableRequest) (*grpc_server.DeleteTableResponse, error) {
		return x.client.Delete(ctx, req)
	})
}
