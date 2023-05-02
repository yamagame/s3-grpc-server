package repository

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type RepositoryCreateFilieInfoClientMessage struct {
}

func (x *RepositoryCreateFilieInfoClientMessage) Input(req *model.FileInfo) *server.CreateFileInfoRequest {
	return &server.CreateFileInfoRequest{
		File: &server.FileInfo{
			Filename: req.Filename,
		},
	}
}

func (x *RepositoryCreateFilieInfoClientMessage) Output(res *server.CreateFileInfoResponse) *model.FileInfo {
	return &model.FileInfo{
		ID: res.GetID(),
	}
}

type RepositoryReadFileInfoClientMessage struct {
}

func (x *RepositoryReadFileInfoClientMessage) Input(req *model.FileInfo) *server.ReadFileInfoRequest {
	return &server.ReadFileInfoRequest{
		ID: req.ID,
	}
}

func (x *RepositoryReadFileInfoClientMessage) Output(res *server.ReadFileInfoResponse) *model.FileInfo {
	return &model.FileInfo{
		ID:       res.GetID(),
		Filename: res.File.Filename,
	}
}

type RepositoryUpdateFileInfoClientMessage struct {
}

func (x *RepositoryUpdateFileInfoClientMessage) Input(req *model.FileInfo) *server.UpdateFileInfoRequest {
	return &server.UpdateFileInfoRequest{
		ID: req.ID,
		File: &server.FileInfo{
			Filename: req.Filename,
		},
	}
}

func (x *RepositoryUpdateFileInfoClientMessage) Output(res *server.UpdateFileInfoResponse) *model.FileInfo {
	return &model.FileInfo{
		ID: res.GetID(),
	}
}

type RepositoryDeleteFileInfoClientMessage struct {
}

func (x *RepositoryDeleteFileInfoClientMessage) Input(req *model.FileInfo) *server.DeleteFileInfoRequest {
	return &server.DeleteFileInfoRequest{
		ID: req.ID,
	}
}

func (x *RepositoryDeleteFileInfoClientMessage) Output(res *server.DeleteFileInfoResponse) *model.FileInfo {
	return &model.FileInfo{
		ID: res.GetID(),
	}
}
