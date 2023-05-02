package fileinfo

import (
	"context"
	"sample/s3-grpc-server/grpc_server/service/repository/dto"
	"sample/s3-grpc-server/infra/repository/fileinfo"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	createFileInfo dto.CreateFileInfo
	readFileInfo   dto.ReadFileInfo
	updateFileInfo dto.UpdateFileInfo
	deleteFileInfo dto.DeleteFileInfo
}

type FileInfoServer struct {
	serverDTO
	service fileinfo.RepositoryInterface
	server.UnimplementedRepositoryServer
}

func NewFileInfoServer(service fileinfo.RepositoryInterface) *FileInfoServer {
	return &FileInfoServer{
		service: service,
	}
}

// Create implements RepositoryServer.Create
func (s *FileInfoServer) Create(ctx context.Context, in *server.CreateFileInfoRequest) (*server.CreateFileInfoResponse, error) {
	fileInfo := s.createFileInfo.Input(in)
	fileInfo, err := s.service.Create(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.createFileInfo.Output(fileInfo), nil
}

// Read implements RepositoryServer.Read
func (s *FileInfoServer) Read(ctx context.Context, in *server.ReadFileInfoRequest) (*server.ReadFileInfoResponse, error) {
	fileInfo := s.readFileInfo.Input(in)
	fileInfo, err := s.service.Read(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.readFileInfo.Output(fileInfo), nil
}

// Update implements RepositoryServer.Update
func (s *FileInfoServer) Update(ctx context.Context, in *server.UpdateFileInfoRequest) (*server.UpdateFileInfoResponse, error) {
	fileInfo := s.updateFileInfo.Input(in)
	fileInfo, err := s.service.Update(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.updateFileInfo.Output(fileInfo), nil
}

// Delete implements RepositoryServer.Delete
func (s *FileInfoServer) Delete(ctx context.Context, in *server.DeleteFileInfoRequest) (*server.DeleteFileInfoResponse, error) {
	fileInfo := s.deleteFileInfo.Input(in)
	fileInfo, err := s.service.Delete(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.deleteFileInfo.Output(fileInfo), nil
}
