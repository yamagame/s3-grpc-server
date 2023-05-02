package cell

import (
	"sample/s3-grpc-server/grpc_server/service/repository"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
	"time"
)

type CreateCell struct {
}

func (x *CreateCell) Domain(req *server.CreateCellRequest, call func(table *model.Cell) (*model.Cell, error)) (*server.CreateCellResponse, error) {
	return repository.Domain[model.Cell, server.CreateCellRequest, server.CreateCellResponse](x, req, call)
}

func (x *CreateCell) Input(req *server.CreateCellRequest) *model.Cell {
	return &model.Cell{
		TableID:   req.Cell.TableId,
		Row:       req.Cell.Row,
		Col:       req.Cell.Col,
		Text:      req.Cell.Text,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (x *CreateCell) Output(res *model.Cell) *server.CreateCellResponse {
	return &server.CreateCellResponse{
		ID: res.ID,
	}
}

type ReadCell struct {
}

func (x *ReadCell) Domain(req *server.ReadCellRequest, call func(table *model.Cell) (*model.Cell, error)) (*server.ReadCellResponse, error) {
	return repository.Domain[model.Cell, server.ReadCellRequest, server.ReadCellResponse](x, req, call)
}

func (x *ReadCell) Input(req *server.ReadCellRequest) *model.Cell {
	return &model.Cell{
		ID: req.GetID(),
	}
}

func (x *ReadCell) Output(res *model.Cell) *server.ReadCellResponse {
	return &server.ReadCellResponse{
		Cell: &server.Cell{
			Text: res.Text,
		},
	}
}

type UpdateCell struct {
}

func (x *UpdateCell) Domain(req *server.UpdateCellRequest, call func(table *model.Cell) (*model.Cell, error)) (*server.UpdateCellResponse, error) {
	return repository.Domain[model.Cell, server.UpdateCellRequest, server.UpdateCellResponse](x, req, call)
}

func (x *UpdateCell) Input(req *server.UpdateCellRequest) *model.Cell {
	return &model.Cell{
		ID:        req.Cell.Id,
		Text:      req.Cell.Text,
		UpdatedAt: time.Now(),
	}
}

func (x *UpdateCell) Output(res *model.Cell) *server.UpdateCellResponse {
	return &server.UpdateCellResponse{
		ID: res.ID,
	}
}

type DeleteCell struct {
}

func (x *DeleteCell) Domain(req *server.DeleteCellRequest, call func(table *model.Cell) (*model.Cell, error)) (*server.DeleteCellResponse, error) {
	return repository.Domain[model.Cell, server.DeleteCellRequest, server.DeleteCellResponse](x, req, call)
}

func (x *DeleteCell) Input(req *server.DeleteCellRequest) *model.Cell {
	return &model.Cell{
		ID: req.GetID(),
	}
}

func (x *DeleteCell) Output(res *model.Cell) *server.DeleteCellResponse {
	return &server.DeleteCellResponse{
		ID: res.ID,
	}
}
