package repository

import "sample/s3-grpc-server/infra/repository/model"

type RepositoryInterface interface {
	CreateFileInfo() error
	ReadFileInfo() (*model.FileInfo, error)
	UpdateFileInfo(file *model.FileInfo) error
	DeleteFileInfo(file *model.FileInfo) error
}
