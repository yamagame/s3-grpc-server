package cell

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/gateway"
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
	return gateway.ToInfra(cell, x.create.Input, func(cell *grpc_server.CreateCellRequest) (*grpc_server.CreateCellResponse, error) {
		return x.client.Create(ctx, cell)
	}, x.create.Output)
}

// Read implements cellRepository.Read
func (x *CellClientRepository) Read(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return gateway.ToInfra(cell, x.read.Input, func(cell *grpc_server.ReadCellRequest) (*grpc_server.ReadCellResponse, error) {
		return x.client.Read(ctx, cell)
	}, x.read.Output)
}

// Update implements cellRepository.Update
func (x *CellClientRepository) Update(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return gateway.ToInfra(cell, x.update.Input, func(cell *grpc_server.UpdateCellRequest) (*grpc_server.UpdateCellResponse, error) {
		return x.client.Update(ctx, cell)
	}, x.update.Output)
}

// Delete implements cellRepository.Delete
func (x *CellClientRepository) Delete(ctx context.Context, cell *model.Cell) (*model.Cell, error) {
	return gateway.ToInfra(cell, x.delete.Input, func(cell *grpc_server.DeleteCellRequest) (*grpc_server.DeleteCellResponse, error) {
		return x.client.Delete(ctx, cell)
	}, x.delete.Output)
}
