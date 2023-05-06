package file

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository"

	"gorm.io/gorm"
)

type FileRepository struct {
	repository.CRUDRepository[model.File]
}

func NewFileRepository(db *gorm.DB) RepositoryInterface {
	return &FileRepository{
		CRUDRepository: repository.CRUDRepository[model.File]{DB: db},
	}
}

func (s *FileRepository) Read(ctx context.Context, object *model.File) (*model.File, error) {
	if err := s.DB.Model(&model.File{}).Preload("Tables").Take(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}
