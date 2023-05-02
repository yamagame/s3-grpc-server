package repository

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/infra/repository/fileinfo/model"
	"strconv"
)

type RepositoryScannerInterface interface {
	CreateFileInfo() *model.FileInfo
	ReadFileInfo() *model.FileInfo
	UpdateFileInfo() *model.FileInfo
	DeleteFileInfo() *model.FileInfo
}

type RepositoryScanner struct {
	scanner *bufio.Scanner
}

func NewScanner(scanner *bufio.Scanner) *RepositoryScanner {
	return &RepositoryScanner{
		scanner: scanner,
	}
}

func (x *RepositoryScanner) CreateFileInfo() *model.FileInfo {
	fmt.Print("content >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &model.FileInfo{
		Filename: filename,
	}
}

func (x *RepositoryScanner) ReadFileInfo() *model.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.FileInfo{
		ID: int64(id),
	}
}

func (x *RepositoryScanner) UpdateFileInfo() *model.FileInfo {
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

func (x *RepositoryScanner) DeleteFileInfo() *model.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &model.FileInfo{
		ID: int64(id),
	}
}
