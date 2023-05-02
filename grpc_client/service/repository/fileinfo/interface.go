package fileinfo

import "sample/s3-grpc-server/infra/repository/model"

type FileInfoScannerInterface interface {
	CreateFileInfo() *model.FileInfo
	ReadFileInfo() *model.FileInfo
	UpdateFileInfo() *model.FileInfo
	DeleteFileInfo() *model.FileInfo
}
