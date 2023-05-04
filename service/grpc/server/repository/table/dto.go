package table

import (
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/libs/dto"
	"sample/s3-grpc-server/proto/grpc_server"

	"time"
)

type CreateTable struct {
}

func (x *CreateTable) Domain(req *grpc_server.CreateTableRequest, call func(table *model.Table) (*model.Table, error)) (*grpc_server.CreateTableResponse, error) {
	return dto.ToDomain[model.Table, grpc_server.CreateTableRequest, grpc_server.CreateTableResponse](x, req, call)
}

func (x *CreateTable) Input(req *grpc_server.CreateTableRequest) *model.Table {
	return &model.Table{
		FileID:    req.Table.FileId,
		Title:     req.Table.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (x *CreateTable) Output(res *model.Table) *grpc_server.CreateTableResponse {
	return &grpc_server.CreateTableResponse{
		ID: res.ID,
	}
}

type ReadTable struct {
}

func (x *ReadTable) Domain(req *grpc_server.ReadTableRequest, call func(table *model.Table) (*model.Table, error)) (*grpc_server.ReadTableResponse, error) {
	return dto.ToDomain[model.Table, grpc_server.ReadTableRequest, grpc_server.ReadTableResponse](x, req, call)
}

func (x *ReadTable) Input(req *grpc_server.ReadTableRequest) *model.Table {
	return &model.Table{
		ID: req.GetID(),
	}
}

func (x *ReadTable) Output(res *model.Table) *grpc_server.ReadTableResponse {
	cells := make([]*grpc_server.Cell, 0)
	for _, v := range res.Cells {
		cells = append(cells, cell.ToInfra(v))
	}
	return &grpc_server.ReadTableResponse{
		ID: res.ID,
		Table: &grpc_server.Table{
			Title: res.Title,
			Cells: cells,
		},
	}
}

type UpdateTable struct {
}

func (x *UpdateTable) Domain(req *grpc_server.UpdateTableRequest, call func(table *model.Table) (*model.Table, error)) (*grpc_server.UpdateTableResponse, error) {
	return dto.ToDomain[model.Table, grpc_server.UpdateTableRequest, grpc_server.UpdateTableResponse](x, req, call)
}

func (x *UpdateTable) Input(req *grpc_server.UpdateTableRequest) *model.Table {
	return &model.Table{
		ID:        req.GetID(),
		Title:     req.Table.Title,
		UpdatedAt: time.Now(),
	}
}

func (x *UpdateTable) Output(res *model.Table) *grpc_server.UpdateTableResponse {
	return &grpc_server.UpdateTableResponse{
		ID: res.ID,
	}
}

type DeleteTable struct {
}

func (x *DeleteTable) Domain(req *grpc_server.DeleteTableRequest, call func(table *model.Table) (*model.Table, error)) (*grpc_server.DeleteTableResponse, error) {
	return dto.ToDomain[model.Table, grpc_server.DeleteTableRequest, grpc_server.DeleteTableResponse](x, req, call)
}

func (x *DeleteTable) Input(req *grpc_server.DeleteTableRequest) *model.Table {
	return &model.Table{
		ID: req.GetID(),
	}
}

func (x *DeleteTable) Output(res *model.Table) *grpc_server.DeleteTableResponse {
	return &grpc_server.DeleteTableResponse{
		ID: res.ID,
	}
}
