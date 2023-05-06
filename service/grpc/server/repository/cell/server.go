package cell

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/gateway"
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
	return gateway.ToDomain(in, s.create.Input, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Create(ctx, cell)
	}, s.create.Output)
}

// Read implements RepositoryServer.Read
func (s *CellServerRepository) Read(ctx context.Context, in *grpc_server.ReadCellRequest) (*grpc_server.ReadCellResponse, error) {
	return gateway.ToDomain(in, s.read.Input, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Read(ctx, cell)
	}, s.read.Output)
}

// Update implements RepositoryServer.Update
func (s *CellServerRepository) Update(ctx context.Context, in *grpc_server.UpdateCellRequest) (*grpc_server.UpdateCellResponse, error) {
	return gateway.ToDomain(in, s.update.Input, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Update(ctx, cell)
	}, s.update.Output)
}

// Delete implements RepositoryServer.Delete
func (s *CellServerRepository) Delete(ctx context.Context, in *grpc_server.DeleteCellRequest) (*grpc_server.DeleteCellResponse, error) {
	return gateway.ToDomain(in, s.delete.Input, func(cell *model.Cell) (*model.Cell, error) {
		return s.repository.Delete(ctx, cell)
	}, s.delete.Output)
}
