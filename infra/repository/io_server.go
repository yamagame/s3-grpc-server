package repository

import (
	server "sample/s3-grpc-server/proto/grpc_server"
)

type RepositoryCreateFilieInfoServerData struct {
}

func (x *RepositoryCreateFilieInfoServerData) Input(req *server.CreateFileInfoRequest) *FileInfo {
	return &FileInfo{
		Filename: req.File.Filename,
	}
}

func (x *RepositoryCreateFilieInfoServerData) Output(res *FileInfo) *server.CreateFileInfoResponse {
	return &server.CreateFileInfoResponse{}
}

type RepositoryReadFileInfoServerData struct {
}

func (x *RepositoryReadFileInfoServerData) Input(req *server.ReadFileInfoRequest) *FileInfo {
	return &FileInfo{
		ID: req.GetID(),
	}
}

func (x *RepositoryReadFileInfoServerData) Output(res *FileInfo) *server.ReadFileInfoResponse {
	return &server.ReadFileInfoResponse{
		ID: res.ID,
		File: &server.FileInfo{
			Filename: res.Filename,
		},
	}
}

type RepositoryUpdateFileInfoServerData struct {
}

func (x *RepositoryUpdateFileInfoServerData) Input(req *server.UpdateFileInfoRequest) *FileInfo {
	return &FileInfo{
		ID:       req.GetID(),
		Filename: req.File.Filename,
	}
}

func (x *RepositoryUpdateFileInfoServerData) Output(res *FileInfo) *server.UpdateFileInfoResponse {
	return &server.UpdateFileInfoResponse{}
}

type RepositoryDeleteFileInfoServerData struct {
}

func (x *RepositoryDeleteFileInfoServerData) Input(req *server.DeleteFileInfoRequest) *FileInfo {
	return &FileInfo{
		ID: req.GetID(),
	}
}

func (x *RepositoryDeleteFileInfoServerData) Output(res *FileInfo) *server.DeleteFileInfoResponse {
	return &server.DeleteFileInfoResponse{}
}
