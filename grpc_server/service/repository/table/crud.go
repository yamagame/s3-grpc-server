package table

import (
	"context"
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
	service table.RepositoryInterface
	server.UnimplementedRepositoryServer
}

func NewCRUDService(service table.RepositoryInterface) *CRUDService {
	return &CRUDService{
		service: service,
	}
}

// Create implements RepositoryServer.Create
func (s *CRUDService) Create(ctx context.Context, in *server.CreateTableRequest) (*server.CreateTableResponse, error) {
	file := s.create.Input(in)
	file, err := s.service.Create(file)
	if err != nil {
		return nil, err
	}
	return s.create.Output(file), nil
}

// Read implements RepositoryServer.Read
func (s *CRUDService) Read(ctx context.Context, in *server.ReadTableRequest) (*server.ReadTableResponse, error) {
	file := s.read.Input(in)
	file, err := s.service.Read(file)
	if err != nil {
		return nil, err
	}
	return s.read.Output(file), nil
}

// Update implements RepositoryServer.Update
func (s *CRUDService) Update(ctx context.Context, in *server.UpdateTableRequest) (*server.UpdateTableResponse, error) {
	file := s.update.Input(in)
	file, err := s.service.Update(file)
	if err != nil {
		return nil, err
	}
	return s.update.Output(file), nil
}

// Delete implements RepositoryServer.Delete
func (s *CRUDService) Delete(ctx context.Context, in *server.DeleteTableRequest) (*server.DeleteTableResponse, error) {
	file := s.delete.Input(in)
	file, err := s.service.Delete(file)
	if err != nil {
		return nil, err
	}
	return s.delete.Output(file), nil
}
