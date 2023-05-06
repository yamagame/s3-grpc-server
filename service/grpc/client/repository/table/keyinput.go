package table

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

func (x *KeyInput) Create() *model.Table {
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

func (x *KeyInput) Read() *model.Table {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.Table{
		ID: id,
	}
}

func (x *KeyInput) Update() *model.Table {
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

func (x *KeyInput) Delete() *model.Table {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 0, 32)
	return &model.Table{
		ID: id,
	}
}
