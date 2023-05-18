package cell

import (
	"sample/s3-grpc-server/entitiy/repository/gateway/cell"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"

	"time"
)

type CreateCell struct {
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

type ListCell struct {
}

func (x *ListCell) Input(req *grpc_server.ListCellRequest) *model.Cell {
	return &model.Cell{}
}

func (x *ListCell) Output(res []*model.Cell) *grpc_server.ListCellResponse {
	cells := make([]*grpc_server.Cell, 0)
	for _, v := range res {
		cells = append(cells, cell.ToInfra(v))
	}
	return &grpc_server.ListCellResponse{
		Cells: cells,
	}
}
