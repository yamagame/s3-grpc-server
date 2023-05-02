package cell

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"

	"gorm.io/gorm"
)

type cellRepository struct {
	repository.CRUDRepository[model.Cell]
}

func NewRepository(db *gorm.DB) *cellRepository {
	return &cellRepository{
		CRUDRepository: repository.CRUDRepository[model.Cell]{DB: db},
	}
}
