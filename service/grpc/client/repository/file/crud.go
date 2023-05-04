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

type FileRepository struct {
	clientDTO
	client server.FileRepositoryClient
}

func NewFileRepository(client server.FileRepositoryClient) *FileRepository {
	return &FileRepository{
		client: client,
	}
}

// Create implements fileRepository.Create
func (x *FileRepository) Create(ctx context.Context, file *model.File) (*model.File, error) {
	return x.create.Domain(file, func(req *server.CreateFileRequest) (*server.CreateFileResponse, error) {
		return x.client.Create(ctx, req)
	})
}

// Read implements fileRepository.Read
func (x *FileRepository) Read(ctx context.Context, file *model.File) (*model.File, error) {
	return x.read.Domain(file, func(req *server.ReadFileRequest) (*server.ReadFileResponse, error) {
		return x.client.Read(ctx, req)
	})
}

// Update implements fileRepository.Create
func (x *FileRepository) Update(ctx context.Context, file *model.File) (*model.File, error) {
	return x.update.Domain(file, func(req *server.UpdateFileRequest) (*server.UpdateFileResponse, error) {
		return x.client.Update(ctx, req)
	})
}

// Delete implements fileRepository.Delete
func (x *FileRepository) Delete(ctx context.Context, file *model.File) (*model.File, error) {
	return x.delete.Domain(file, func(req *server.DeleteFileRequest) (*server.DeleteFileResponse, error) {
		return x.client.Delete(ctx, req)
	})
}
