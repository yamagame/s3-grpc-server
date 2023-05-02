package model

type Table struct {
	ID     int64
	Title  string
	FileID int64
	Cells  []Cell
}
