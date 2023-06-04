package file

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

func (x *KeyInput) Create() *model.File {
	fmt.Print("filename >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &model.File{
		Filename: filename,
		Attr: model.Attr{
			Owner: "me",
		},
		Tables: []*model.Table{{
			Title: "Sheet1",
		}},
	}
}

func (x *KeyInput) Read() *model.File {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.File{
		ID: id,
		// AttrID: id,
		Attr: model.Attr{
			ID: id,
		},
	}
}

func (x *KeyInput) Update() *model.File {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	fmt.Print("filename >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	fmt.Print("owner >")
	x.scanner.Scan()
	owner := x.scanner.Text()
	return &model.File{
		ID:       id,
		Filename: filename,
		Attr: model.Attr{
			// ID:    id,
			Owner: owner,
		},
	}
}

func (x *KeyInput) Delete() *model.File {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.ParseUint(x.scanner.Text(), 10, 64)
	return &model.File{
		ID: id,
	}
}

func (x *KeyInput) List() *model.File {
	return &model.File{}
}
