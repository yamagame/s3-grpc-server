package cell

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateCell struct {
}

func (x *CreateCell) Input(req *server.CreateCellRequest) *model.Cell {
	return &model.Cell{
		TableID: req.Cell.TableId,
		Row:     req.Cell.Row,
		Col:     req.Cell.Col,
		Text:    req.Cell.Text,
	}
}

func (x *CreateCell) Output(res *model.Cell) *server.CreateCellResponse {
	return &server.CreateCellResponse{
		ID: int64(res.ID),
	}
}

type ReadCell struct {
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

func (x *UpdateCell) Input(req *server.UpdateCellRequest) *model.Cell {
	return &model.Cell{
		ID:   req.Cell.Id,
		Text: req.Cell.Text,
	}
}

func (x *UpdateCell) Output(res *model.Cell) *server.UpdateCellResponse {
	return &server.UpdateCellResponse{
		ID: res.ID,
	}
}

type DeleteCell struct {
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
