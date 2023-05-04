package file

import (
	"context"
	"sample/s3-grpc-server/infra/repository/file"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	create CreateFile
	read   ReadFile
	update UpdateFile
	delete DeleteFile
}

type FileRepository struct {
	serverDTO
	repository file.RepositoryInterface
	grpc_server.UnimplementedFileRepositoryServer
}

func NewFileRepository(repository file.RepositoryInterface) *FileRepository {
	return &FileRepository{
		repository: repository,
	}
}

// Create implements RepositoryServer.Create
func (s *FileRepository) Create(ctx context.Context, in *server.CreateFileRequest) (*server.CreateFileResponse, error) {
	return s.create.Domain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Create(ctx, file)
	})
}

// Read implements RepositoryServer.Read
func (s *FileRepository) Read(ctx context.Context, in *server.ReadFileRequest) (*server.ReadFileResponse, error) {
	return s.read.Domain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Read(ctx, file)
	})
}

// Update implements RepositoryServer.Update
func (s *FileRepository) Update(ctx context.Context, in *server.UpdateFileRequest) (*server.UpdateFileResponse, error) {
	return s.update.Domain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Update(ctx, file)
	})
}

// Delete implements RepositoryServer.Delete
func (s *FileRepository) Delete(ctx context.Context, in *server.DeleteFileRequest) (*server.DeleteFileResponse, error) {
	return s.delete.Domain(in, func(file *model.File) (*model.File, error) {
		return s.repository.Delete(ctx, file)
	})
}
