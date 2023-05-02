package table

import (
	"context"
	"sample/s3-grpc-server/infra/repository/table"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	createTable CreateTable
	readTable   ReadTable
	updateTable UpdateTable
	deleteTable DeleteTable
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
	file := s.createTable.Input(in)
	file, err := s.service.Create(file)
	if err != nil {
		return nil, err
	}
	return s.createTable.Output(file), nil
}

// Read implements RepositoryServer.Read
func (s *CRUDService) Read(ctx context.Context, in *server.ReadTableRequest) (*server.ReadTableResponse, error) {
	file := s.readTable.Input(in)
	file, err := s.service.Read(file)
	if err != nil {
		return nil, err
	}
	return s.readTable.Output(file), nil
}

// Update implements RepositoryServer.Update
func (s *CRUDService) Update(ctx context.Context, in *server.UpdateTableRequest) (*server.UpdateTableResponse, error) {
	file := s.updateTable.Input(in)
	file, err := s.service.Update(file)
	if err != nil {
		return nil, err
	}
	return s.updateTable.Output(file), nil
}

// Delete implements RepositoryServer.Delete
func (s *CRUDService) Delete(ctx context.Context, in *server.DeleteTableRequest) (*server.DeleteTableResponse, error) {
	file := s.deleteTable.Input(in)
	file, err := s.service.Delete(file)
	if err != nil {
		return nil, err
	}
	return s.deleteTable.Output(file), nil
}
