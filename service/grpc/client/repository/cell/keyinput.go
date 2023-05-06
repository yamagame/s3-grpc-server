package cell

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/entitiy/repository/model"
	"strconv"
)

type KeyInput struct {
	scanner *bufio.Scanner
}

func NewKeyInput(scanner *bufio.Scanner) *KeyInput {
	return &KeyInput{
		scanner: scanner,
	}
}

func (x *KeyInput) Create() *model.Cell {
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

func (x *KeyInput) Read() *model.Cell {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.Cell{
		ID: id,
	}
}

func (x *KeyInput) Update() *model.Cell {
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

func (x *KeyInput) Delete() *model.Cell {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.Cell{
		ID: id,
	}
}
