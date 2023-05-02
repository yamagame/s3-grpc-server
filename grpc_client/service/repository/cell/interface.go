package cell

import "sample/s3-grpc-server/infra/repository/model"

type CellScannerInterface interface {
	CreateCell() *model.Cell
	ReadCell() *model.Cell
	UpdateCell() *model.Cell
	DeleteCell() *model.Cell
}
