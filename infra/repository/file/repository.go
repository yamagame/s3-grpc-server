package file

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *repository.CRUDRepository[model.File] {
	return &repository.CRUDRepository[model.File]{
		DB: db,
	}
}
