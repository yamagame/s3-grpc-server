package repository

import (
	"gorm.io/gorm"
)

type CRUDRepository[T any] struct {
	DB *gorm.DB
}

// Create implements CRUDRepository.Create
func (s *CRUDRepository[T]) Create(object *T) (*T, error) {
	s.DB.Create(object)
	return object, nil
}

// Read implements CRUDRepository.Read
func (s *CRUDRepository[T]) Read(object *T) (*T, error) {
	s.DB.Take(object)
	return object, nil
}

// Update implements CRUDRepository.Update
func (s *CRUDRepository[T]) Update(object *T) (*T, error) {
	s.DB.Updates(object)
	return object, nil
}

// Delete implements CRUDRepository.Delete
func (s *CRUDRepository[T]) Delete(object *T) (*T, error) {
	s.DB.Delete(object)
	return object, nil
}
