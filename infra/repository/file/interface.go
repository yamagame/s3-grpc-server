package file

import "sample/s3-grpc-server/infra/repository/model"

type RepositoryInterface interface {
	Create(file *model.File) (*model.File, error)
	Read(file *model.File) (*model.File, error)
	Update(file *model.File) (*model.File, error)
	Delete(file *model.File) (*model.File, error)
}
