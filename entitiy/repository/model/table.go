package model

import (
	"time"

	"gorm.io/gorm"
)

type Table struct {
	ID        uint64
	FileID    uint64
	Title     string
	Cells     []*Cell
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
