package model

import (
	"sample/s3-grpc-server/share"
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID        uint64
	Filename  string
	FValue    float64
	Attr      Attr `gorm:"foreignKey:ID"`
	Tables    []*Table
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (x *File) String() string {
	return share.JsonString(x)
	// return fmt.Sprintf("{%d, %s, %s}", x.ID, x.Filename, x.Attr.Owner)
}
