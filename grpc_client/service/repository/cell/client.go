package cell

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type clientDTO struct {
	createCell CreateCellDTO
	readCell   ReadCellDTO
	updateCell UpdateCellDTO
	deleteCell DeleteCellDTO
}

type CellClient struct {
	clientDTO
	scanner CellScannerInterface
	client  server.RepositoryClient
}

func NewCellClient(client server.RepositoryClient, scanner CellScannerInterface) *CellClient {
	return &CellClient{
		scanner: scanner,
		client:  client,
	}
}

func (x *CellClient) Create(ctx context.Context) (*model.Cell, error) {
	req := x.createCell.Input(x.scanner.CreateCell())
	res, err := x.client.CreateCell(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createCell.Output(res), nil
}

func (x *CellClient) Read(ctx context.Context) (*model.Cell, error) {
	req := x.readCell.Input(x.scanner.ReadCell())
	res, err := x.client.ReadCell(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.readCell.Output(res), nil
}

func (x *CellClient) Update(ctx context.Context) (*model.Cell, error) {
	req := x.updateCell.Input(x.scanner.UpdateCell())
	res, err := x.client.UpdateCell(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.updateCell.Output(res), nil
}

func (x *CellClient) Delete(ctx context.Context) (*model.Cell, error) {
	req := x.deleteCell.Input(x.scanner.DeleteCell())
	res, err := x.client.DeleteCell(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteCell.Output(res), nil
}
