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
	return s.CRUDRepository.Read(object)
}
