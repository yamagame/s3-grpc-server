package table

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/gateway"
)

type serverGateway struct {
	create CreateTable
	read   ReadTable
	update UpdateTable
	delete DeleteTable
	list   ListTable
}

type TableServerRepository struct {
	serverGateway
	repository table.RepositoryInterface
	grpc_server.UnimplementedTableRepositoryServer
}

func NewTableServerRepository(repository table.RepositoryInterface) *TableServerRepository {
	return &TableServerRepository{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *TableServerRepository) Create(ctx context.Context, in *grpc_server.CreateTableRequest) (*grpc_server.CreateTableResponse, error) {
	return gateway.ToDomain(in, s.create.Input, func(table *model.Table) (*model.Table, error) {
		return s.repository.Create(ctx, table)
	}, s.create.Output)
}

// Read implements RepositoryServer.Read
func (s *TableServerRepository) Read(ctx context.Context, in *grpc_server.ReadTableRequest) (*grpc_server.ReadTableResponse, error) {
	return gateway.ToDomain(in, s.read.Input, func(table *model.Table) (*model.Table, error) {
		return s.repository.Read(ctx, table)
	}, s.read.Output)
}

// Update implements RepositoryServer.Update
func (s *TableServerRepository) Update(ctx context.Context, in *grpc_server.UpdateTableRequest) (*grpc_server.UpdateTableResponse, error) {
	return gateway.ToDomain(in, s.update.Input, func(table *model.Table) (*model.Table, error) {
		return s.repository.Update(ctx, table)
	}, s.update.Output)
}

// Delete implements RepositoryServer.Delete
func (s *TableServerRepository) Delete(ctx context.Context, in *grpc_server.DeleteTableRequest) (*grpc_server.DeleteTableResponse, error) {
	return gateway.ToDomain(in, s.delete.Input, func(table *model.Table) (*model.Table, error) {
		return s.repository.Delete(ctx, table)
	}, s.delete.Output)
}

// List implements RepositoryServer.List
func (s *TableServerRepository) List(ctx context.Context, in *grpc_server.ListTableRequest) (*grpc_server.ListTableResponse, error) {
	return gateway.ToDomainList(in, s.list.Input, func(table *model.Table) ([]*model.Table, error) {
		return s.repository.List(ctx, table)
	}, s.list.Output)
}
