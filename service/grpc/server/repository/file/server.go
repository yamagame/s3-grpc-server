package file

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository/file"
	"sample/s3-grpc-server/proto/grpc_server"
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
	return s.create.ToDomain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Create(ctx, file)
	})
}

// Read implements RepositoryServer.Read
func (s *FileServerRepository) Read(ctx context.Context, in *grpc_server.ReadFileRequest) (*grpc_server.ReadFileResponse, error) {
	return s.read.ToDomain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Read(ctx, file)
	})
}

// Update implements RepositoryServer.Update
func (s *FileServerRepository) Update(ctx context.Context, in *grpc_server.UpdateFileRequest) (*grpc_server.UpdateFileResponse, error) {
	return s.update.ToDomain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Update(ctx, file)
	})
}

// Delete implements RepositoryServer.Delete
func (s *FileServerRepository) Delete(ctx context.Context, in *grpc_server.DeleteFileRequest) (*grpc_server.DeleteFileResponse, error) {
	return s.delete.ToDomain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Delete(ctx, file)
	})
}
