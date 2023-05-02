package repository

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateFilieInfoDTO struct {
}

func (x *CreateFilieInfoDTO) Input(req *model.FileInfo) *server.CreateFileInfoRequest {
	return &server.CreateFileInfoRequest{
		File: &server.FileInfo{
			Filename: req.Filename,
		},
	}
}

func (x *CreateFilieInfoDTO) Output(res *server.CreateFileInfoResponse) *model.FileInfo {
	return &model.FileInfo{
		ID: res.GetID(),
	}
}

type ReadFileInfoDTO struct {
}

func (x *ReadFileInfoDTO) Input(req *model.FileInfo) *server.ReadFileInfoRequest {
	return &server.ReadFileInfoRequest{
		ID: req.ID,
	}
}

func (x *ReadFileInfoDTO) Output(res *server.ReadFileInfoResponse) *model.FileInfo {
	return &model.FileInfo{
		ID:       res.GetID(),
		Filename: res.File.Filename,
	}
}

type UpdateFileInfoDTO struct {
}

func (x *UpdateFileInfoDTO) Input(req *model.FileInfo) *server.UpdateFileInfoRequest {
	return &server.UpdateFileInfoRequest{
		ID: req.ID,
		File: &server.FileInfo{
			Filename: req.Filename,
		},
	}
}

func (x *UpdateFileInfoDTO) Output(res *server.UpdateFileInfoResponse) *model.FileInfo {
	return &model.FileInfo{
		ID: res.GetID(),
	}
}

type DeleteFileInfoDTO struct {
}

func (x *DeleteFileInfoDTO) Input(req *model.FileInfo) *server.DeleteFileInfoRequest {
	return &server.DeleteFileInfoRequest{
		ID: req.ID,
	}
}

func (x *DeleteFileInfoDTO) Output(res *server.DeleteFileInfoResponse) *model.FileInfo {
	return &model.FileInfo{
		ID: res.GetID(),
	}
}
