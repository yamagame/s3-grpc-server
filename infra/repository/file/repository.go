package file

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"

	"gorm.io/gorm"
)

type fileRepository struct {
	repository.CRUDRepository[model.File]
}

func NewRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{
		CRUDRepository: repository.CRUDRepository[model.File]{DB: db},
	}
}
