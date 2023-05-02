package cell

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateCellDTO struct {
}

func (x *CreateCellDTO) Input(req *model.Cell) *server.CreateCellRequest {
	return &server.CreateCellRequest{
		Cell: &server.Cell{
			Text: req.Text,
		},
	}
}

func (x *CreateCellDTO) Output(res *server.CreateCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}

type ReadCellDTO struct {
}

func (x *ReadCellDTO) Input(req *model.Cell) *server.ReadCellRequest {
	return &server.ReadCellRequest{
		ID: req.ID,
	}
}

func (x *ReadCellDTO) Output(res *server.ReadCellResponse) *model.Cell {
	return &model.Cell{
		ID:   res.GetID(),
		Text: res.Cell.Text,
	}
}

type UpdateCellDTO struct {
}

func (x *UpdateCellDTO) Input(req *model.Cell) *server.UpdateCellRequest {
	return &server.UpdateCellRequest{
		ID: req.ID,
		Cell: &server.Cell{
			Text: req.Text,
		},
	}
}

func (x *UpdateCellDTO) Output(res *server.UpdateCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}

type DeleteCellDTO struct {
}

func (x *DeleteCellDTO) Input(req *model.Cell) *server.DeleteCellRequest {
	return &server.DeleteCellRequest{
		ID: req.ID,
	}
}

func (x *DeleteCellDTO) Output(res *server.DeleteCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}
