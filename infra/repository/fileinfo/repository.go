package fileinfo

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *repository.CRUDRepository[model.FileInfo] {
	return &repository.CRUDRepository[model.FileInfo]{
		DB: db,
	}
}
