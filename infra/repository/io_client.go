package repository

import (
	server "sample/s3-grpc-server/proto/grpc_server"
)

type RepositoryCreateFilieInfoClientData struct {
}

func (x *RepositoryCreateFilieInfoClientData) Input(req *FileInfo) *server.CreateFileInfoRequest {
	return &server.CreateFileInfoRequest{
		File: &server.FileInfo{
			Filename: req.Filename,
		},
	}
}

func (x *RepositoryCreateFilieInfoClientData) Output(res *server.CreateFileInfoResponse) *FileInfo {
	return &FileInfo{
		ID: res.GetID(),
	}
}

type RepositoryReadFileInfoClientData struct {
}

func (x *RepositoryReadFileInfoClientData) Input(req *FileInfo) *server.ReadFileInfoRequest {
	return &server.ReadFileInfoRequest{
		ID: req.ID,
	}
}

func (x *RepositoryReadFileInfoClientData) Output(res *server.ReadFileInfoResponse) *FileInfo {
	return &FileInfo{
		ID:       res.GetID(),
		Filename: res.File.Filename,
	}
}

type RepositoryUpdateFileInfoClientData struct {
}

func (x *RepositoryUpdateFileInfoClientData) Input(req *FileInfo) *server.UpdateFileInfoRequest {
	return &server.UpdateFileInfoRequest{
		ID: req.ID,
		File: &server.FileInfo{
			Filename: req.Filename,
		},
	}
}

func (x *RepositoryUpdateFileInfoClientData) Output(res *server.UpdateFileInfoResponse) *FileInfo {
	return &FileInfo{
		ID: res.GetID(),
	}
}

type RepositoryDeleteFileInfoClientData struct {
}

func (x *RepositoryDeleteFileInfoClientData) Input(req *FileInfo) *server.DeleteFileInfoRequest {
	return &server.DeleteFileInfoRequest{
		ID: req.ID,
	}
}

func (x *RepositoryDeleteFileInfoClientData) Output(res *server.DeleteFileInfoResponse) *FileInfo {
	return &FileInfo{
		ID: res.GetID(),
	}
}
