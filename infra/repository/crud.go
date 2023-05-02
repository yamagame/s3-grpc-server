package repository

import (
	"gorm.io/gorm"
)

type CRUDRepository[T any] struct {
	DB *gorm.DB
}

// Create implements CRUDRepository.Create
func (s *CRUDRepository[T]) Create(object *T) (*T, error) {
	if err := s.DB.Create(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Read implements CRUDRepository.Read
func (s *CRUDRepository[T]) Read(object *T) (*T, error) {
	if err := s.DB.Take(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Update implements CRUDRepository.Update
func (s *CRUDRepository[T]) Update(object *T) (*T, error) {
	if err := s.DB.Updates(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Delete implements CRUDRepository.Delete
func (s *CRUDRepository[T]) Delete(object *T) (*T, error) {
	if err := s.DB.Delete(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}
