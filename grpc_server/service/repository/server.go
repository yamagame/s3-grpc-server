package repository

import (
	"context"
	"sample/s3-grpc-server/infra/repository/fileinfo"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type repositoryServerDTO struct {
	createFileInfo CreateFilieInfoDTO
	readFileInfo   ReadFileInfoDTO
	updateFileInfo UpdateFileInfoDTO
	deleteFileInfo DeleteFileInfoDTO
}

type repositoryServer struct {
	repositoryServerDTO
	service fileinfo.RepositoryInterface
	server.UnimplementedRepositoryServer
}

func NewRepositoryServer(service fileinfo.RepositoryInterface) *repositoryServer {
	return &repositoryServer{
		service: service,
	}
}

// CreateFileInfo implements repositoryServer.CreateFileInfo
func (s *repositoryServer) CreateFileInfo(ctx context.Context, in *server.CreateFileInfoRequest) (*server.CreateFileInfoResponse, error) {
	fileInfo := s.createFileInfo.Input(in)
	fileInfo, err := s.service.CreateFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.createFileInfo.Output(fileInfo), nil
}

// ReadFileInfo implements repositoryServer.ReadFileInfo
func (s *repositoryServer) ReadFileInfo(ctx context.Context, in *server.ReadFileInfoRequest) (*server.ReadFileInfoResponse, error) {
	fileInfo := s.readFileInfo.Input(in)
	fileInfo, err := s.service.ReadFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.readFileInfo.Output(fileInfo), nil
}

// UpdateFileInfo implements repositoryServer.UpdateFileInfo
func (s *repositoryServer) UpdateFileInfo(ctx context.Context, in *server.UpdateFileInfoRequest) (*server.UpdateFileInfoResponse, error) {
	fileInfo := s.updateFileInfo.Input(in)
	fileInfo, err := s.service.UpdateFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.updateFileInfo.Output(fileInfo), nil
}

// DeleteFileInfo implements repositoryServer.DeleteFileInfo
func (s *repositoryServer) DeleteFileInfo(ctx context.Context, in *server.DeleteFileInfoRequest) (*server.DeleteFileInfoResponse, error) {
	fileInfo := s.deleteFileInfo.Input(in)
	fileInfo, err := s.service.DeleteFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.deleteFileInfo.Output(fileInfo), nil
}
