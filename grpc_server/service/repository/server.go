package repository

import (
	"context"
	"sample/s3-grpc-server/grpc_server/service/repository/fileinfo"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type repositoryServer struct {
	fileinfo *fileinfo.FileInfoServer
	server.UnimplementedRepositoryServer
}

func NewRepositoryServer(fileinfo *fileinfo.FileInfoServer) *repositoryServer {
	return &repositoryServer{
		fileinfo: fileinfo,
	}
}

// CreateFileInfo implements repositoryServer.CreateFileInfo
func (s *repositoryServer) CreateFileInfo(ctx context.Context, in *server.CreateFileInfoRequest) (*server.CreateFileInfoResponse, error) {
	return s.fileinfo.CreateFileInfo(ctx, in)
}

// ReadFileInfo implements repositoryServer.ReadFileInfo
func (s *repositoryServer) ReadFileInfo(ctx context.Context, in *server.ReadFileInfoRequest) (*server.ReadFileInfoResponse, error) {
	return s.fileinfo.ReadFileInfo(ctx, in)
}

// UpdateFileInfo implements repositoryServer.UpdateFileInfo
func (s *repositoryServer) UpdateFileInfo(ctx context.Context, in *server.UpdateFileInfoRequest) (*server.UpdateFileInfoResponse, error) {
	return s.fileinfo.UpdateFileInfo(ctx, in)
}

// DeleteFileInfo implements repositoryServer.DeleteFileInfo
func (s *repositoryServer) DeleteFileInfo(ctx context.Context, in *server.DeleteFileInfoRequest) (*server.DeleteFileInfoResponse, error) {
	return s.fileinfo.DeleteFileInfo(ctx, in)
}
