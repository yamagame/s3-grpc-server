package sheet

import (
	"encoding/csv"
	"reflect"
	"strings"
)

type Sheet struct {
	Initialized bool
	rowmax      int
	colmax      int
	Cells       [][]string
}

func (x *Sheet) init() {
	if !x.Initialized {
		x.rowmax = len(x.Cells)
		x.colmax = 0
		for _, row := range x.Cells {
			if x.colmax < len(row) {
				x.colmax = len(row)
			}
		}
		x.Initialized = true
	}
}

func (x *Sheet) Cell(row, col int) string {
	x.init()
	if row >= 0 && row < len(x.Cells) {
		if col >= 0 && col < len(x.Cells[row]) {
			return x.Cells[row][col]
		}
	}
	return ""
}

func (x *Sheet) SetCell(row, col int, str string) {
	if row < 0 || col < 0 {
		return
	}
	x.init()
	if x.rowmax <= row {
		for i := x.rowmax; i <= row; i++ {
			x.Cells = append(x.Cells, []string{})
		}
		x.rowmax = row + 1
	}
	if len(x.Cells[row]) <= col {
		for i := len(x.Cells[row]); i <= col; i++ {
			x.Cells[row] = append(x.Cells[row], "")
			x.Initialized = false
		}
		if x.colmax <= col {
			x.colmax = col + 1
		}
	}
	x.Cells[row][col] = str
}

func (x *Sheet) Compare(t SpreadsheetInterface) bool {
	return reflect.DeepEqual(x.Cells, t.GetCells())
}

func (x *Sheet) CSVString() string {
	b := new(strings.Builder)
	writer := csv.NewWriter(b)
	writer.WriteAll(x.Cells)
	return b.String()
}
