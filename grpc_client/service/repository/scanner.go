package repository

import (
	"bufio"
	"fmt"
	"sample/s3-grpc-server/infra/repository"
	"strconv"
)

type RepositoryScannerInterface interface {
	CreateFileInfo() *repository.FileInfo
	ReadFileInfo() *repository.FileInfo
	UpdateFileInfo() *repository.FileInfo
	DeleteFileInfo() *repository.FileInfo
}

type RepositoryScanner struct {
	scanner *bufio.Scanner
}

func NewScanner(scanner *bufio.Scanner) *RepositoryScanner {
	return &RepositoryScanner{
		scanner: scanner,
	}
}

func (x *RepositoryScanner) CreateFileInfo() *repository.FileInfo {
	fmt.Print("content >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &repository.FileInfo{
		Filename: filename,
	}
}

func (x *RepositoryScanner) ReadFileInfo() *repository.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &repository.FileInfo{
		ID: int64(id),
	}
}

func (x *RepositoryScanner) UpdateFileInfo() *repository.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	fmt.Print("content >")
	x.scanner.Scan()
	filename := x.scanner.Text()
	return &repository.FileInfo{
		ID:       int64(id),
		Filename: filename,
	}
}

func (x *RepositoryScanner) DeleteFileInfo() *repository.FileInfo {
	fmt.Print("ID >")
	x.scanner.Scan()
	id, _ := strconv.Atoi(x.scanner.Text())
	return &repository.FileInfo{
		ID: int64(id),
	}
}
