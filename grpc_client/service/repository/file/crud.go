package file

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type clientDTO struct {
	create CreateFile
	read   ReadFile
	update UpdateFile
	delete DeleteFile
}

type fileRepository struct {
	clientDTO
	client server.RepositoryClient
}

func NewFileRepository(client server.RepositoryClient) *fileRepository {
	return &fileRepository{
		client: client,
	}
}

func (x *fileRepository) Create(ctx context.Context, file *model.File) (*model.File, error) {
	return x.create.Domain(file, func(req *server.CreateFileRequest) (*server.CreateFileResponse, error) {
		return x.client.CreateFile(ctx, req)
	})
}

func (x *fileRepository) Read(ctx context.Context, file *model.File) (*model.File, error) {
	return x.read.Domain(file, func(req *server.ReadFileRequest) (*server.ReadFileResponse, error) {
		return x.client.ReadFile(ctx, req)
	})
}

func (x *fileRepository) Update(ctx context.Context, file *model.File) (*model.File, error) {
	return x.update.Domain(file, func(req *server.UpdateFileRequest) (*server.UpdateFileResponse, error) {
		return x.client.UpdateFile(ctx, req)
	})
}

func (x *fileRepository) Delete(ctx context.Context, file *model.File) (*model.File, error) {
	return x.delete.Domain(file, func(req *server.DeleteFileRequest) (*server.DeleteFileResponse, error) {
		return x.client.DeleteFile(ctx, req)
	})
}
