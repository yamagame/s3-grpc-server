package file

import (
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/libs/dto"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateFile struct{}

func (x *CreateFile) ToInfra(req *model.File, call func(file *grpc_server.CreateFileRequest) (*grpc_server.CreateFileResponse, error)) (*model.File, error) {
	return dto.ToInfra[model.File, grpc_server.CreateFileRequest, grpc_server.CreateFileResponse](x, req, call)
}

func (x *CreateFile) Input(req *model.File) *grpc_server.CreateFileRequest {
	return &grpc_server.CreateFileRequest{
		File: &grpc_server.File{
			Filename: req.Filename,
		},
	}
}

func (x *CreateFile) Output(res *grpc_server.CreateFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type ReadFile struct{}

func (x *ReadFile) ToInfra(req *model.File, call func(file *grpc_server.ReadFileRequest) (*grpc_server.ReadFileResponse, error)) (*model.File, error) {
	return dto.ToInfra[model.File, grpc_server.ReadFileRequest, grpc_server.ReadFileResponse](x, req, call)
}

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
		Tables:   tables,
	}
}

type UpdateFile struct{}

func (x *UpdateFile) ToInfra(req *model.File, call func(file *grpc_server.UpdateFileRequest) (*grpc_server.UpdateFileResponse, error)) (*model.File, error) {
	return dto.ToInfra[model.File, grpc_server.UpdateFileRequest, grpc_server.UpdateFileResponse](x, req, call)
}

func (x *UpdateFile) Input(req *model.File) *grpc_server.UpdateFileRequest {
	return &grpc_server.UpdateFileRequest{
		ID: req.ID,
		File: &grpc_server.File{
			Filename: req.Filename,
		},
	}
}

func (x *UpdateFile) Output(res *grpc_server.UpdateFileResponse) *model.File {
	return &model.File{
		ID: res.GetID(),
	}
}

type DeleteFile struct{}

func (x *DeleteFile) ToInfra(req *model.File, call func(file *grpc_server.DeleteFileRequest) (*grpc_server.DeleteFileResponse, error)) (*model.File, error) {
	return dto.ToInfra[model.File, grpc_server.DeleteFileRequest, grpc_server.DeleteFileResponse](x, req, call)
}

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
