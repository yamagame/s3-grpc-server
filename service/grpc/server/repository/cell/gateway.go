package cell

import (
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/libs/gateway"
	"sample/s3-grpc-server/proto/grpc_server"

	"time"
)

type CreateCell struct {
}

func (x *CreateCell) ToDomain(req *grpc_server.CreateCellRequest, call func(table *model.Cell) (*model.Cell, error)) (*grpc_server.CreateCellResponse, error) {
	return gateway.ToDomain[model.Cell, grpc_server.CreateCellRequest, grpc_server.CreateCellResponse](x, req, call)
}

func (x *CreateCell) Input(req *grpc_server.CreateCellRequest) *model.Cell {
	return &model.Cell{
		TableID:   req.Cell.TableId,
		Row:       req.Cell.Row,
		Col:       req.Cell.Col,
		Text:      req.Cell.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (x *CreateCell) Output(res *model.Cell) *grpc_server.CreateCellResponse {
	return &grpc_server.CreateCellResponse{
		ID: res.ID,
	}
}

type ReadCell struct {
}

func (x *ReadCell) ToDomain(req *grpc_server.ReadCellRequest, call func(table *model.Cell) (*model.Cell, error)) (*grpc_server.ReadCellResponse, error) {
	return gateway.ToDomain[model.Cell, grpc_server.ReadCellRequest, grpc_server.ReadCellResponse](x, req, call)
}

func (x *ReadCell) Input(req *grpc_server.ReadCellRequest) *model.Cell {
	return &model.Cell{
		ID: req.GetID(),
	}
}

func (x *ReadCell) Output(res *model.Cell) *grpc_server.ReadCellResponse {
	return &grpc_server.ReadCellResponse{
		Cell: &grpc_server.Cell{
			Text: res.Text,
		},
	}
}

type UpdateCell struct {
}

func (x *UpdateCell) ToDomain(req *grpc_server.UpdateCellRequest, call func(table *model.Cell) (*model.Cell, error)) (*grpc_server.UpdateCellResponse, error) {
	return gateway.ToDomain[model.Cell, grpc_server.UpdateCellRequest, grpc_server.UpdateCellResponse](x, req, call)
}

func (x *UpdateCell) Input(req *grpc_server.UpdateCellRequest) *model.Cell {
	return &model.Cell{
		ID:        req.Cell.Id,
		Text:      req.Cell.Text,
		UpdatedAt: time.Now(),
	}
}

func (x *UpdateCell) Output(res *model.Cell) *grpc_server.UpdateCellResponse {
	return &grpc_server.UpdateCellResponse{
		ID: res.ID,
	}
}

type DeleteCell struct {
}

func (x *DeleteCell) ToDomain(req *grpc_server.DeleteCellRequest, call func(table *model.Cell) (*model.Cell, error)) (*grpc_server.DeleteCellResponse, error) {
	return gateway.ToDomain[model.Cell, grpc_server.DeleteCellRequest, grpc_server.DeleteCellResponse](x, req, call)
}

func (x *DeleteCell) Input(req *grpc_server.DeleteCellRequest) *model.Cell {
	return &model.Cell{
		ID: req.GetID(),
	}
}

func (x *DeleteCell) Output(res *model.Cell) *grpc_server.DeleteCellResponse {
	return &grpc_server.DeleteCellResponse{
		ID: res.ID,
	}
}
