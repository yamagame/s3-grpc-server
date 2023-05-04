package cell

import (
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

func ToDomain(cell *grpc_server.Cell) *model.Cell {
	return &model.Cell{
		ID:   cell.Id,
		Row:  cell.Row,
		Col:  cell.Col,
		Text: cell.Text,
	}
}

func ToInfra(cell *model.Cell) *grpc_server.Cell {
	return &grpc_server.Cell{
		Id:   cell.ID,
		Row:  cell.Row,
		Col:  cell.Col,
		Text: cell.Text,
	}
}
