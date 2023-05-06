package app

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sample/s3-grpc-server/service/grpc/client/repository/cell"
	"sample/s3-grpc-server/service/grpc/client/repository/file"
	"sample/s3-grpc-server/service/grpc/client/repository/table"
	"sample/s3-grpc-server/service/grpc/client/storage"

	cellInfra "sample/s3-grpc-server/infra/repository/cell"
	fileInfra "sample/s3-grpc-server/infra/repository/file"
	tableInfra "sample/s3-grpc-server/infra/repository/table"
	storageInfra "sample/s3-grpc-server/infra/storage"

	"sample/s3-grpc-server/usecase"
	"strconv"

	"go.uber.org/dig"
)

type AppIn struct {
	dig.In
	FileRepository    *file.FileClientRepository
	TableRepository   *table.TableClientRepository
	CellRepository    *cell.CellClientRepository
	StorageRepository *storage.StorageClientRepository
	FileScanner       *file.KeyInput
	TableScanner      *table.KeyInput
	CellScanner       *cell.KeyInput
	StorageScanner    *storage.KeyInput
	DBWriter          *usecase.DBWriter
	Keyboard          *bufio.Scanner
}

type App struct {
	fileRepository    fileInfra.RepositoryInterface
	tableRepository   tableInfra.RepositoryInterface
	cellRepository    cellInfra.RepositoryInterface
	storageRepository storageInfra.RepositoryInterface
	fileScanner       *file.KeyInput
	tableScanner      *table.KeyInput
	cellScanner       *cell.KeyInput
	storageScanner    *storage.KeyInput
	dbwriter          *usecase.DBWriter
	keyboard          *bufio.Scanner
	ctx               context.Context
}

func NewApp(ctx context.Context, in AppIn) *App {
	return &App{
		fileRepository:    in.FileRepository,
		tableRepository:   in.TableRepository,
		cellRepository:    in.CellRepository,
		storageRepository: in.StorageRepository,
		fileScanner:       in.FileScanner,
		tableScanner:      in.TableScanner,
		cellScanner:       in.CellScanner,
		storageScanner:    in.StorageScanner,
		keyboard:          in.Keyboard,
		dbwriter:          in.DBWriter,
		ctx:               ctx,
	}
}

func (x *App) Start() {
	ctx := x.ctx
	tbl := []struct {
		name string
		call func() error
	}{
		{"", nil},
		{"CreateBucket", func() error {
			ent, _ := x.storageRepository.CreateBucket(ctx, x.storageScanner.CreateBucket())
			fmt.Println(ent)
			return nil
		}},
		{"ListBuckets", func() error {
			ent, _ := x.storageRepository.ListBuckets(ctx, x.storageScanner.ListBuckets())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"PutObject", func() error {
			ent, _ := x.storageRepository.PutObject(ctx, x.storageScanner.PutObject())
			fmt.Println(ent)
			return nil
		}},
		{"GetObject", func() error {
			ent, _ := x.storageRepository.GetObject(ctx, x.storageScanner.GetObject())
			fmt.Println(ent)
			return nil
		}},
		{"DeleteObject", func() error {
			ent, _ := x.storageRepository.DeleteObject(ctx, x.storageScanner.DeleteObject())
			fmt.Println(ent)
			return nil
		}},
		{"ListObjects", func() error {
			ent, _ := x.storageRepository.ListObjects(ctx, x.storageScanner.ListObjects())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"CreateFile", func() error {
			ent, _ := x.fileRepository.Create(ctx, x.fileScanner.Create())
			fmt.Println(ent)
			return nil
		}},
		{"ReadFile", func() error {
			ent, _ := x.fileRepository.Read(ctx, x.fileScanner.Read())
			fmt.Println(ent)
			if ent != nil {
				for _, v := range ent.Tables {
					fmt.Println(v)
				}
			}
			return nil
		}},
		{"UpdateFile", func() error {
			ent, _ := x.fileRepository.Update(ctx, x.fileScanner.Update())
			fmt.Println(ent)
			return nil
		}},
		{"DeleteFile", func() error {
			ent, _ := x.fileRepository.Delete(ctx, x.fileScanner.Delete())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"CreateTable", func() error {
			ent, _ := x.tableRepository.Create(ctx, x.tableScanner.Create())
			fmt.Println(ent)
			return nil
		}},
		{"ReadTable", func() error {
			ent, _ := x.tableRepository.Read(ctx, x.tableScanner.Read())
			fmt.Println(ent)
			if ent != nil {
				for _, v := range ent.Cells {
					fmt.Println(v)
				}
			}
			return nil
		}},
		{"UpdateTable", func() error {
			ent, _ := x.tableRepository.Update(ctx, x.tableScanner.Update())
			fmt.Println(ent)
			return nil
		}},
		{"DeleteTable", func() error {
			ent, _ := x.tableRepository.Delete(ctx, x.tableScanner.Delete())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"CreateCell", func() error {
			ent, _ := x.cellRepository.Create(ctx, x.cellScanner.Create())
			fmt.Println(ent)
			return nil
		}},
		{"ReadCell", func() error {
			ent, _ := x.cellRepository.Read(ctx, x.cellScanner.Read())
			fmt.Println(ent)
			return nil
		}},
		{"UpdateCell", func() error {
			ent, _ := x.cellRepository.Update(ctx, x.cellScanner.Update())
			fmt.Println(ent)
			return nil
		}},
		{"DeleteCell", func() error {
			ent, _ := x.cellRepository.Delete(ctx, x.cellScanner.Delete())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"WriteCSV", func() error {
			if err := x.dbwriter.CreateFakeCSV(ctx, "sample.csv"); err != nil {
				return err
			}
			x.dbwriter.CreateAll(ctx, "sample.csv")
			return nil
		}},
		{"", nil},
	}

	for {
		n := 1
		var ids []int
		for i, v := range tbl {
			if v.call == nil {
				fmt.Println(v.name)
			} else {
				fmt.Printf("%d: %s\n", n, v.name)
				n++
				ids = append(ids, i)
			}
		}
		fmt.Println("?: exit")
		fmt.Print("please enter >")
		x.keyboard.Scan()
		in := x.keyboard.Text()
		skip := false
		id, _ := strconv.Atoi(in)
		for i, v := range tbl {
			if !(id >= 1 && id <= len(ids)) {
				break
			}
			if i == ids[id-1] {
				if v.call != nil {
					if err := v.call(); err != nil {
						panic("fail")
					}
				}
				skip = true
			}
		}
		if !skip {
			fmt.Println("bye.")
			os.Exit(0)
		}
	}
}
