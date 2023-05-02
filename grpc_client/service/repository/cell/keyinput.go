package cell

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/infra/repository/model"
	"strconv"
)

type keyInput struct {
	scanner *bufio.Scanner
}

func NewKeyInput(scanner *bufio.Scanner) *keyInput {
	return &keyInput{
		scanner: scanner,
	}
}

func (x *keyInput) Create() *model.Cell {
	fmt.Print("table_id >")
	x.scanner.Scan()
	tableID, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("row >")
	x.scanner.Scan()
	row, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("col >")
	x.scanner.Scan()
	col, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("text >")
	x.scanner.Scan()
	text := x.scanner.Text()
	return &model.Cell{
		TableID: int64(tableID),
		Row:     int64(row),
		Col:     int64(col),
		Text:    text,
	}
}

func (x *keyInput) Read() *model.Cell {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.Cell{
		ID: int64(id),
	}
}

func (x *keyInput) Update() *model.Cell {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("text >")
	x.scanner.Scan()
	text := x.scanner.Text()
	return &model.Cell{
		ID:   int64(id),
		Text: text,
	}
}

func (x *keyInput) Delete() *model.Cell {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.Cell{
		ID: int64(id),
	}
}
