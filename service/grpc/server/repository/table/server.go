package table

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/proto/grpc_server"
)

type serverGateway struct {
	create CreateTable
	read   ReadTable
	update UpdateTable
	delete DeleteTable
}

type TableRepositoryServer struct {
	serverGateway
	repository table.RepositoryInterface
	grpc_server.UnimplementedTableRepositoryServer
}

func NewTableRepositoryServer(repository table.RepositoryInterface) *TableRepositoryServer {
	return &TableRepositoryServer{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *TableRepositoryServer) Create(ctx context.Context, in *grpc_server.CreateTableRequest) (*grpc_server.CreateTableResponse, error) {
	return s.create.ToDomain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Create(ctx, table)
	})
}

// Read implements RepositoryServer.Read
func (s *TableRepositoryServer) Read(ctx context.Context, in *grpc_server.ReadTableRequest) (*grpc_server.ReadTableResponse, error) {
	return s.read.ToDomain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Read(ctx, table)
	})
}

// Update implements RepositoryServer.Update
func (s *TableRepositoryServer) Update(ctx context.Context, in *grpc_server.UpdateTableRequest) (*grpc_server.UpdateTableResponse, error) {
	return s.update.ToDomain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Update(ctx, table)
	})
}

// Delete implements RepositoryServer.Delete
func (s *TableRepositoryServer) Delete(ctx context.Context, in *grpc_server.DeleteTableRequest) (*grpc_server.DeleteTableResponse, error) {
	return s.delete.ToDomain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Delete(ctx, table)
	})
}
