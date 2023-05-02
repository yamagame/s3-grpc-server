package file

import (
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/infra/repository/table"
	server "sample/s3-grpc-server/proto/grpc_server"
	"time"
)

type CreateFile struct {
}

func (x *CreateFile) Domain(req *server.CreateFileRequest, call func(table *model.File) (*model.File, error)) (*server.CreateFileResponse, error) {
	file := x.input(req)
	file, err := call(file)
	if err != nil {
		return nil, err
	}
	return x.output(file), nil
}

func (x *CreateFile) input(req *server.CreateFileRequest) *model.File {
	return &model.File{
		Filename:  req.File.Filename,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (x *CreateFile) output(res *model.File) *server.CreateFileResponse {
	return &server.CreateFileResponse{
		ID: res.ID,
	}
}

type ReadFile struct {
}

func (x *ReadFile) Domain(req *server.ReadFileRequest, call func(table *model.File) (*model.File, error)) (*server.ReadFileResponse, error) {
	file := x.input(req)
	file, err := call(file)
	if err != nil {
		return nil, err
	}
	return x.output(file), nil
}

func (x *ReadFile) input(req *server.ReadFileRequest) *model.File {
	return &model.File{
		ID: req.GetID(),
	}
}

func (x *ReadFile) output(res *model.File) *server.ReadFileResponse {
	tables := make([]*server.Table, 0)
	for _, v := range res.Tables {
		tables = append(tables, table.ToInfra(v))
	}
	return &server.ReadFileResponse{
		ID: res.ID,
		File: &server.File{
			Filename: res.Filename,
			Tables:   tables,
		},
	}
}

type UpdateFile struct {
}

func (x *UpdateFile) Domain(req *server.UpdateFileRequest, call func(table *model.File) (*model.File, error)) (*server.UpdateFileResponse, error) {
	file := x.input(req)
	file, err := call(file)
	if err != nil {
		return nil, err
	}
	return x.output(file), nil
}

func (x *UpdateFile) input(req *server.UpdateFileRequest) *model.File {
	return &model.File{
		ID:        req.GetID(),
		Filename:  req.File.Filename,
		UpdatedAt: time.Now(),
	}
}

func (x *UpdateFile) output(res *model.File) *server.UpdateFileResponse {
	return &server.UpdateFileResponse{
		ID: res.ID,
	}
}

type DeleteFile struct {
}

func (x *DeleteFile) Domain(req *server.DeleteFileRequest, call func(table *model.File) (*model.File, error)) (*server.DeleteFileResponse, error) {
	file := x.input(req)
	file, err := call(file)
	if err != nil {
		return nil, err
	}
	return x.output(file), nil
}

func (x *DeleteFile) input(req *server.DeleteFileRequest) *model.File {
	return &model.File{
		ID: req.GetID(),
	}
}

func (x *DeleteFile) output(res *model.File) *server.DeleteFileResponse {
	return &server.DeleteFileResponse{
		ID: res.ID,
	}
}
