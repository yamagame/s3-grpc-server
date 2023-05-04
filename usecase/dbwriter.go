package usecase

import (
	"context"
	"fmt"
	repositoryModel "sample/s3-grpc-server/entitiy/repository/model"
	storageModel "sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/file"
	"sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/infra/storage"
	"sample/s3-grpc-server/service/sheet"
	"strings"
	"time"
)

type dbWriter struct {
	fileRepository    file.RepositoryInterface
	tableRepository   table.RepositoryInterface
	cellRepository    cell.RepositoryInterface
	storageRepository storage.RepositoryInterface
	sheet             sheet.SpreadsheetInterface
}

func NewDBWriter(
	fileRepository file.RepositoryInterface,
	tableRepository table.RepositoryInterface,
	cellRepository cell.RepositoryInterface,
	storageRepository storage.RepositoryInterface,
	sheet sheet.SpreadsheetInterface,
) *dbWriter {
	return &dbWriter{
		fileRepository:    fileRepository,
		tableRepository:   tableRepository,
		cellRepository:    cellRepository,
		storageRepository: storageRepository,
		sheet:             sheet,
	}
}

func fakeCSV(num int, sheet sheet.SpreadsheetInterface) string {
	csv := sheet
	for i := 0; i < num; i++ {
		csv.SetCell(i, 0, fmt.Sprintf("file%05d.csv", i))
	}
	return csv.CSVString()
}

func (x *dbWriter) CreateFakeCSV(ctx context.Context, filename string) error {
	fmt.Println(time.Now(), "makecsv")
	req := &storageModel.PutObject{
		Key:     filename,
		Content: fakeCSV(10000, x.sheet),
	}
	_, err := x.storageRepository.PutObject(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (x *dbWriter) CreateAll(ctx context.Context, filename string) error {
	fmt.Println(time.Now(), "start")
	req := &storageModel.GetObject{
		Key: filename,
	}
	ent, err := x.storageRepository.GetObject(ctx, req)
	if err != nil {
		return err
	}
	csv := x.sheet
	csv.Read(strings.NewReader(ent.Content))
	records := csv.GetCells()
	fmt.Println(time.Now(), "write")
	for _, record := range records {
		file := &repositoryModel.File{
			Filename: record[0],
		}
		_, err := x.fileRepository.Create(ctx, file)
		if err != nil {
			return err
		}
	}
	fmt.Println(time.Now(), "end")
	return nil
}
