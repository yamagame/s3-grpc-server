package cell

import (
	"context"
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	create CreateCell
	read   ReadCell
	update UpdateCell
	delete DeleteCell
}

type CellRepository struct {
	serverDTO
	repository cell.RepositoryInterface
}

func NewCellRepository(repository cell.RepositoryInterface) *CellRepository {
	return &CellRepository{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *CellRepository) Create(ctx context.Context, in *server.CreateCellRequest) (*server.CreateCellResponse, error) {
	return s.create.Domain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Create(ctx, cell)
	})
}

// Read implements RepositoryServer.Read
func (s *CellRepository) Read(ctx context.Context, in *server.ReadCellRequest) (*server.ReadCellResponse, error) {
	return s.read.Domain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Read(ctx, cell)
	})
}

// Update implements RepositoryServer.Update
func (s *CellRepository) Update(ctx context.Context, in *server.UpdateCellRequest) (*server.UpdateCellResponse, error) {
	return s.update.Domain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Update(ctx, cell)
	})
}

// Delete implements RepositoryServer.Delete
func (s *CellRepository) Delete(ctx context.Context, in *server.DeleteCellRequest) (*server.DeleteCellResponse, error) {
	return s.delete.Domain(in, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Delete(ctx, cell)
	})
}
