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
	tableID, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	fmt.Print("row >")
	x.scanner.Scan()
	row, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	fmt.Print("col >")
	x.scanner.Scan()
	col, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	fmt.Print("text >")
	x.scanner.Scan()
	text := x.scanner.Text()
	return &model.Cell{
		TableID: tableID,
		Row:     row,
		Col:     col,
		Text:    text,
	}
}

func (x *keyInput) Read() *model.Cell {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.Cell{
		ID: id,
	}
}

func (x *keyInput) Update() *model.Cell {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 0, 64)
	fmt.Print("text >")
	x.scanner.Scan()
	text := x.scanner.Text()
	return &model.Cell{
		ID:   id,
		Text: text,
	}
}

func (x *keyInput) Delete() *model.Cell {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.Cell{
		ID: id,
	}
}
