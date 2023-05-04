package cell

import (
	"context"
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	create CreateCell
	read   ReadCell
	update UpdateCell
	delete DeleteCell
}

type CellRepositoryServer struct {
	serverDTO
	repository cell.RepositoryInterface
	grpc_server.UnimplementedCellRepositoryServer
}

func NewCellRepositoryServer(repository cell.RepositoryInterface) *CellRepositoryServer {
	return &CellRepositoryServer{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *CellRepositoryServer) Create(ctx context.Context, in *grpc_server.CreateCellRequest) (*grpc_server.CreateCellResponse, error) {
	return s.create.ToDomain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Create(ctx, cell)
	})
}

// Read implements RepositoryServer.Read
func (s *CellRepositoryServer) Read(ctx context.Context, in *grpc_server.ReadCellRequest) (*grpc_server.ReadCellResponse, error) {
	return s.read.ToDomain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Read(ctx, cell)
	})
}

// Update implements RepositoryServer.Update
func (s *CellRepositoryServer) Update(ctx context.Context, in *grpc_server.UpdateCellRequest) (*grpc_server.UpdateCellResponse, error) {
	return s.update.ToDomain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Update(ctx, cell)
	})
}

// Delete implements RepositoryServer.Delete
func (s *CellRepositoryServer) Delete(ctx context.Context, in *grpc_server.DeleteCellRequest) (*grpc_server.DeleteCellResponse, error) {
	return s.delete.ToDomain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Delete(ctx, cell)
	})
}
