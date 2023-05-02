package table

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/infra/repository/table"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	create CreateTable
	read   ReadTable
	update UpdateTable
	delete DeleteTable
}

type CRUDService struct {
	serverDTO
	repository table.RepositoryInterface
}

func NewCRUDService(repository table.RepositoryInterface) *CRUDService {
	return &CRUDService{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *CRUDService) Create(ctx context.Context, in *server.CreateTableRequest) (*server.CreateTableResponse, error) {
	return s.create.Domain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Create(table)
	})
}

// Read implements RepositoryServer.Read
func (s *CRUDService) Read(ctx context.Context, in *server.ReadTableRequest) (*server.ReadTableResponse, error) {
	return s.read.Domain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Read(table)
	})
}

// Update implements RepositoryServer.Update
func (s *CRUDService) Update(ctx context.Context, in *server.UpdateTableRequest) (*server.UpdateTableResponse, error) {
	return s.update.Domain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Update(table)
	})
}

// Delete implements RepositoryServer.Delete
func (s *CRUDService) Delete(ctx context.Context, in *server.DeleteTableRequest) (*server.DeleteTableResponse, error) {
	return s.delete.Domain(in, func(table *model.Table) (*model.Table, error) {
		return s.repository.Delete(table)
	})
}
