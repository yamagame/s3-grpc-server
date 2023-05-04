package cell

import (
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/libs/dto"
	"sample/s3-grpc-server/proto/grpc_server"
)

type CreateCell struct{}

func (x *CreateCell) ToInfra(req *model.Cell, call func(cell *grpc_server.CreateCellRequest) (*grpc_server.CreateCellResponse, error)) (*model.Cell, error) {
	return dto.ToInfra[model.Cell, grpc_server.CreateCellRequest, grpc_server.CreateCellResponse](x, req, call)
}

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

func (x *ReadCell) ToInfra(req *model.Cell, call func(cell *grpc_server.ReadCellRequest) (*grpc_server.ReadCellResponse, error)) (*model.Cell, error) {
	return dto.ToInfra[model.Cell, grpc_server.ReadCellRequest, grpc_server.ReadCellResponse](x, req, call)
}

func (x *ReadCell) Input(req *model.Cell) *grpc_server.ReadCellRequest {
	return &grpc_server.ReadCellRequest{
		ID: req.ID,
	}
}

func (x *ReadCell) Output(res *grpc_server.ReadCellResponse) *model.Cell {
	return cell.ToDomain(res.Cell)
}

type UpdateCell struct{}

func (x *UpdateCell) ToInfra(req *model.Cell, call func(cell *grpc_server.UpdateCellRequest) (*grpc_server.UpdateCellResponse, error)) (*model.Cell, error) {
	return dto.ToInfra[model.Cell, grpc_server.UpdateCellRequest, grpc_server.UpdateCellResponse](x, req, call)
}

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

func (x *DeleteCell) ToInfra(req *model.Cell, call func(cell *grpc_server.DeleteCellRequest) (*grpc_server.DeleteCellResponse, error)) (*model.Cell, error) {
	return dto.ToInfra[model.Cell, grpc_server.DeleteCellRequest, grpc_server.DeleteCellResponse](x, req, call)
}

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
