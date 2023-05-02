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

type TableClient struct {
	clientDTO
	scanner TableScannerInterface
	client  server.RepositoryClient
}

func NewTableClient(client server.RepositoryClient, scanner TableScannerInterface) *TableClient {
	return &TableClient{
		scanner: scanner,
		client:  client,
	}
}

func (x *TableClient) Create(ctx context.Context) (*model.Table, error) {
	req := x.createTable.Input(x.scanner.CreateTable())
	res, err := x.client.CreateTable(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createTable.Output(res), nil
}

func (x *TableClient) Read(ctx context.Context) (*model.Table, error) {
	req := x.readTable.Input(x.scanner.ReadTable())
	res, err := x.client.ReadTable(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.readTable.Output(res), nil
}

func (x *TableClient) Update(ctx context.Context) (*model.Table, error) {
	req := x.updateTable.Input(x.scanner.UpdateTable())
	res, err := x.client.UpdateTable(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.updateTable.Output(res), nil
}

func (x *TableClient) Delete(ctx context.Context) (*model.Table, error) {
	req := x.deleteTable.Input(x.scanner.DeleteTable())
	res, err := x.client.DeleteTable(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteTable.Output(res), nil
}
