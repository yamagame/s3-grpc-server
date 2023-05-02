package fileinfo

import (
	"context"
	"sample/s3-grpc-server/grpc_server/service/repository/dto"
	"sample/s3-grpc-server/infra/repository/fileinfo"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	createFileInfo dto.CreateFilieInfo
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

// CreateFileInfo implements RepositoryServer.CreateFileInfo
func (s *FileInfoServer) CreateFileInfo(ctx context.Context, in *server.CreateFileInfoRequest) (*server.CreateFileInfoResponse, error) {
	fileInfo := s.createFileInfo.Input(in)
	fileInfo, err := s.service.CreateFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.createFileInfo.Output(fileInfo), nil
}

// ReadFileInfo implements RepositoryServer.ReadFileInfo
func (s *FileInfoServer) ReadFileInfo(ctx context.Context, in *server.ReadFileInfoRequest) (*server.ReadFileInfoResponse, error) {
	fileInfo := s.readFileInfo.Input(in)
	fileInfo, err := s.service.ReadFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.readFileInfo.Output(fileInfo), nil
}

// UpdateFileInfo implements RepositoryServer.UpdateFileInfo
func (s *FileInfoServer) UpdateFileInfo(ctx context.Context, in *server.UpdateFileInfoRequest) (*server.UpdateFileInfoResponse, error) {
	fileInfo := s.updateFileInfo.Input(in)
	fileInfo, err := s.service.UpdateFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.updateFileInfo.Output(fileInfo), nil
}

// DeleteFileInfo implements RepositoryServer.DeleteFileInfo
func (s *FileInfoServer) DeleteFileInfo(ctx context.Context, in *server.DeleteFileInfoRequest) (*server.DeleteFileInfoResponse, error) {
	fileInfo := s.deleteFileInfo.Input(in)
	fileInfo, err := s.service.DeleteFileInfo(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.deleteFileInfo.Output(fileInfo), nil
}
