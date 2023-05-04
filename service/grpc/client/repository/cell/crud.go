package cell

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type clientDTO struct {
	create CreateCell
	read   ReadCell
	update UpdateCell
	delete DeleteCell
}

type CellRepository struct {
	clientDTO
	client server.CellRepositoryClient
}

func NewCellRepository(client server.CellRepositoryClient) *CellRepository {
	return &CellRepository{
		client: client,
	}
}

// Create implements cellRepository.Create
func (x *CellRepository) Create(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return x.create.Domain(cell, func(cell *server.CreateCellRequest) (*server.CreateCellResponse, error) {
		return x.client.Create(ctx, cell)
	})
}

// Read implements cellRepository.Read
func (x *CellRepository) Read(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return x.read.Domain(cell, func(cell *server.ReadCellRequest) (*server.ReadCellResponse, error) {
		return x.client.Read(ctx, cell)
	})
}

// Update implements cellRepository.Update
func (x *CellRepository) Update(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return x.update.Domain(cell, func(cell *server.UpdateCellRequest) (*server.UpdateCellResponse, error) {
		return x.client.Update(ctx, cell)
	})
}

// Delete implements cellRepository.Delete
func (x *CellRepository) Delete(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return x.delete.Domain(cell, func(cell *server.DeleteCellRequest) (*server.DeleteCellResponse, error) {
		return x.client.Delete(ctx, cell)
	})
}
