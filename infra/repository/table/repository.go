package table

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *repository.CRUDRepository[model.Table] {
	return &repository.CRUDRepository[model.Table]{
		DB: db,
	}
}
