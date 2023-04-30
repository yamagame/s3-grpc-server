package repository

import (
	db "sample/s3-grpc-server/infra/repository"

	"gorm.io/gorm"
)

type repositoryService struct {
	db *gorm.DB
}

func NewRepositoryService(db *gorm.DB) *repositoryService {
	return &repositoryService{
		db: db,
	}
}

// CreateFileInfo implements repositoryService.CreateFileInfo
func (s *repositoryService) CreateFileInfo(file *db.FileInfo) (*db.FileInfo, error) {
	s.db.Create(file)
	return file, nil
}

// ReadFileInfo implements repositoryService.ReadFileInfo
func (s *repositoryService) ReadFileInfo(file *db.FileInfo) (*db.FileInfo, error) {
	s.db.Take(file)
	return file, nil
}

// UpdateFileInfo implements repositoryService.UpdateFileInfo
func (s *repositoryService) UpdateFileInfo(file *db.FileInfo) (*db.FileInfo, error) {
	s.db.Updates(file)
	return file, nil
}

// DeleteFileInfo implements repositoryService.DeleteFileInfo
func (s *repositoryService) DeleteFileInfo(file *db.FileInfo) (*db.FileInfo, error) {
	s.db.Delete(file)
	return file, nil
}
