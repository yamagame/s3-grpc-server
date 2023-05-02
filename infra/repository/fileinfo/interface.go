package fileinfo

import "sample/s3-grpc-server/infra/repository/model"

type RepositoryInterface interface {
	Create(file *model.FileInfo) (*model.FileInfo, error)
	Read(file *model.FileInfo) (*model.FileInfo, error)
	Update(file *model.FileInfo) (*model.FileInfo, error)
	Delete(file *model.FileInfo) (*model.FileInfo, error)
}
