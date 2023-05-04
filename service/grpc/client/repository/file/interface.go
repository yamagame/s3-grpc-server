package file

import "sample/s3-grpc-server/entitiy/repository/model"

type ScannerInterface interface {
	Create() *model.File
	Read() *model.File
	Update() *model.File
	Delete() *model.File
}
