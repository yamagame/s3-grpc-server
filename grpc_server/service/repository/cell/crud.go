package cell

import (
	"context"
	"sample/s3-grpc-server/infra/repository/cell"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	create CreateCell
	read   ReadCell
	update UpdateCell
	delete DeleteCell
}

type CRUDService struct {
	serverDTO
	service cell.RepositoryInterface
	server.UnimplementedRepositoryServer
}

func NewCRUDService(service cell.RepositoryInterface) *CRUDService {
	return &CRUDService{
		service: service,
	}
}

// Create implements RepositoryServer.Create
func (s *CRUDService) Create(ctx context.Context, in *server.CreateCellRequest) (*server.CreateCellResponse, error) {
	file := s.create.Input(in)
	file, err := s.service.Create(file)
	if err != nil {
		return nil, err
	}
	return s.create.Output(file), nil
}

// Read implements RepositoryServer.Read
func (s *CRUDService) Read(ctx context.Context, in *server.ReadCellRequest) (*server.ReadCellResponse, error) {
	file := s.read.Input(in)
	file, err := s.service.Read(file)
	if err != nil {
		return nil, err
	}
	return s.read.Output(file), nil
}

// Update implements RepositoryServer.Update
func (s *CRUDService) Update(ctx context.Context, in *server.UpdateCellRequest) (*server.UpdateCellResponse, error) {
	file := s.update.Input(in)
	file, err := s.service.Update(file)
	if err != nil {
		return nil, err
	}
	return s.update.Output(file), nil
}

// Delete implements RepositoryServer.Delete
func (s *CRUDService) Delete(ctx context.Context, in *server.DeleteCellRequest) (*server.DeleteCellResponse, error) {
	file := s.delete.Input(in)
	file, err := s.service.Delete(file)
	if err != nil {
		return nil, err
	}
	return s.delete.Output(file), nil
}
