package sheet

import (
	"encoding/csv"
	"io"
)

type CSVSheet struct {
	Sheet
}

func NewCSVSheet() *CSVSheet {
	return &CSVSheet{
		Sheet: Sheet{
			Cells: [][]string{},
		},
	}
}

func (x *CSVSheet) Read(r io.Reader) error {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}
	x.Initialized = false
	x.Sheet.Cells = rows
	return nil
}

func (x *CSVSheet) Save(w io.Writer) error {
	writer := csv.NewWriter(w)
	return writer.WriteAll(x.Sheet.Cells)
}

func (x *CSVSheet) GetCells() [][]string {
	return x.Sheet.Cells
}
