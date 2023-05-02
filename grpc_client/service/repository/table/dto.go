package table

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateTableDTO struct {
}

func (x *CreateTableDTO) Input(req *model.Table) *server.CreateTableRequest {
	return &server.CreateTableRequest{
		Table: &server.Table{
			FileId: req.FileID,
			Title:  req.Title,
		},
	}
}

func (x *CreateTableDTO) Output(res *server.CreateTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}

type ReadTableDTO struct {
}

func (x *ReadTableDTO) Input(req *model.Table) *server.ReadTableRequest {
	return &server.ReadTableRequest{
		ID: req.ID,
	}
}

func (x *ReadTableDTO) Output(res *server.ReadTableResponse) *model.Table {
	return &model.Table{
		ID:    res.GetID(),
		Title: res.Table.Title,
	}
}

type UpdateTableDTO struct {
}

func (x *UpdateTableDTO) Input(req *model.Table) *server.UpdateTableRequest {
	return &server.UpdateTableRequest{
		ID: req.ID,
		Table: &server.Table{
			Title: req.Title,
		},
	}
}

func (x *UpdateTableDTO) Output(res *server.UpdateTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}

type DeleteTableDTO struct {
}

func (x *DeleteTableDTO) Input(req *model.Table) *server.DeleteTableRequest {
	return &server.DeleteTableRequest{
		ID: req.ID,
	}
}

func (x *DeleteTableDTO) Output(res *server.DeleteTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}
