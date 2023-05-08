package sheet

import (
	"io"

	"github.com/xuri/excelize/v2"
)

type XLSXSheet struct {
	Sheet
	Sheetname string
}

func NewXLSXSheet(sheet string) *XLSXSheet {
	return &XLSXSheet{
		Sheetname: sheet,
	}
}

func (x *XLSXSheet) Read(r io.Reader) error {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	defer f.Close()
	rows, err := f.GetRows(x.Sheetname)
	if err != nil {
		return err
	}
	x.Initialized = false
	x.Sheet.Cells = rows
	return nil
}

func (x *XLSXSheet) Save(w io.Writer) (err error) {
	f := excelize.NewFile()
	defer func() error {
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
		return nil
	}()
	sw, err := f.NewStreamWriter(x.Sheetname)
	if err != nil {
		return err
	}
	for rowID := 0; rowID < len(x.Sheet.Cells); rowID++ {
		row := make([]interface{}, len(x.Sheet.Cells[rowID]))
		for colID := 0; colID < len(x.Sheet.Cells[rowID]); colID++ {
			row[colID] = x.Sheet.Cells[rowID][colID]
		}
		cell, err := excelize.CoordinatesToCellName(1, rowID+1)
		if err != nil {
			return err
		}
		if err := sw.SetRow(cell, row); err != nil {
			return err
		}
	}
	if err := sw.Flush(); err != nil {
		return err
	}
	if err := f.Write(w); err != nil {
		return err
	}
	return nil
}

func (x *XLSXSheet) GetCells() [][]string {
	return x.Sheet.Cells
}
