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

type CellServer struct {
	serverDTO
	service cell.RepositoryInterface
	server.UnimplementedRepositoryServer
}

func NewCellServer(service cell.RepositoryInterface) *CellServer {
	return &CellServer{
		service: service,
	}
}

// Create implements RepositoryServer.Create
func (s *CellServer) Create(ctx context.Context, in *server.CreateCellRequest) (*server.CreateCellResponse, error) {
	fileInfo := s.createCell.Input(in)
	fileInfo, err := s.service.Create(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.createCell.Output(fileInfo), nil
}

// Read implements RepositoryServer.Read
func (s *CellServer) Read(ctx context.Context, in *server.ReadCellRequest) (*server.ReadCellResponse, error) {
	fileInfo := s.readCell.Input(in)
	fileInfo, err := s.service.Read(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.readCell.Output(fileInfo), nil
}

// Update implements RepositoryServer.Update
func (s *CellServer) Update(ctx context.Context, in *server.UpdateCellRequest) (*server.UpdateCellResponse, error) {
	fileInfo := s.updateCell.Input(in)
	fileInfo, err := s.service.Update(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.updateCell.Output(fileInfo), nil
}

// Delete implements RepositoryServer.Delete
func (s *CellServer) Delete(ctx context.Context, in *server.DeleteCellRequest) (*server.DeleteCellResponse, error) {
	fileInfo := s.deleteCell.Input(in)
	fileInfo, err := s.service.Delete(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.deleteCell.Output(fileInfo), nil
}
