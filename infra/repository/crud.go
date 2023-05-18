package repository

import (
	"context"

	"gorm.io/gorm"
)

type CRUDRepositoryInterface[T any] interface {
	Create(ctx context.Context, file *T) (*T, error)
	Read(ctx context.Context, file *T) (*T, error)
	Update(ctx context.Context, file *T) (*T, error)
	Delete(ctx context.Context, file *T) (*T, error)
	List(ctx context.Context, file *T) ([]*T, error)
}

type CRUDRepository[T any] struct {
	DB *gorm.DB
}

// Create implements CRUDRepository.Create
func (s *CRUDRepository[T]) Create(ctx context.Context, object *T) (*T, error) {
	if err := s.DB.CreateInBatches([]*T{object}, 1).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Read implements CRUDRepository.Read
func (s *CRUDRepository[T]) Read(ctx context.Context, object *T) (*T, error) {
	if err := s.DB.Take(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Update implements CRUDRepository.Update
func (s *CRUDRepository[T]) Update(ctx context.Context, object *T) (*T, error) {
	if err := s.DB.Updates(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// Delete implements CRUDRepository.Delete
func (s *CRUDRepository[T]) Delete(ctx context.Context, object *T) (*T, error) {
	if err := s.DB.Delete(object).Error; err != nil {
		return nil, err
	}
	return object, nil
}

// List implements CRUDRepository.List
func (s *CRUDRepository[T]) List(ctx context.Context, object *T) ([]*T, error) {
	t := []*T{}
	if err := s.DB.Find(&t).Error; err != nil {
		return nil, err
	}
	return t, nil
}
