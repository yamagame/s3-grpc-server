package table

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

func (x *keyInput) Create() *model.Table {
	fmt.Print("file_id >")
	x.scanner.Scan()
	fileID, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("title >")
	x.scanner.Scan()
	title := x.scanner.Text()
	return &model.Table{
		FileID: int64(fileID),
		Title:  title,
	}
}

func (x *keyInput) Read() *model.Table {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.Table{
		ID: int64(id),
	}
}

func (x *keyInput) Update() *model.Table {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("title >")
	x.scanner.Scan()
	title := x.scanner.Text()
	return &model.Table{
		ID:    int64(id),
		Title: title,
	}
}

func (x *keyInput) Delete() *model.Table {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.Table{
		ID: int64(id),
	}
}
