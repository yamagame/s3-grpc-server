package file

import (
	"context"
	"sample/s3-grpc-server/infra/repository/file"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	create CreateFile
	read   ReadFile
	update UpdateFile
	delete DeleteFile
}

type FileRepositoryServer struct {
	serverDTO
	repository file.RepositoryInterface
	grpc_server.UnimplementedFileRepositoryServer
}

func NewFileRepositoryServer(repository file.RepositoryInterface) *FileRepositoryServer {
	return &FileRepositoryServer{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *FileRepositoryServer) Create(ctx context.Context, in *grpc_server.CreateFileRequest) (*grpc_server.CreateFileResponse, error) {
	return s.create.ToDomain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Create(ctx, file)
	})
}

// Read implements RepositoryServer.Read
func (s *FileRepositoryServer) Read(ctx context.Context, in *grpc_server.ReadFileRequest) (*grpc_server.ReadFileResponse, error) {
	return s.read.ToDomain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Read(ctx, file)
	})
}

// Update implements RepositoryServer.Update
func (s *FileRepositoryServer) Update(ctx context.Context, in *grpc_server.UpdateFileRequest) (*grpc_server.UpdateFileResponse, error) {
	return s.update.ToDomain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Update(ctx, file)
	})
}

// Delete implements RepositoryServer.Delete
func (s *FileRepositoryServer) Delete(ctx context.Context, in *grpc_server.DeleteFileRequest) (*grpc_server.DeleteFileResponse, error) {
	return s.delete.ToDomain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Delete(ctx, file)
	})
}
