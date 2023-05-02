package file

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateFile struct {
}

func (x *CreateFile) Input(req *server.CreateFileRequest) *model.File {
	return &model.File{
		Filename: req.File.Filename,
	}
}

func (x *CreateFile) Output(res *model.File) *server.CreateFileResponse {
	return &server.CreateFileResponse{
		ID: res.ID,
	}
}

type ReadFile struct {
}

func (x *ReadFile) Input(req *server.ReadFileRequest) *model.File {
	return &model.File{
		ID: req.GetID(),
	}
}

func (x *ReadFile) Output(res *model.File) *server.ReadFileResponse {
	return &server.ReadFileResponse{
		ID: res.ID,
		File: &server.File{
			Filename: res.Filename,
		},
	}
}

type UpdateFile struct {
}

func (x *UpdateFile) Input(req *server.UpdateFileRequest) *model.File {
	return &model.File{
		ID:       req.GetID(),
		Filename: req.File.Filename,
	}
}

func (x *UpdateFile) Output(res *model.File) *server.UpdateFileResponse {
	return &server.UpdateFileResponse{
		ID: res.ID,
	}
}

type DeleteFile struct {
}

func (x *DeleteFile) Input(req *server.DeleteFileRequest) *model.File {
	return &model.File{
		ID: req.GetID(),
	}
}

func (x *DeleteFile) Output(res *model.File) *server.DeleteFileResponse {
	return &server.DeleteFileResponse{
		ID: res.ID,
	}
}
