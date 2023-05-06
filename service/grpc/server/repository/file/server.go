package file

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository/file"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/gateway"
)

type serverGateway struct {
	create CreateFile
	read   ReadFile
	update UpdateFile
	delete DeleteFile
}

type FileServerRepository struct {
	serverGateway
	repository file.RepositoryInterface
	grpc_server.UnimplementedFileRepositoryServer
}

func NewFileServerRepository(repository file.RepositoryInterface) *FileServerRepository {
	return &FileServerRepository{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *FileServerRepository) Create(ctx context.Context, in *grpc_server.CreateFileRequest) (*grpc_server.CreateFileResponse, error) {
	return gateway.ToDomain(in, s.create.Input, func(file *model.File) (*model.File, error) {
		return s.repository.Create(ctx, file)
	}, s.create.Output)
}

// Read implements RepositoryServer.Read
func (s *FileServerRepository) Read(ctx context.Context, in *grpc_server.ReadFileRequest) (*grpc_server.ReadFileResponse, error) {
	return gateway.ToDomain(in, s.read.Input, func(file *model.File) (*model.File, error) {
		return s.repository.Read(ctx, file)
	}, s.read.Output)
}

// Update implements RepositoryServer.Update
func (s *FileServerRepository) Update(ctx context.Context, in *grpc_server.UpdateFileRequest) (*grpc_server.UpdateFileResponse, error) {
	return gateway.ToDomain(in, s.update.Input, func(file *model.File) (*model.File, error) {
		return s.repository.Update(ctx, file)
	}, s.update.Output)
}

// Delete implements RepositoryServer.Delete
func (s *FileServerRepository) Delete(ctx context.Context, in *grpc_server.DeleteFileRequest) (*grpc_server.DeleteFileResponse, error) {
	return gateway.ToDomain(in, s.delete.Input, func(file *model.File) (*model.File, error) {
		return s.repository.Delete(ctx, file)
	}, s.delete.Output)
}
