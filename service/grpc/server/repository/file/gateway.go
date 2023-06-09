package file

import (
	"sample/s3-grpc-server/entitiy/repository/gateway/file"
	"sample/s3-grpc-server/entitiy/repository/gateway/table"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"

	"time"
)

type CreateFile struct {
}

func (x *CreateFile) Input(req *grpc_server.CreateFileRequest) *model.File {
	tables := make([]*model.Table, 0)
	for _, v := range req.File.Tables {
		tables = append(tables, table.ToDomain(v))
	}
	return &model.File{
		Filename: req.File.Filename,
		Attr: model.Attr{
			Owner: req.File.Owner,
		},
		Tables:    tables,
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
			Owner:    res.Attr.Owner,
			Tables:   tables,
		},
	}
}

type UpdateFile struct {
}

func (x *UpdateFile) Input(req *grpc_server.UpdateFileRequest) *model.File {
	return &model.File{
		ID:       req.GetID(),
		Filename: req.File.Filename,
		Attr: model.Attr{
			ID:    req.GetID(),
			Owner: req.File.Owner,
		},
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

type ListFile struct {
}

func (x *ListFile) Input(req *grpc_server.ListFileRequest) *model.File {
	return &model.File{}
}

func (x *ListFile) Output(res []*model.File) *grpc_server.ListFileResponse {
	files := make([]*grpc_server.File, 0)
	for _, v := range res {
		files = append(files, file.ToInfra(v))
	}
	return &grpc_server.ListFileResponse{
		Files: files,
	}
}
