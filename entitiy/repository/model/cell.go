package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Cell struct {
	ID        uint64
	Row       uint64 `gorm:"index:idx_position,unique"`
	Col       uint64 `gorm:"index:idx_position,unique"`
	TableID   uint64 `gorm:"index:idx_position,unique"`
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (x *Cell) String() string {
	return fmt.Sprintf("{%d, %s}", x.ID, x.Text)
}
