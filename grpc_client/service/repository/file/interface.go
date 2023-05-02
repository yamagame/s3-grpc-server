package file

import "sample/s3-grpc-server/infra/repository/model"

type FileScannerInterface interface {
	CreateFile() *model.File
	ReadFile() *model.File
	UpdateFile() *model.File
	DeleteFile() *model.File
}
