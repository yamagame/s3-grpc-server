package file

import (
	"sample/s3-grpc-server/entitiy/repository/gateway/file"
	"sample/s3-grpc-server/entitiy/repository/gateway/table"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateFile struct{}

func (x *CreateFile) Input(req *model.File) *grpc_server.CreateFileRequest {
	tables := []*grpc_server.Table{}
	for _, v := range req.Tables {
		tables = append(tables, table.ToInfra(v))
	}
	return &grpc_server.CreateFileRequest{
		File: &grpc_server.File{
			Filename: req.Filename,
			Owner:    req.Attr.Owner,
			Tables:   tables,
		},
	}
}

func (x *CreateFile) Output(res *grpc_server.CreateFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type ReadFile struct{}

func (x *ReadFile) Input(req *model.File) *grpc_server.ReadFileRequest {
	return &grpc_server.ReadFileRequest{
		ID: req.ID,
	}
}

func (x *ReadFile) Output(res *grpc_server.ReadFileResponse) *model.File {
	tables := make([]*model.Table, 0)
	for _, v := range res.File.Tables {
		tables = append(tables, table.ToDomain(v))
	}
	return &model.File{
		ID:       res.GetID(),
		Filename: res.File.Filename,
		Attr: model.Attr{
			ID:    res.GetID(),
			Owner: res.File.Owner,
		},
		Tables: tables,
	}
}

type UpdateFile struct{}

func (x *UpdateFile) Input(req *model.File) *grpc_server.UpdateFileRequest {
	return &grpc_server.UpdateFileRequest{
		ID: req.ID,
		File: &grpc_server.File{
			Filename: req.Filename,
			Owner:    req.Attr.Owner,
		},
	}
}

func (x *UpdateFile) Output(res *grpc_server.UpdateFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type DeleteFile struct{}

func (x *DeleteFile) Input(req *model.File) *grpc_server.DeleteFileRequest {
	return &grpc_server.DeleteFileRequest{
		ID: req.ID,
	}
}

func (x *DeleteFile) Output(res *grpc_server.DeleteFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type ListFile struct{}

func (x *ListFile) Input(req *model.File) *grpc_server.ListFileRequest {
	return &grpc_server.ListFileRequest{
		//
	}
}

func (x *ListFile) Output(res *grpc_server.ListFileResponse) []*model.File {
	ret := []*model.File{}
	for _, v := range res.Files {
		ret = append(ret, file.ToDomain(v))
	}
	return ret
}
