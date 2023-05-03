package table

import "sample/s3-grpc-server/infra/repository/model"

type ScannerInterface interface {
	Create() *model.Table
	Read() *model.Table
	Update() *model.Table
	Delete() *model.Table
}
