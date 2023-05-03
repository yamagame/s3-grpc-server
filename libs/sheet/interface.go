package sheet

import "io"

type SpreadsheetInterface interface {
	Read(r io.Reader) error
	Save(w io.Writer) error
	GetCells() [][]string
}
