package table

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

func ToDomain(table *server.Table) *model.Table {
	return &model.Table{
		ID:    table.Id,
		Title: table.Title,
	}
}

func ToInfra(table *model.Table) *server.Table {
	return &server.Table{
		Id:    table.ID,
		Title: table.Title,
	}
}
