package repository

import (
	"context"
	"fmt"
	"sample/s3-grpc-server/infra/repository"
	server "sample/s3-grpc-server/proto/grpc-server"

	"gorm.io/gorm"
)

type repositoryServerDomain struct {
	createFileInfo repository.RepositoryCreateFilieInfoServerData
	readFileInfo   repository.RepositoryReadFileInfoServerData
	updateFileInfo repository.RepositoryUpdateFileInfoServerData
	deleteFileInfo repository.RepositoryDeleteFileInfoServerData
}

type repositoryServer struct {
	service *repositoryService
	domain  *repositoryServerDomain
	server.UnimplementedRepositoryServer
}

func NewRepositoryServer(db *gorm.DB) *repositoryServer {
	return &repositoryServer{
		domain:  &repositoryServerDomain{},
		service: NewRepositoryService(db),
	}
}

// CreateFileInfo implements repositoryServer.CreateFileInfo
func (s *repositoryServer) CreateFileInfo(ctx context.Context, in *server.CreateFileInfoRequest) (*server.CreateFileInfoResponse, error) {
	fmt.Println(in)
	fileInfo := s.domain.createFileInfo.Input(in)
	fmt.Println(in)
	fileInfo, err := s.service.CreateFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.domain.createFileInfo.Output(fileInfo), nil
}

// ReadFileInfo implements repositoryServer.ReadFileInfo
func (s *repositoryServer) ReadFileInfo(ctx context.Context, in *server.ReadFileInfoRequest) (*server.ReadFileInfoResponse, error) {
	fileInfo := s.domain.readFileInfo.Input(in)
	fileInfo, err := s.service.ReadFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.domain.readFileInfo.Output(fileInfo), nil
}

// UpdateFileInfo implements repositoryServer.UpdateFileInfo
func (s *repositoryServer) UpdateFileInfo(ctx context.Context, in *server.UpdateFileInfoRequest) (*server.UpdateFileInfoResponse, error) {
	fileInfo := s.domain.updateFileInfo.Input(in)
	fileInfo, err := s.service.UpdateFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.domain.updateFileInfo.Output(fileInfo), nil
}

// DeleteFileInfo implements repositoryServer.DeleteFileInfo
func (s *repositoryServer) DeleteFileInfo(ctx context.Context, in *server.DeleteFileInfoRequest) (*server.DeleteFileInfoResponse, error) {
	fileInfo := s.domain.deleteFileInfo.Input(in)
	fileInfo, err := s.service.DeleteFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.domain.deleteFileInfo.Output(fileInfo), nil
}
