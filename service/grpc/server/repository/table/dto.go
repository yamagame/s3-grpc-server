package table

import (
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/server/repository"
	"time"
)

type CreateTable struct {
}

func (x *CreateTable) Domain(req *server.CreateTableRequest, call func(table *model.Table) (*model.Table, error)) (*server.CreateTableResponse, error) {
	return repository.Domain[model.Table, server.CreateTableRequest, server.CreateTableResponse](x, req, call)
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

func (x *ReadTable) Domain(req *server.ReadTableRequest, call func(table *model.Table) (*model.Table, error)) (*server.ReadTableResponse, error) {
	return repository.Domain[model.Table, server.ReadTableRequest, server.ReadTableResponse](x, req, call)
}

func (x *ReadTable) Input(req *server.ReadTableRequest) *model.Table {
	return &model.Table{
		ID: req.GetID(),
	}
}

func (x *ReadTable) Output(res *model.Table) *server.ReadTableResponse {
	cells := make([]*server.Cell, 0)
	for _, v := range res.Cells {
		cells = append(cells, cell.ToInfra(v))
	}
	return &server.ReadTableResponse{
		ID: res.ID,
		Table: &server.Table{
			Title: res.Title,
			Cells: cells,
		},
	}
}

type UpdateTable struct {
}

func (x *UpdateTable) Domain(req *server.UpdateTableRequest, call func(table *model.Table) (*model.Table, error)) (*server.UpdateTableResponse, error) {
	return repository.Domain[model.Table, server.UpdateTableRequest, server.UpdateTableResponse](x, req, call)
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

func (x *DeleteTable) Domain(req *server.DeleteTableRequest, call func(table *model.Table) (*model.Table, error)) (*server.DeleteTableResponse, error) {
	return repository.Domain[model.Table, server.DeleteTableRequest, server.DeleteTableResponse](x, req, call)
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
