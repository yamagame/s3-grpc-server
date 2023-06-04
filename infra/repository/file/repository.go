package file

import (
	"context"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type FileRepository struct {
	repository.CRUDRepository[model.File]
}

func NewFileRepository(db *gorm.DB) RepositoryInterface {
	return &FileRepository{
		CRUDRepository: repository.CRUDRepository[model.File]{DB: db},
	}
}

// Create implements FileRepository.Create
func (s *FileRepository) Create(ctx context.Context, object *model.File) (*model.File, error) {
	if err := s.DB.
		// Clauses(clause.OnConflict{
		// 	Columns:   []clause.Column{{Name: "id"}},
		// 	DoUpdates: clause.Assignments(map[string]interface{}{"filename": object.Filename}),
		// }).
		Create(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Read implements FileRepository.Read
func (s *FileRepository) Read(ctx context.Context, object *model.File) (*model.File, error) {
	if err := s.DB.Model(&model.File{}).Preload(clause.Associations).Take(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Update implements FileRepository.Update
func (s *FileRepository) Update(ctx context.Context, object *model.File) (*model.File, error) {
	if err := s.DB.Model(object).Session(&gorm.Session{FullSaveAssociations: true}).Updates(object).Error; err != nil {
		return nil, err
	}
	// if err := s.DB.Model(object).Omit(clause.Associations).Updates(object).Error; err != nil {
	// 	return nil, err
	// }
	// if err := s.DB.Model(&object.Attr).Updates(&object.Attr).Error; err != nil {
	// 	return nil, err
	// }
	return object, nil
}
