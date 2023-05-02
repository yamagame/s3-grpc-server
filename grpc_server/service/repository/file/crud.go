package file

import (
	"context"
	"sample/s3-grpc-server/infra/repository/file"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	createFile CreateFile
	readFile   ReadFile
	updateFile UpdateFile
	deleteFile DeleteFile
}

type CRUDService struct {
	serverDTO
	service file.RepositoryInterface
	server.UnimplementedRepositoryServer
}

func NewCRUDService(service file.RepositoryInterface) *CRUDService {
	return &CRUDService{
		service: service,
	}
}

// Create implements RepositoryServer.Create
func (s *CRUDService) Create(ctx context.Context, in *server.CreateFileRequest) (*server.CreateFileResponse, error) {
	file := s.createFile.Input(in)
	file, err := s.service.Create(file)
	if err != nil {
		return nil, err
	}
	return s.createFile.Output(file), nil
}

// Read implements RepositoryServer.Read
func (s *CRUDService) Read(ctx context.Context, in *server.ReadFileRequest) (*server.ReadFileResponse, error) {
	file := s.readFile.Input(in)
	file, err := s.service.Read(file)
	if err != nil {
		return nil, err
	}
	return s.readFile.Output(file), nil
}

// Update implements RepositoryServer.Update
func (s *CRUDService) Update(ctx context.Context, in *server.UpdateFileRequest) (*server.UpdateFileResponse, error) {
	file := s.updateFile.Input(in)
	file, err := s.service.Update(file)
	if err != nil {
		return nil, err
	}
	return s.updateFile.Output(file), nil
}

// Delete implements RepositoryServer.Delete
func (s *CRUDService) Delete(ctx context.Context, in *server.DeleteFileRequest) (*server.DeleteFileResponse, error) {
	file := s.deleteFile.Input(in)
	file, err := s.service.Delete(file)
	if err != nil {
		return nil, err
	}
	return s.deleteFile.Output(file), nil
}
