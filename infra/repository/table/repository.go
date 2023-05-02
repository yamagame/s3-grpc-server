package table

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"

	"gorm.io/gorm"
)

type tableRepository struct {
	repository.CRUDRepository[model.Table]
}

func NewRepository(db *gorm.DB) *tableRepository {
	return &tableRepository{
		CRUDRepository: repository.CRUDRepository[model.Table]{DB: db},
	}
}

func (s *tableRepository) Read(object *model.Table) (*model.Table, error) {
	if err := s.DB.Model(&model.Table{}).Preload("Cells").Take(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}
