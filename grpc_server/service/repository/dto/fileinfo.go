package dto

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateFilieInfo struct {
}

func (x *CreateFilieInfo) Input(req *server.CreateFileInfoRequest) *model.FileInfo {
	return &model.FileInfo{
		Filename: req.File.Filename,
	}
}

func (x *CreateFilieInfo) Output(res *model.FileInfo) *server.CreateFileInfoResponse {
	return &server.CreateFileInfoResponse{}
}

type ReadFileInfo struct {
}

func (x *ReadFileInfo) Input(req *server.ReadFileInfoRequest) *model.FileInfo {
	return &model.FileInfo{
		ID: req.GetID(),
	}
}

func (x *ReadFileInfo) Output(res *model.FileInfo) *server.ReadFileInfoResponse {
	return &server.ReadFileInfoResponse{
		ID: res.ID,
		File: &server.FileInfo{
			Filename: res.Filename,
		},
	}
}

type UpdateFileInfo struct {
}

func (x *UpdateFileInfo) Input(req *server.UpdateFileInfoRequest) *model.FileInfo {
	return &model.FileInfo{
		ID:       req.GetID(),
		Filename: req.File.Filename,
	}
}

func (x *UpdateFileInfo) Output(res *model.FileInfo) *server.UpdateFileInfoResponse {
	return &server.UpdateFileInfoResponse{}
}

type DeleteFileInfo struct {
}

func (x *DeleteFileInfo) Input(req *server.DeleteFileInfoRequest) *model.FileInfo {
	return &model.FileInfo{
		ID: req.GetID(),
	}
}

func (x *DeleteFileInfo) Output(res *model.FileInfo) *server.DeleteFileInfoResponse {
	return &server.DeleteFileInfoResponse{}
}
