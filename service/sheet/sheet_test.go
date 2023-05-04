package sheet

import (
	"os"
	"testing"
)

func TestSheet(t *testing.T) {
	csv := NewCSVSheet()
	// CSV 読み込みテスト
	{

		fp, _ := os.Open("./testdata/sample.csv")
		defer fp.Close()
		csv.Read(fp)
	}
	// CSV 書き出しテスト
	{
		fp, _ := os.CreateTemp("", "sample-new.csv")
		// fp, _ := os.Create("./testdata/sample-new.csv")
		defer fp.Close()
		csv.Save(fp)
	}

	{
		tests := []struct {
			row  int
			col  int
			want string
		}{
			{0, 0, "0"},
			{-1, 0, ""},
			{0, -1, ""},
			{10, 10, ""},
			{3, 4, "3"},
			{4, 4, ""},
			{3, 5, ""},
		}
		for _, tt := range tests {
			if csv.Cell(tt.row, tt.col) != tt.want {
				t.Errorf("expected %s got %s", tt.want, csv.Cell(tt.row, tt.col))
			}
		}
	}

	xlsx := NewXLSXSheet("Sheet1")
	// XLSX 読み込みテスト
	{
		fp, _ := os.Open("./testdata/sample.xlsx")
		defer fp.Close()
		xlsx.Read(fp)
	}
	// XLSX 書き出しテスト
	{
		fp, _ := os.CreateTemp("", "sample-new.xlsx")
		// fp, _ := os.Create("./testdata/sample-new.xlsx")
		defer fp.Close()
		xlsx.Save(fp)
	}

	if !csv.Compare(xlsx) {
		t.Errorf("differ csv.Cells(), xlsx.Cells()")
	}

	// 範囲外テスト
	{
		tests := []struct {
			row  int
			col  int
			want string
		}{
			{0, 0, "00"},
			{-1, 0, ""},
			{0, -1, ""},
			{5, 5, "55"},
			{3, 4, "34"},
			{4, 4, "44"},
			{3, 5, "35"},
		}
		for _, tt := range tests {
			csv.SetCell(tt.row, tt.col, tt.want)
			if csv.Cell(tt.row, tt.col) != tt.want {
				t.Errorf("expected %s got %s", tt.want, csv.Cell(tt.row, tt.col))
			}
			xlsx.SetCell(tt.row, tt.col, tt.want)
			if xlsx.Cell(tt.row, tt.col) != tt.want {
				t.Errorf("expected %s got %s", tt.want, xlsx.Cell(tt.row, tt.col))
			}
		}
	}
	// CSV 書き出しテスト
	{
		fp, _ := os.CreateTemp("", "sample-range.csv")
		defer fp.Close()
		csv.Save(fp)
	}
	// XLSX 書き出しテスト
	{
		fp, _ := os.CreateTemp("", "sample-range.xlsx")
		defer fp.Close()
		xlsx.Save(fp)
	}

	if !csv.Compare(xlsx) {
		t.Errorf("expected csv same xlsx")
	}

	xlsx2 := NewXLSXSheet("Sheet1")
	// XLSX 読み込みテスト
	{
		fp, _ := os.Open("./testdata/sample-range.xlsx")
		defer fp.Close()
		xlsx2.Read(fp)
	}

	// テストデータと比較
	if !xlsx2.Compare(xlsx) {
		t.Errorf("expected xlsx2 same xlsx")
	}

	csv2 := NewCSVSheet()
	// CSV 読み込みテスト
	{
		fp, _ := os.Open("./testdata/sample-range.csv")
		defer fp.Close()
		csv2.Read(fp)
	}

	// テストデータと比較
	if !csv2.Compare(csv) {
		t.Errorf("expected csv2 same csv")
	}
}
