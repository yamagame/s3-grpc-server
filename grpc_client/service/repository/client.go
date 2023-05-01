package repository

import (
	"context"
	"sample/s3-grpc-server/infra/repository"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type repositoryClientMessage struct {
	createFileInfo RepositoryCreateFilieInfoClientMessage
	readFileInfo   RepositoryReadFileInfoClientMessage
	updateFileInfo RepositoryUpdateFileInfoClientMessage
	deleteFileInfo RepositoryDeleteFileInfoClientMessage
}

type RepositoryClient struct {
	repositoryClientMessage
	scanner RepositoryScannerInterface
	client  server.RepositoryClient
}

func NewRepositoryClient(client server.RepositoryClient, scanner RepositoryScannerInterface) *RepositoryClient {
	return &RepositoryClient{
		scanner: scanner,
		client:  client,
	}
}

func (x *RepositoryClient) CreateFileInfo(ctx context.Context) (*repository.FileInfo, error) {
	req := x.createFileInfo.Input(x.scanner.CreateFileInfo())
	res, err := x.client.CreateFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createFileInfo.Output(res), nil
}

func (x *RepositoryClient) ReadFileInfo(ctx context.Context) (*repository.FileInfo, error) {
	req := x.readFileInfo.Input(x.scanner.ReadFileInfo())
	res, err := x.client.ReadFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.readFileInfo.Output(res), nil
}

func (x *RepositoryClient) UpdateFileInfo(ctx context.Context) (*repository.FileInfo, error) {
	req := x.updateFileInfo.Input(x.scanner.UpdateFileInfo())
	res, err := x.client.UpdateFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.updateFileInfo.Output(res), nil
}

func (x *RepositoryClient) DeleteFileInfo(ctx context.Context) (*repository.FileInfo, error) {
	req := x.deleteFileInfo.Input(x.scanner.DeleteFileInfo())
	res, err := x.client.DeleteFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteFileInfo.Output(res), nil
}
