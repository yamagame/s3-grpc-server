package file

import (
	"sample/s3-grpc-server/entitiy/repository/gateway/table"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"

	"time"
)

type CreateFile struct {
}

func (x *CreateFile) Input(req *grpc_server.CreateFileRequest) *model.File {
	return &model.File{
		Filename:  req.File.Filename,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (x *CreateFile) Output(res *model.File) *grpc_server.CreateFileResponse {
	return &grpc_server.CreateFileResponse{
		ID: res.ID,
	}
}

type ReadFile struct {
}

func (x *ReadFile) Input(req *grpc_server.ReadFileRequest) *model.File {
	return &model.File{
		ID: req.GetID(),
	}
}

func (x *ReadFile) Output(res *model.File) *grpc_server.ReadFileResponse {
	tables := make([]*grpc_server.Table, 0)
	for _, v := range res.Tables {
		tables = append(tables, table.ToInfra(v))
	}
	return &grpc_server.ReadFileResponse{
		ID: res.ID,
		File: &grpc_server.File{
			Filename: res.Filename,
			Tables:   tables,
		},
	}
}

type UpdateFile struct {
}

func (x *UpdateFile) Input(req *grpc_server.UpdateFileRequest) *model.File {
	return &model.File{
		ID:        req.GetID(),
		Filename:  req.File.Filename,
		UpdatedAt: time.Now(),
	}
}

func (x *UpdateFile) Output(res *model.File) *grpc_server.UpdateFileResponse {
	return &grpc_server.UpdateFileResponse{
		ID: res.ID,
	}
}

type DeleteFile struct {
}

func (x *DeleteFile) Input(req *grpc_server.DeleteFileRequest) *model.File {
	return &model.File{
		ID: req.GetID(),
	}
}

func (x *DeleteFile) Output(res *model.File) *grpc_server.DeleteFileResponse {
	return &grpc_server.DeleteFileResponse{
		ID: res.ID,
	}
}
