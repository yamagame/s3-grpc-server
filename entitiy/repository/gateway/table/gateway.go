package table

import (
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

func ToDomain(table *grpc_server.Table) *model.Table {
	return &model.Table{
		ID:    table.Id,
		Title: table.Title,
	}
}

func ToInfra(table *model.Table) *grpc_server.Table {
	return &grpc_server.Table{
		Id:    table.ID,
		Title: table.Title,
	}
}
