package cell

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

func ToDomain(cell *server.Cell) *model.Cell {
	return &model.Cell{
		ID:   cell.Id,
		Row:  cell.Row,
		Col:  cell.Col,
		Text: cell.Text,
	}
}

func ToInfra(cell *model.Cell) *server.Cell {
	return &server.Cell{
		Id:   cell.ID,
		Row:  cell.Row,
		Col:  cell.Col,
		Text: cell.Text,
	}
}
