package model

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        uint64
	Filename  string
	Tables    []Table
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
