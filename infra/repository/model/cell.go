package model

type Cell struct {
	ID      int64 `gorm:"primaryKey"`
	Row     int64 `gorm:"index:idx_position,unique"`
	Col     int64 `gorm:"index:idx_position,unique"`
	TableID int64 `gorm:"index:idx_position,unique"`
	Text    string
}
