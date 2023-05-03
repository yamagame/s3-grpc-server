package table

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/client/repository"
)

type CreateTable struct {
}

func (x *CreateTable) Domain(req *model.Table, call func(req *server.CreateTableRequest) (*server.CreateTableResponse, error)) (*model.Table, error) {
	return repository.Domain[model.Table, server.CreateTableRequest, server.CreateTableResponse](x, req, call)
}

func (x *CreateTable) Input(req *model.Table) *server.CreateTableRequest {
	return &server.CreateTableRequest{
		Table: &server.Table{
			FileId: req.FileID,
			Title:  req.Title,
		},
	}
}

func (x *CreateTable) Output(res *server.CreateTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}

type ReadTable struct{}

func (x *ReadTable) Domain(req *model.Table, call func(table *server.ReadTableRequest) (*server.ReadTableResponse, error)) (*model.Table, error) {
	return repository.Domain[model.Table, server.ReadTableRequest, server.ReadTableResponse](x, req, call)
}

func (x *ReadTable) Input(req *model.Table) *server.ReadTableRequest {
	return &server.ReadTableRequest{
		ID: req.ID,
	}
}

func (x *ReadTable) Output(res *server.ReadTableResponse) *model.Table {
	cells := make([]*model.Cell, 0)
	for _, cell := range res.Table.Cells {
		cells = append(cells, &model.Cell{
			Text: cell.Text,
			Row:  cell.Row,
			Col:  cell.Col,
		})
	}
	return &model.Table{
		ID:    res.GetID(),
		Title: res.Table.Title,
		Cells: cells,
	}
}

type UpdateTable struct{}

func (x *UpdateTable) Domain(req *model.Table, call func(table *server.UpdateTableRequest) (*server.UpdateTableResponse, error)) (*model.Table, error) {
	return repository.Domain[model.Table, server.UpdateTableRequest, server.UpdateTableResponse](x, req, call)
}

func (x *UpdateTable) Input(req *model.Table) *server.UpdateTableRequest {
	return &server.UpdateTableRequest{
		ID: req.ID,
		Table: &server.Table{
			Title: req.Title,
		},
	}
}

func (x *UpdateTable) Output(res *server.UpdateTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}

type DeleteTable struct{}

func (x *DeleteTable) Domain(req *model.Table, call func(table *server.DeleteTableRequest) (*server.DeleteTableResponse, error)) (*model.Table, error) {
	return repository.Domain[model.Table, server.DeleteTableRequest, server.DeleteTableResponse](x, req, call)
}

func (x *DeleteTable) Input(req *model.Table) *server.DeleteTableRequest {
	return &server.DeleteTableRequest{
		ID: req.ID,
	}
}

func (x *DeleteTable) Output(res *server.DeleteTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}
