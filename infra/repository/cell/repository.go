package cell

import (
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository"

	"gorm.io/gorm"
)

type CellRepository struct {
	repository.CRUDRepository[model.Cell]
}

func NewCellRepository(db *gorm.DB) RepositoryInterface {
	return &CellRepository{
		CRUDRepository: repository.CRUDRepository[model.Cell]{DB: db},
	}
}
