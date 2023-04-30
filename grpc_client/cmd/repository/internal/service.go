package internal

import (
	"context"
	"sample/s3-grpc-server/infra/repository"
	server "sample/s3-grpc-server/proto/grpc-server"
)

type repositoryClientDomain struct {
	createFileInfo repository.RepositoryCreateFilieInfoClientData
	readFileInfo   repository.RepositoryReadFileInfoClientData
	updateFileInfo repository.RepositoryUpdateFileInfoClientData
	deleteFileInfo repository.RepositoryDeleteFileInfoClientData
}

type RepositoryClient struct {
	scanner RepositoryScannerInterface
	domain  repositoryClientDomain
	client  server.RepositoryClient
}

func NewRepositoryClient(aws server.RepositoryClient, scanner RepositoryScannerInterface) *RepositoryClient {
	return &RepositoryClient{
		scanner: scanner,
		domain:  repositoryClientDomain{},
		client:  aws,
	}
}

func (x *RepositoryClient) CreateFileInfo(ctx context.Context) (*repository.FileInfo, error) {
	req := x.domain.createFileInfo.Input(x.scanner.CreateFileInfo())
	res, err := x.client.CreateFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.createFileInfo.Output(res), nil
}

func (x *RepositoryClient) ReadFileInfo(ctx context.Context) (*repository.FileInfo, error) {
	req := x.domain.readFileInfo.Input(x.scanner.ReadFileInfo())
	res, err := x.client.ReadFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.readFileInfo.Output(res), nil
}

func (x *RepositoryClient) UpdateFileInfo(ctx context.Context) (*repository.FileInfo, error) {
	req := x.domain.updateFileInfo.Input(x.scanner.UpdateFileInfo())
	res, err := x.client.UpdateFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.updateFileInfo.Output(res), nil
}

func (x *RepositoryClient) DeleteFileInfo(ctx context.Context) (*repository.FileInfo, error) {
	req := x.domain.deleteFileInfo.Input(x.scanner.DeleteFileInfo())
	res, err := x.client.DeleteFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.deleteFileInfo.Output(res), nil
}
