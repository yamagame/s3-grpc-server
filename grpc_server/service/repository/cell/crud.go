package cell

import (
	"context"
	"sample/s3-grpc-server/infra/repository/cell"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	createCell CreateCell
	readCell   ReadCell
	updateCell UpdateCell
	deleteCell DeleteCell
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
	file := s.createCell.Input(in)
	file, err := s.service.Create(file)
	if err != nil {
		return nil, err
	}
	return s.createCell.Output(file), nil
}

// Read implements RepositoryServer.Read
func (s *CRUDService) Read(ctx context.Context, in *server.ReadCellRequest) (*server.ReadCellResponse, error) {
	file := s.readCell.Input(in)
	file, err := s.service.Read(file)
	if err != nil {
		return nil, err
	}
	return s.readCell.Output(file), nil
}

// Update implements RepositoryServer.Update
func (s *CRUDService) Update(ctx context.Context, in *server.UpdateCellRequest) (*server.UpdateCellResponse, error) {
	file := s.updateCell.Input(in)
	file, err := s.service.Update(file)
	if err != nil {
		return nil, err
	}
	return s.updateCell.Output(file), nil
}

// Delete implements RepositoryServer.Delete
func (s *CRUDService) Delete(ctx context.Context, in *server.DeleteCellRequest) (*server.DeleteCellResponse, error) {
	file := s.deleteCell.Input(in)
	file, err := s.service.Delete(file)
	if err != nil {
		return nil, err
	}
	return s.deleteCell.Output(file), nil
}
