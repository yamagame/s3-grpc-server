package repository

import (
	"gorm.io/gorm"
)

type RepositoryService struct {
	db *gorm.DB
}

func NewRepositoryService(db *gorm.DB) *RepositoryService {
	return &RepositoryService{
		db: db,
	}
}

// CreateFileInfo implements RepositoryService.CreateFileInfo
func (s *RepositoryService) CreateFileInfo(file *FileInfo) (*FileInfo, error) {
	s.db.Create(file)
	return file, nil
}

// ReadFileInfo implements RepositoryService.ReadFileInfo
func (s *RepositoryService) ReadFileInfo(file *FileInfo) (*FileInfo, error) {
	s.db.Take(file)
	return file, nil
}

// UpdateFileInfo implements RepositoryService.UpdateFileInfo
func (s *RepositoryService) UpdateFileInfo(file *FileInfo) (*FileInfo, error) {
	s.db.Updates(file)
	return file, nil
}

// DeleteFileInfo implements RepositoryService.DeleteFileInfo
func (s *RepositoryService) DeleteFileInfo(file *FileInfo) (*FileInfo, error) {
	s.db.Delete(file)
	return file, nil
}