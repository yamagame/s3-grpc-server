package file

import (
	"context"
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

func (s *fileRepository) Read(ctx context.Context, object *model.File) (*model.File, error) {
	if err := s.DB.Model(&model.File{}).Preload("Tables").Take(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}
