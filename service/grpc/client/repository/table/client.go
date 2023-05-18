package table

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/gateway"
)

type clientGateway struct {
	create CreateTable
	read   ReadTable
	update UpdateTable
	delete DeleteTable
	list   ListTable
}

type TableClientRepository struct {
	clientGateway
	client grpc_server.TableRepositoryClient
}

func NewTableRepository(client grpc_server.TableRepositoryClient) *TableClientRepository {
	return &TableClientRepository{
		client: client,
	}
}

// Create implements tableRepository.Create
func (x *TableClientRepository) Create(ctx context.Context, table *model.Table) (*model.Table, error) {
	return gateway.ToInfra(table, x.create.Input, func(req *grpc_server.CreateTableRequest) (*grpc_server.CreateTableResponse, error) {
		return x.client.Create(ctx, req)
	}, x.create.Output)
}

// Read implements tableRepository.Read
func (x *TableClientRepository) Read(ctx context.Context, table *model.Table) (*model.Table, error) {
	return gateway.ToInfra(table, x.read.Input, func(req *grpc_server.ReadTableRequest) (*grpc_server.ReadTableResponse, error) {
		return x.client.Read(ctx, req)
	}, x.read.Output)
}

// Update implements tableRepository.Update
func (x *TableClientRepository) Update(ctx context.Context, table *model.Table) (*model.Table, error) {
	return gateway.ToInfra(table, x.update.Input, func(req *grpc_server.UpdateTableRequest) (*grpc_server.UpdateTableResponse, error) {
		return x.client.Update(ctx, req)
	}, x.update.Output)
}

// Delete implements tableRepository.Delete
func (x *TableClientRepository) Delete(ctx context.Context, table *model.Table) (*model.Table, error) {
	return gateway.ToInfra(table, x.delete.Input, func(req *grpc_server.DeleteTableRequest) (*grpc_server.DeleteTableResponse, error) {
		return x.client.Delete(ctx, req)
	}, x.delete.Output)
}

// List implements tableRepository.List
func (x *TableClientRepository) List(ctx context.Context, cell *model.Table) ([]*model.Table, error) {
	return gateway.ToInfraList(cell, x.list.Input, func(cell *grpc_server.ListTableRequest) (*grpc_server.ListTableResponse, error) {
		return x.client.List(ctx, cell)
	}, x.list.Output)
}
