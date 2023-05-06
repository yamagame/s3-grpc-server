package table

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository"

	"gorm.io/gorm"
)

type TableRepository struct {
	repository.CRUDRepository[model.Table]
}

func NewTableRepository(db *gorm.DB) RepositoryInterface {
	return &TableRepository{
		CRUDRepository: repository.CRUDRepository[model.Table]{DB: db},
	}
}

func (s *TableRepository) Read(ctx context.Context, object *model.Table) (*model.Table, error) {
	if err := s.DB.Model(&model.Table{}).Preload("Cells").Take(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}
