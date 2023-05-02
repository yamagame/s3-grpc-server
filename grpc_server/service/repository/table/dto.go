package table

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
	"time"
)

type CreateTable struct {
}

func (x *CreateTable) Input(req *server.CreateTableRequest) *model.Table {
	return &model.Table{
		FileID:    req.Table.FileId,
		Title:     req.Table.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (x *CreateTable) Output(res *model.Table) *server.CreateTableResponse {
	return &server.CreateTableResponse{
		ID: res.ID,
	}
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
		ID:        req.GetID(),
		Title:     req.Table.Title,
		UpdatedAt: time.Now(),
	}
}

func (x *UpdateTable) Output(res *model.Table) *server.UpdateTableResponse {
	return &server.UpdateTableResponse{
		ID: res.ID,
	}
}

type DeleteTable struct {
}

func (x *DeleteTable) Input(req *server.DeleteTableRequest) *model.Table {
	return &model.Table{
		ID: req.GetID(),
	}
}

func (x *DeleteTable) Output(res *model.Table) *server.DeleteTableResponse {
	return &server.DeleteTableResponse{
		ID: res.ID,
	}
}
