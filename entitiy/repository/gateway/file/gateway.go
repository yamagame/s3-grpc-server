package file

import (
	"sample/s3-grpc-server/entitiy/repository/gateway/table"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

func ToDomain(file *grpc_server.File) *model.File {
	tables := []*model.Table{}
	for _, v := range file.Tables {
		tables = append(tables, table.ToDomain(v))
	}
	return &model.File{
		ID:       file.Id,
		Filename: file.Filename,
		Attr: model.Attr{
			Owner: file.Owner,
		},
		Tables: tables,
	}
}

func ToInfra(file *model.File) *grpc_server.File {
	tables := []*grpc_server.Table{}
	for _, v := range file.Tables {
		tables = append(tables, table.ToInfra(v))
	}
	return &grpc_server.File{
		Id:       file.ID,
		Filename: file.Filename,
		Owner:    file.Attr.Owner,
		Tables:   tables,
	}
}
