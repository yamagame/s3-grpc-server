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

type FileClient struct {
	clientDTO
	scanner FileScannerInterface
	client  server.RepositoryClient
}

func NewFileClient(client server.RepositoryClient, scanner FileScannerInterface) *FileClient {
	return &FileClient{
		scanner: scanner,
		client:  client,
	}
}

func (x *FileClient) Create(ctx context.Context) (*model.File, error) {
	req := x.createFile.Input(x.scanner.CreateFile())
	res, err := x.client.CreateFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createFile.Output(res), nil
}

func (x *FileClient) Read(ctx context.Context) (*model.File, error) {
	req := x.readFile.Input(x.scanner.ReadFile())
	res, err := x.client.ReadFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.readFile.Output(res), nil
}

func (x *FileClient) Update(ctx context.Context) (*model.File, error) {
	req := x.updateFile.Input(x.scanner.UpdateFile())
	res, err := x.client.UpdateFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.updateFile.Output(res), nil
}

func (x *FileClient) Delete(ctx context.Context) (*model.File, error) {
	req := x.deleteFile.Input(x.scanner.DeleteFile())
	res, err := x.client.DeleteFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteFile.Output(res), nil
}
