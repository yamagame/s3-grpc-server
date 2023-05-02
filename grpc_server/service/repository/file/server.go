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

type FileServer struct {
	serverDTO
	service file.RepositoryInterface
	server.UnimplementedRepositoryServer
}

func NewFileServer(service file.RepositoryInterface) *FileServer {
	return &FileServer{
		service: service,
	}
}

// Create implements RepositoryServer.Create
func (s *FileServer) Create(ctx context.Context, in *server.CreateFileRequest) (*server.CreateFileResponse, error) {
	fileInfo := s.createFile.Input(in)
	fileInfo, err := s.service.Create(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.createFile.Output(fileInfo), nil
}

// Read implements RepositoryServer.Read
func (s *FileServer) Read(ctx context.Context, in *server.ReadFileRequest) (*server.ReadFileResponse, error) {
	fileInfo := s.readFile.Input(in)
	fileInfo, err := s.service.Read(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.readFile.Output(fileInfo), nil
}

// Update implements RepositoryServer.Update
func (s *FileServer) Update(ctx context.Context, in *server.UpdateFileRequest) (*server.UpdateFileResponse, error) {
	fileInfo := s.updateFile.Input(in)
	fileInfo, err := s.service.Update(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.updateFile.Output(fileInfo), nil
}

// Delete implements RepositoryServer.Delete
func (s *FileServer) Delete(ctx context.Context, in *server.DeleteFileRequest) (*server.DeleteFileResponse, error) {
	fileInfo := s.deleteFile.Input(in)
	fileInfo, err := s.service.Delete(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.deleteFile.Output(fileInfo), nil
}
