package model

import (
	"fmt"
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

func (x *Table) String() string {
	return fmt.Sprintf("{%d, %s}", x.ID, x.Title)
}
