package cell

import (
	"sample/s3-grpc-server/grpc_client/service/repository"
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type CreateCell struct{}

func (x *CreateCell) Domain(req *model.Cell, call func(cell *server.CreateCellRequest) (*server.CreateCellResponse, error)) (*model.Cell, error) {
	return repository.Domain[model.Cell, server.CreateCellRequest, server.CreateCellResponse](x, req, call)
}

func (x *CreateCell) Input(req *model.Cell) *server.CreateCellRequest {
	return &server.CreateCellRequest{
		Cell: cell.ToInfra(req),
	}
}

func (x *CreateCell) Output(res *server.CreateCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}

type ReadCell struct{}

func (x *ReadCell) Domain(req *model.Cell, call func(cell *server.ReadCellRequest) (*server.ReadCellResponse, error)) (*model.Cell, error) {
	return repository.Domain[model.Cell, server.ReadCellRequest, server.ReadCellResponse](x, req, call)
}

func (x *ReadCell) Input(req *model.Cell) *server.ReadCellRequest {
	return &server.ReadCellRequest{
		ID: req.ID,
	}
}

func (x *ReadCell) Output(res *server.ReadCellResponse) *model.Cell {
	return cell.ToDomain(res.Cell)
}

type UpdateCell struct{}

func (x *UpdateCell) Domain(req *model.Cell, call func(cell *server.UpdateCellRequest) (*server.UpdateCellResponse, error)) (*model.Cell, error) {
	return repository.Domain[model.Cell, server.UpdateCellRequest, server.UpdateCellResponse](x, req, call)
}

func (x *UpdateCell) Input(req *model.Cell) *server.UpdateCellRequest {
	return &server.UpdateCellRequest{
		Cell: cell.ToInfra(req),
	}
}

func (x *UpdateCell) Output(res *server.UpdateCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}

type DeleteCell struct{}

func (x *DeleteCell) Domain(req *model.Cell, call func(cell *server.DeleteCellRequest) (*server.DeleteCellResponse, error)) (*model.Cell, error) {
	return repository.Domain[model.Cell, server.DeleteCellRequest, server.DeleteCellResponse](x, req, call)
}

func (x *DeleteCell) Input(req *model.Cell) *server.DeleteCellRequest {
	return &server.DeleteCellRequest{
		ID: req.ID,
	}
}

func (x *DeleteCell) Output(res *server.DeleteCellResponse) *model.Cell {
	return &model.Cell{
		ID: res.GetID(),
	}
}
