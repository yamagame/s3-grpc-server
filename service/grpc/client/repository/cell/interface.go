package cell

import (
	"sample/s3-grpc-server/entitiy/repository/model"
)

type ScannerInterface interface {
	Create() *model.Cell
	Read() *model.Cell
	Update() *model.Cell
	Delete() *model.Cell
}
