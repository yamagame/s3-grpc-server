package fileinfo

import (
	"context"
	"sample/s3-grpc-server/grpc_client/service/repository/dto"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type clientDTO struct {
	createFileInfo dto.CreateFilieInfoDTO
	readFileInfo   dto.ReadFileInfoDTO
	updateFileInfo dto.UpdateFileInfoDTO
	deleteFileInfo dto.DeleteFileInfoDTO
}

type FileInfoClient struct {
	clientDTO
	scanner FileInfoScannerInterface
	client  server.RepositoryClient
}

func NewFileInfoClient(client server.RepositoryClient, scanner FileInfoScannerInterface) *FileInfoClient {
	return &FileInfoClient{
		scanner: scanner,
		client:  client,
	}
}

func (x *FileInfoClient) CreateFileInfo(ctx context.Context) (*model.FileInfo, error) {
	req := x.createFileInfo.Input(x.scanner.CreateFileInfo())
	res, err := x.client.CreateFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.createFileInfo.Output(res), nil
}

func (x *FileInfoClient) ReadFileInfo(ctx context.Context) (*model.FileInfo, error) {
	req := x.readFileInfo.Input(x.scanner.ReadFileInfo())
	res, err := x.client.ReadFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.readFileInfo.Output(res), nil
}

func (x *FileInfoClient) UpdateFileInfo(ctx context.Context) (*model.FileInfo, error) {
	req := x.updateFileInfo.Input(x.scanner.UpdateFileInfo())
	res, err := x.client.UpdateFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.updateFileInfo.Output(res), nil
}

func (x *FileInfoClient) DeleteFileInfo(ctx context.Context) (*model.FileInfo, error) {
	req := x.deleteFileInfo.Input(x.scanner.DeleteFileInfo())
	res, err := x.client.DeleteFileInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.deleteFileInfo.Output(res), nil
}
