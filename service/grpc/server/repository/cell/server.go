package cell

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/proto/grpc_server"
)

type serverGateway struct {
	create CreateCell
	read   ReadCell
	update UpdateCell
	delete DeleteCell
}

type CellServerRepository struct {
	serverGateway
	repository cell.RepositoryInterface
	grpc_server.UnimplementedCellRepositoryServer
}

func NewCellServerRepository(repository cell.RepositoryInterface) *CellServerRepository {
	return &CellServerRepository{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *CellServerRepository) Create(ctx context.Context, in *grpc_server.CreateCellRequest) (*grpc_server.CreateCellResponse, error) {
	return s.create.ToDomain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Create(ctx, cell)
	})
}

// Read implements RepositoryServer.Read
func (s *CellServerRepository) Read(ctx context.Context, in *grpc_server.ReadCellRequest) (*grpc_server.ReadCellResponse, error) {
	return s.read.ToDomain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Read(ctx, cell)
	})
}

// Update implements RepositoryServer.Update
func (s *CellServerRepository) Update(ctx context.Context, in *grpc_server.UpdateCellRequest) (*grpc_server.UpdateCellResponse, error) {
	return s.update.ToDomain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Update(ctx, cell)
	})
}

// Delete implements RepositoryServer.Delete
func (s *CellServerRepository) Delete(ctx context.Context, in *grpc_server.DeleteCellRequest) (*grpc_server.DeleteCellResponse, error) {
	return s.delete.ToDomain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Delete(ctx, cell)
	})
}
