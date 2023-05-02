package repository

import (
	"sample/s3-grpc-server/infra/repository/fileinfo/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateFilieInfoDTO struct {
}

func (x *CreateFilieInfoDTO) Input(req *server.CreateFileInfoRequest) *model.FileInfo {
	return &model.FileInfo{
		Filename: req.File.Filename,
	}
}

func (x *CreateFilieInfoDTO) Output(res *model.FileInfo) *server.CreateFileInfoResponse {
	return &server.CreateFileInfoResponse{}
}

type ReadFileInfoDTO struct {
}

func (x *ReadFileInfoDTO) Input(req *server.ReadFileInfoRequest) *model.FileInfo {
	return &model.FileInfo{
		ID: req.GetID(),
	}
}

func (x *ReadFileInfoDTO) Output(res *model.FileInfo) *server.ReadFileInfoResponse {
	return &server.ReadFileInfoResponse{
		ID: res.ID,
		File: &server.FileInfo{
			Filename: res.Filename,
		},
	}
}

type UpdateFileInfoDTO struct {
}

func (x *UpdateFileInfoDTO) Input(req *server.UpdateFileInfoRequest) *model.FileInfo {
	return &model.FileInfo{
		ID:       req.GetID(),
		Filename: req.File.Filename,
	}
}

func (x *UpdateFileInfoDTO) Output(res *model.FileInfo) *server.UpdateFileInfoResponse {
	return &server.UpdateFileInfoResponse{}
}

type DeleteFileInfoDTO struct {
}

func (x *DeleteFileInfoDTO) Input(req *server.DeleteFileInfoRequest) *model.FileInfo {
	return &model.FileInfo{
		ID: req.GetID(),
	}
}

func (x *DeleteFileInfoDTO) Output(res *model.FileInfo) *server.DeleteFileInfoResponse {
	return &server.DeleteFileInfoResponse{}
}
