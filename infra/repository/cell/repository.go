package cell

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"

	"gorm.io/gorm"
)

type CellRepository struct {
	repository.CRUDRepository[model.Cell]
}

func NewCellRepository(db *gorm.DB) *CellRepository {
	return &CellRepository{
		CRUDRepository: repository.CRUDRepository[model.Cell]{DB: db},
	}
}
