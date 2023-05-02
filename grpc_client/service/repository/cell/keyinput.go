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
	fmt.Print("text >")
	x.scanner.Scan()
	text := x.scanner.Text()
	return &model.Cell{
		Text: text,
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