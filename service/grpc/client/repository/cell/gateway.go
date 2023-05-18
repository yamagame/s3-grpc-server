package cell

import (
	"sample/s3-grpc-server/entitiy/repository/gateway/cell"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateCell struct{}

func (x *CreateCell) Input(req *model.Cell) *grpc_server.CreateCellRequest {
	return &grpc_server.CreateCellRequest{
		Cell: cell.ToInfra(req),
	}
}

func (x *CreateCell) Output(res *grpc_server.CreateCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}

type ReadCell struct{}

func (x *ReadCell) Input(req *model.Cell) *grpc_server.ReadCellRequest {
	return &grpc_server.ReadCellRequest{
		ID: req.ID,
	}
}

func (x *ReadCell) Output(res *grpc_server.ReadCellResponse) *model.Cell {
	return cell.ToDomain(res.Cell)
}

type UpdateCell struct{}

func (x *UpdateCell) Input(req *model.Cell) *grpc_server.UpdateCellRequest {
	return &grpc_server.UpdateCellRequest{
		Cell: cell.ToInfra(req),
	}
}

func (x *UpdateCell) Output(res *grpc_server.UpdateCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}

type DeleteCell struct{}

func (x *DeleteCell) Input(req *model.Cell) *grpc_server.DeleteCellRequest {
	return &grpc_server.DeleteCellRequest{
		ID: req.ID,
	}
}

func (x *DeleteCell) Output(res *grpc_server.DeleteCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}

type ListCell struct{}

func (x *ListCell) Input(req *model.Cell) *grpc_server.ListCellRequest {
	return &grpc_server.ListCellRequest{
		//
	}
}

func (x *ListCell) Output(res *grpc_server.ListCellResponse) []*model.Cell {
	ret := []*model.Cell{}
	for _, v := range res.Cells {
		ret = append(ret, cell.ToDomain(v))
	}
	return ret
}
