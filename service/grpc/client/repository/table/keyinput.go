package table

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/entitiy/repository/model"
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

func (x *keyInput) Create() *model.Table {
	fmt.Print("file_id >")
	x.scanner.Scan()
	fileID, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	fmt.Print("title >")
	x.scanner.Scan()
	title := x.scanner.Text()
	return &model.Table{
		FileID: fileID,
		Title:  title,
	}
}

func (x *keyInput) Read() *model.Table {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.Table{
		ID: id,
	}
}

func (x *keyInput) Update() *model.Table {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	fmt.Print("title >")
	x.scanner.Scan()
	title := x.scanner.Text()
	return &model.Table{
		ID:    id,
		Title: title,
	}
}

func (x *keyInput) Delete() *model.Table {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 0, 32)
	return &model.Table{
		ID: id,
	}
}
