package model

type Attr struct {
	ID     uint64 `gorm:"primaryKey"`
	FileID uint64 `gorm:"index"`
	Owner  string
}
