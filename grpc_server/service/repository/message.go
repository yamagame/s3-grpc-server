package repository

import (
	"sample/s3-grpc-server/infra/repository"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type RepositoryCreateFilieInfoServerMessage struct {
}

func (x *RepositoryCreateFilieInfoServerMessage) Input(req *server.CreateFileInfoRequest) *repository.FileInfo {
	return &repository.FileInfo{
		Filename: req.File.Filename,
	}
}

func (x *RepositoryCreateFilieInfoServerMessage) Output(res *repository.FileInfo) *server.CreateFileInfoResponse {
	return &server.CreateFileInfoResponse{}
}

type RepositoryReadFileInfoServerMessage struct {
}

func (x *RepositoryReadFileInfoServerMessage) Input(req *server.ReadFileInfoRequest) *repository.FileInfo {
	return &repository.FileInfo{
		ID: req.GetID(),
	}
}

func (x *RepositoryReadFileInfoServerMessage) Output(res *repository.FileInfo) *server.ReadFileInfoResponse {
	return &server.ReadFileInfoResponse{
		ID: res.ID,
		File: &server.FileInfo{
			Filename: res.Filename,
		},
	}
}

type RepositoryUpdateFileInfoServerMessage struct {
}

func (x *RepositoryUpdateFileInfoServerMessage) Input(req *server.UpdateFileInfoRequest) *repository.FileInfo {
	return &repository.FileInfo{
		ID:       req.GetID(),
		Filename: req.File.Filename,
	}
}

func (x *RepositoryUpdateFileInfoServerMessage) Output(res *repository.FileInfo) *server.UpdateFileInfoResponse {
	return &server.UpdateFileInfoResponse{}
}

type RepositoryDeleteFileInfoServerMessage struct {
}

func (x *RepositoryDeleteFileInfoServerMessage) Input(req *server.DeleteFileInfoRequest) *repository.FileInfo {
	return &repository.FileInfo{
		ID: req.GetID(),
	}
}

func (x *RepositoryDeleteFileInfoServerMessage) Output(res *repository.FileInfo) *server.DeleteFileInfoResponse {
	return &server.DeleteFileInfoResponse{}
}
