package table

import "sample/s3-grpc-server/infra/repository/model"

type TableScannerInterface interface {
	CreateTable() *model.Table
	ReadTable() *model.Table
	UpdateTable() *model.Table
	DeleteTable() *model.Table
}
