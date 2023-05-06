package table

import (
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateTable struct {
}

func (x *CreateTable) Input(req *model.Table) *grpc_server.CreateTableRequest {
	return &grpc_server.CreateTableRequest{
		Table: &grpc_server.Table{
			FileId: req.FileID,
			Title:  req.Title,
		},
	}
}

func (x *CreateTable) Output(res *grpc_server.CreateTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}

type ReadTable struct{}

func (x *ReadTable) Input(req *model.Table) *grpc_server.ReadTableRequest {
	return &grpc_server.ReadTableRequest{
		ID: req.ID,
	}
}

func (x *ReadTable) Output(res *grpc_server.ReadTableResponse) *model.Table {
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

func (x *UpdateTable) Input(req *model.Table) *grpc_server.UpdateTableRequest {
	return &grpc_server.UpdateTableRequest{
		ID: req.ID,
		Table: &grpc_server.Table{
			Title: req.Title,
		},
	}
}

func (x *UpdateTable) Output(res *grpc_server.UpdateTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}

type DeleteTable struct{}

func (x *DeleteTable) Input(req *model.Table) *grpc_server.DeleteTableRequest {
	return &grpc_server.DeleteTableRequest{
		ID: req.ID,
	}
}

func (x *DeleteTable) Output(res *grpc_server.DeleteTableResponse) *model.Table {
	return &model.Table{
		ID: res.GetID(),
	}
}
