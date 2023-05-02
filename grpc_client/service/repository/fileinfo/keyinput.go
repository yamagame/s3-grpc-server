package fileinfo

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/infra/repository/model"
	"strconv"
)

type FileInfoKeyInput struct {
	scanner *bufio.Scanner
}

func NewKeyInput(scanner *bufio.Scanner) *FileInfoKeyInput {
	return &FileInfoKeyInput{
		scanner: scanner,
	}
}

func (x *FileInfoKeyInput) CreateFileInfo() *model.FileInfo {
	fmt.Print("content >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &model.FileInfo{
		Filename: filename,
	}
}

func (x *FileInfoKeyInput) ReadFileInfo() *model.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.FileInfo{
		ID: int64(id),
	}
}

func (x *FileInfoKeyInput) UpdateFileInfo() *model.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("content >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &model.FileInfo{
		ID:       int64(id),
		Filename: filename,
	}
}

func (x *FileInfoKeyInput) DeleteFileInfo() *model.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.FileInfo{
		ID: int64(id),
	}
}
