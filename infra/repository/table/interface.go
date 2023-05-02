package table

import "sample/s3-grpc-server/infra/repository/model"

type RepositoryInterface interface {
	Create(file *model.Table) (*model.Table, error)
	Read(file *model.Table) (*model.Table, error)
	Update(file *model.Table) (*model.Table, error)
	Delete(file *model.Table) (*model.Table, error)
}
