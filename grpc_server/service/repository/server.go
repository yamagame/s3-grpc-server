package repository

import (
	"context"
	"sample/s3-grpc-server/infra/repository"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type repositoryServerMessage struct {
	createFileInfo RepositoryCreateFilieInfoServerMessage
	readFileInfo   RepositoryReadFileInfoServerMessage
	updateFileInfo RepositoryUpdateFileInfoServerMessage
	deleteFileInfo RepositoryDeleteFileInfoServerMessage
}

type repositoryServer struct {
	service *repository.RepositoryService
	message *repositoryServerMessage
	server.UnimplementedRepositoryServer
}

func NewRepositoryServer(service *repository.RepositoryService) *repositoryServer {
	return &repositoryServer{
		service: service,
		message: &repositoryServerMessage{},
	}
}

// CreateFileInfo implements repositoryServer.CreateFileInfo
func (s *repositoryServer) CreateFileInfo(ctx context.Context, in *server.CreateFileInfoRequest) (*server.CreateFileInfoResponse, error) {
	fileInfo := s.message.createFileInfo.Input(in)
	fileInfo, err := s.service.CreateFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.message.createFileInfo.Output(fileInfo), nil
}

// ReadFileInfo implements repositoryServer.ReadFileInfo
func (s *repositoryServer) ReadFileInfo(ctx context.Context, in *server.ReadFileInfoRequest) (*server.ReadFileInfoResponse, error) {
	fileInfo := s.message.readFileInfo.Input(in)
	fileInfo, err := s.service.ReadFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.message.readFileInfo.Output(fileInfo), nil
}

// UpdateFileInfo implements repositoryServer.UpdateFileInfo
func (s *repositoryServer) UpdateFileInfo(ctx context.Context, in *server.UpdateFileInfoRequest) (*server.UpdateFileInfoResponse, error) {
	fileInfo := s.message.updateFileInfo.Input(in)
	fileInfo, err := s.service.UpdateFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.message.updateFileInfo.Output(fileInfo), nil
}

// DeleteFileInfo implements repositoryServer.DeleteFileInfo
func (s *repositoryServer) DeleteFileInfo(ctx context.Context, in *server.DeleteFileInfoRequest) (*server.DeleteFileInfoResponse, error) {
	fileInfo := s.message.deleteFileInfo.Input(in)
	fileInfo, err := s.service.DeleteFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.message.deleteFileInfo.Output(fileInfo), nil
}
