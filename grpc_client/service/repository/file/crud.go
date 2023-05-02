package file

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type clientDTO struct {
	createFile CreateFileDTO
	readFile   ReadFileDTO
	updateFile UpdateFileDTO
	deleteFile DeleteFileDTO
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

func (x *CRUDClient) Create(ctx context.Context) (*model.File, error) {
	req := x.createFile.Input(x.scanner.Create())
	res, err := x.client.CreateFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createFile.Output(res), nil
}

func (x *CRUDClient) Read(ctx context.Context) (*model.File, error) {
	req := x.readFile.Input(x.scanner.Read())
	res, err := x.client.ReadFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.readFile.Output(res), nil
}

func (x *CRUDClient) Update(ctx context.Context) (*model.File, error) {
	req := x.updateFile.Input(x.scanner.Update())
	res, err := x.client.UpdateFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.updateFile.Output(res), nil
}

func (x *CRUDClient) Delete(ctx context.Context) (*model.File, error) {
	req := x.deleteFile.Input(x.scanner.Delete())
	res, err := x.client.DeleteFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteFile.Output(res), nil
}
