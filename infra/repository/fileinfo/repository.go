package fileinfo

import (
	"sample/s3-grpc-server/infra/repository/model"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CreateFileInfo implements RepositoryService.CreateFileInfo
func (s *Repository) CreateFileInfo(file *model.FileInfo) (*model.FileInfo, error) {
	s.db.Create(file)
	return file, nil
}

// ReadFileInfo implements RepositoryService.ReadFileInfo
func (s *Repository) ReadFileInfo(file *model.FileInfo) (*model.FileInfo, error) {
	s.db.Take(file)
	return file, nil
}

// UpdateFileInfo implements RepositoryService.UpdateFileInfo
func (s *Repository) UpdateFileInfo(file *model.FileInfo) (*model.FileInfo, error) {
	s.db.Updates(file)
	return file, nil
}

// DeleteFileInfo implements RepositoryService.DeleteFileInfo
func (s *Repository) DeleteFileInfo(file *model.FileInfo) (*model.FileInfo, error) {
	s.db.Delete(file)
	return file, nil
}
