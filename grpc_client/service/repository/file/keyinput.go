package file

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

func (x *keyInput) Create() *model.File {
	fmt.Print("filename >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &model.File{
		Filename: filename,
	}
}

func (x *keyInput) Read() *model.File {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.File{
		ID: id,
	}
}

func (x *keyInput) Update() *model.File {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	fmt.Print("filename >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &model.File{
		ID:       id,
		Filename: filename,
	}
}

func (x *keyInput) Delete() *model.File {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.File{
		ID: id,
	}
}
