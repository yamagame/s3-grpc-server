package cell

import "sample/s3-grpc-server/infra/repository/model"

type RepositoryInterface interface {
	Create(file *model.Cell) (*model.Cell, error)
	Read(file *model.Cell) (*model.Cell, error)
	Update(file *model.Cell) (*model.Cell, error)
	Delete(file *model.Cell) (*model.Cell, error)
}
