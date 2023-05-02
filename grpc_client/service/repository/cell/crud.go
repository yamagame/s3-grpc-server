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

type CRUDClient struct {
	clientDTO
	scanner ScannerInterface
	client  server.RepositoryClient
}

func NewCRUDClient(client server.RepositoryClient, scanner ScannerInterface) *CRUDClient {
	return &CRUDClient{
		scanner: scanner,
		client:  client,
	}
}

func (x *CRUDClient) Create(ctx context.Context) (*model.Cell, error) {
	req := x.createCell.Input(x.scanner.Create())
	res, err := x.client.CreateCell(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createCell.Output(res), nil
}

func (x *CRUDClient) Read(ctx context.Context) (*model.Cell, error) {
	req := x.readCell.Input(x.scanner.Read())
	res, err := x.client.ReadCell(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.readCell.Output(res), nil
}

func (x *CRUDClient) Update(ctx context.Context) (*model.Cell, error) {
	req := x.updateCell.Input(x.scanner.Update())
	res, err := x.client.UpdateCell(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.updateCell.Output(res), nil
}

func (x *CRUDClient) Delete(ctx context.Context) (*model.Cell, error) {
	req := x.deleteCell.Input(x.scanner.Delete())
	res, err := x.client.DeleteCell(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteCell.Output(res), nil
}
