package repository

import (
	"sample/s3-grpc-server/infra/repository"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type RepositoryCreateFilieInfoClientMessage struct {
}

func (x *RepositoryCreateFilieInfoClientMessage) Input(req *repository.FileInfo) *server.CreateFileInfoRequest {
	return &server.CreateFileInfoRequest{
		File: &server.FileInfo{
			Filename: req.Filename,
		},
	}
}

func (x *RepositoryCreateFilieInfoClientMessage) Output(res *server.CreateFileInfoResponse) *repository.FileInfo {
	return &repository.FileInfo{
		ID: res.GetID(),
	}
}

type RepositoryReadFileInfoClientMessage struct {
}

func (x *RepositoryReadFileInfoClientMessage) Input(req *repository.FileInfo) *server.ReadFileInfoRequest {
	return &server.ReadFileInfoRequest{
		ID: req.ID,
	}
}

func (x *RepositoryReadFileInfoClientMessage) Output(res *server.ReadFileInfoResponse) *repository.FileInfo {
	return &repository.FileInfo{
		ID:       res.GetID(),
		Filename: res.File.Filename,
	}
}

type RepositoryUpdateFileInfoClientMessage struct {
}

func (x *RepositoryUpdateFileInfoClientMessage) Input(req *repository.FileInfo) *server.UpdateFileInfoRequest {
	return &server.UpdateFileInfoRequest{
		ID: req.ID,
		File: &server.FileInfo{
			Filename: req.Filename,
		},
	}
}

func (x *RepositoryUpdateFileInfoClientMessage) Output(res *server.UpdateFileInfoResponse) *repository.FileInfo {
	return &repository.FileInfo{
		ID: res.GetID(),
	}
}

type RepositoryDeleteFileInfoClientMessage struct {
}

func (x *RepositoryDeleteFileInfoClientMessage) Input(req *repository.FileInfo) *server.DeleteFileInfoRequest {
	return &server.DeleteFileInfoRequest{
		ID: req.ID,
	}
}

func (x *RepositoryDeleteFileInfoClientMessage) Output(res *server.DeleteFileInfoResponse) *repository.FileInfo {
	return &repository.FileInfo{
		ID: res.GetID(),
	}
}
