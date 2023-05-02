package table

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateTable struct {
}

func (x *CreateTable) Input(req *server.CreateTableRequest) *model.Table {
	return &model.Table{
		Title: req.Table.Title,
	}
}

func (x *CreateTable) Output(res *model.Table) *server.CreateTableResponse {
	return &server.CreateTableResponse{}
}

type ReadTable struct {
}

func (x *ReadTable) Input(req *server.ReadTableRequest) *model.Table {
	return &model.Table{
		ID: req.GetID(),
	}
}

func (x *ReadTable) Output(res *model.Table) *server.ReadTableResponse {
	return &server.ReadTableResponse{
		ID: res.ID,
		Table: &server.Table{
			Title: res.Title,
		},
	}
}

type UpdateTable struct {
}

func (x *UpdateTable) Input(req *server.UpdateTableRequest) *model.Table {
	return &model.Table{
		ID:    req.GetID(),
		Title: req.Table.Title,
	}
}

func (x *UpdateTable) Output(res *model.Table) *server.UpdateTableResponse {
	return &server.UpdateTableResponse{}
}

type DeleteTable struct {
}

func (x *DeleteTable) Input(req *server.DeleteTableRequest) *model.Table {
	return &model.Table{
		ID: req.GetID(),
	}
}

func (x *DeleteTable) Output(res *model.Table) *server.DeleteTableResponse {
	return &server.DeleteTableResponse{}
}