package cell

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type clientGateway struct {
	create CreateCell
	read   ReadCell
	update UpdateCell
	delete DeleteCell
}

type CellClientRepository struct {
	clientGateway
	client grpc_server.CellRepositoryClient
}

func NewCellRepository(client grpc_server.CellRepositoryClient) *CellClientRepository {
	return &CellClientRepository{
		client: client,
	}
}

// Create implements cellRepository.Create
func (x *CellClientRepository) Create(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return x.create.ToInfra(cell, func(cell *grpc_server.CreateCellRequest) (*grpc_server.CreateCellResponse, error) {
		return x.client.Create(ctx, cell)
	})
}

// Read implements cellRepository.Read
func (x *CellClientRepository) Read(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return x.read.ToInfra(cell, func(cell *grpc_server.ReadCellRequest) (*grpc_server.ReadCellResponse, error) {
		return x.client.Read(ctx, cell)
	})
}

// Update implements cellRepository.Update
func (x *CellClientRepository) Update(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return x.update.ToInfra(cell, func(cell *grpc_server.UpdateCellRequest) (*grpc_server.UpdateCellResponse, error) {
		return x.client.Update(ctx, cell)
	})
}

// Delete implements cellRepository.Delete
func (x *CellClientRepository) Delete(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return x.delete.ToInfra(cell, func(cell *grpc_server.DeleteCellRequest) (*grpc_server.DeleteCellResponse, error) {
		return x.client.Delete(ctx, cell)
	})
}
