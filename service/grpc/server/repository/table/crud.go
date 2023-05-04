package table

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/proto/grpc_server"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	create CreateTable
	read   ReadTable
	update UpdateTable
	delete DeleteTable
}

type TableRepository struct {
	serverDTO
	repository table.RepositoryInterface
	grpc_server.UnimplementedTableRepositoryServer
}

func NewTableRepository(repository table.RepositoryInterface) *TableRepository {
	return &TableRepository{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *TableRepository) Create(ctx context.Context, in *server.CreateTableRequest) (*server.CreateTableResponse, error) {
	return s.create.Domain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Create(ctx, table)
	})
}

// Read implements RepositoryServer.Read
func (s *TableRepository) Read(ctx context.Context, in *server.ReadTableRequest) (*server.ReadTableResponse, error) {
	return s.read.Domain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Read(ctx, table)
	})
}

// Update implements RepositoryServer.Update
func (s *TableRepository) Update(ctx context.Context, in *server.UpdateTableRequest) (*server.UpdateTableResponse, error) {
	return s.update.Domain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Update(ctx, table)
	})
}

// Delete implements RepositoryServer.Delete
func (s *TableRepository) Delete(ctx context.Context, in *server.DeleteTableRequest) (*server.DeleteTableResponse, error) {
	return s.delete.Domain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Delete(ctx, table)
	})
}
