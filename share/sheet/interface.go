package sheet

import "io"

type SpreadsheetInterface interface {
	Read(r io.Reader) error
	Save(w io.Writer) error
	GetCells() [][]string
	SetCell(row, col int, str string)
	Cell(row, col int) string
	Compare(t SpreadsheetInterface) bool
	CSVString() string
}
