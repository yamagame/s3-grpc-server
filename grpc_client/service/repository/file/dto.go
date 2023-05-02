package file

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateFileDTO struct {
}

func (x *CreateFileDTO) Input(req *model.File) *server.CreateFileRequest {
	return &server.CreateFileRequest{
		File: &server.File{
			Filename: req.Filename,
		},
	}
}

func (x *CreateFileDTO) Output(res *server.CreateFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type ReadFileDTO struct {
}

func (x *ReadFileDTO) Input(req *model.File) *server.ReadFileRequest {
	return &server.ReadFileRequest{
		ID: req.ID,
	}
}

func (x *ReadFileDTO) Output(res *server.ReadFileResponse) *model.File {
	return &model.File{
		ID:       res.GetID(),
		Filename: res.File.Filename,
	}
}

type UpdateFileDTO struct {
}

func (x *UpdateFileDTO) Input(req *model.File) *server.UpdateFileRequest {
	return &server.UpdateFileRequest{
		ID: req.ID,
		File: &server.File{
			Filename: req.Filename,
		},
	}
}

func (x *UpdateFileDTO) Output(res *server.UpdateFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type DeleteFileDTO struct {
}

func (x *DeleteFileDTO) Input(req *model.File) *server.DeleteFileRequest {
	return &server.DeleteFileRequest{
		ID: req.ID,
	}
}

func (x *DeleteFileDTO) Output(res *server.DeleteFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}
