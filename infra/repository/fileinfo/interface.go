package fileinfo

import "sample/s3-grpc-server/infra/repository/fileinfo/model"

type RepositoryInterface interface {
	CreateFileInfo(file *model.FileInfo) (*model.FileInfo, error)
	ReadFileInfo(file *model.FileInfo) (*model.FileInfo, error)
	UpdateFileInfo(file *model.FileInfo) (*model.FileInfo, error)
	DeleteFileInfo(file *model.FileInfo) (*model.FileInfo, error)
}
