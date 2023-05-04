package file

import (
	"context"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type clientGateway struct {
	create CreateFile
	read   ReadFile
	update UpdateFile
	delete DeleteFile
}

type FileRepository struct {
	clientGateway
	client grpc_server.FileRepositoryClient
}

func NewFileRepository(client grpc_server.FileRepositoryClient) *FileRepository {
	return &FileRepository{
		client: client,
	}
}

// Create implements fileRepository.Create
func (x *FileRepository) Create(ctx context.Context, file *model.File) (*model.File, error) {
	return x.create.ToInfra(file, func(req *grpc_server.CreateFileRequest) (*grpc_server.CreateFileResponse, error) {
		return x.client.Create(ctx, req)
	})
}

// Read implements fileRepository.Read
func (x *FileRepository) Read(ctx context.Context, file *model.File) (*model.File, error) {
	return x.read.ToInfra(file, func(req *grpc_server.ReadFileRequest) (*grpc_server.ReadFileResponse, error) {
		return x.client.Read(ctx, req)
	})
}

// Update implements fileRepository.Create
func (x *FileRepository) Update(ctx context.Context, file *model.File) (*model.File, error) {
	return x.update.ToInfra(file, func(req *grpc_server.UpdateFileRequest) (*grpc_server.UpdateFileResponse, error) {
		return x.client.Update(ctx, req)
	})
}

// Delete implements fileRepository.Delete
func (x *FileRepository) Delete(ctx context.Context, file *model.File) (*model.File, error) {
	return x.delete.ToInfra(file, func(req *grpc_server.DeleteFileRequest) (*grpc_server.DeleteFileResponse, error) {
		return x.client.Delete(ctx, req)
	})
}
