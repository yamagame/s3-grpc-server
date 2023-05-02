package table

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type clientDTO struct {
	createTable CreateTableDTO
	readTable   ReadTableDTO
	updateTable UpdateTableDTO
	deleteTable DeleteTableDTO
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

func (x *CRUDClient) Create(ctx context.Context) (*model.Table, error) {
	req := x.createTable.Input(x.scanner.Create())
	res, err := x.client.CreateTable(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createTable.Output(res), nil
}

func (x *CRUDClient) Read(ctx context.Context) (*model.Table, error) {
	req := x.readTable.Input(x.scanner.Read())
	res, err := x.client.ReadTable(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.readTable.Output(res), nil
}

func (x *CRUDClient) Update(ctx context.Context) (*model.Table, error) {
	req := x.updateTable.Input(x.scanner.Update())
	res, err := x.client.UpdateTable(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.updateTable.Output(res), nil
}

func (x *CRUDClient) Delete(ctx context.Context) (*model.Table, error) {
	req := x.deleteTable.Input(x.scanner.Delete())
	res, err := x.client.DeleteTable(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteTable.Output(res), nil
}
