package fileinfo

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

func (x *keyInput) CreateFileInfo() *model.FileInfo {
	fmt.Print("filename >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &model.FileInfo{
		Filename: filename,
	}
}

func (x *keyInput) ReadFileInfo() *model.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.FileInfo{
		ID: int64(id),
	}
}

func (x *keyInput) UpdateFileInfo() *model.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("filename >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &model.FileInfo{
		ID:       int64(id),
		Filename: filename,
	}
}

func (x *keyInput) DeleteFileInfo() *model.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.FileInfo{
		ID: int64(id),
	}
}
