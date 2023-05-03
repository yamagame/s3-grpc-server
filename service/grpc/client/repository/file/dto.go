package file

import (
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/infra/repository/table"
	server "sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/client/repository"
)

type CreateFile struct{}

func (x *CreateFile) Domain(req *model.File, call func(file *server.CreateFileRequest) (*server.CreateFileResponse, error)) (*model.File, error) {
	return repository.Domain[model.File, server.CreateFileRequest, server.CreateFileResponse](x, req, call)
}

func (x *CreateFile) Input(req *model.File) *server.CreateFileRequest {
	return &server.CreateFileRequest{
		File: &server.File{
			Filename: req.Filename,
		},
	}
}

func (x *CreateFile) Output(res *server.CreateFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type ReadFile struct{}

func (x *ReadFile) Domain(req *model.File, call func(file *server.ReadFileRequest) (*server.ReadFileResponse, error)) (*model.File, error) {
	return repository.Domain[model.File, server.ReadFileRequest, server.ReadFileResponse](x, req, call)
}

func (x *ReadFile) Input(req *model.File) *server.ReadFileRequest {
	return &server.ReadFileRequest{
		ID: req.ID,
	}
}

func (x *ReadFile) Output(res *server.ReadFileResponse) *model.File {
	tables := make([]*model.Table, 0)
	for _, v := range res.File.Tables {
		tables = append(tables, table.ToDomain(v))
	}
	return &model.File{
		ID:       res.GetID(),
		Filename: res.File.Filename,
		Tables:   tables,
	}
}

type UpdateFile struct{}

func (x *UpdateFile) Domain(req *model.File, call func(file *server.UpdateFileRequest) (*server.UpdateFileResponse, error)) (*model.File, error) {
	return repository.Domain[model.File, server.UpdateFileRequest, server.UpdateFileResponse](x, req, call)
}

func (x *UpdateFile) Input(req *model.File) *server.UpdateFileRequest {
	return &server.UpdateFileRequest{
		ID: req.ID,
		File: &server.File{
			Filename: req.Filename,
		},
	}
}

func (x *UpdateFile) Output(res *server.UpdateFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type DeleteFile struct{}

func (x *DeleteFile) Domain(req *model.File, call func(file *server.DeleteFileRequest) (*server.DeleteFileResponse, error)) (*model.File, error) {
	return repository.Domain[model.File, server.DeleteFileRequest, server.DeleteFileResponse](x, req, call)
}

func (x *DeleteFile) Input(req *model.File) *server.DeleteFileRequest {
	return &server.DeleteFileRequest{
		ID: req.ID,
	}
}

func (x *DeleteFile) Output(res *server.DeleteFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}
