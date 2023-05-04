package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"sample/s3-grpc-server/service/grpc/client/repository/cell"
	"sample/s3-grpc-server/service/grpc/client/repository/file"
	"sample/s3-grpc-server/service/grpc/client/repository/table"
	"sample/s3-grpc-server/service/grpc/client/storage"
	"sample/s3-grpc-server/service/sheet"
	"sample/s3-grpc-server/usecase"

	grpc_server "sample/s3-grpc-server/proto/grpc_server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	fmt.Println("start gRPC Client.")

	// 1. キーボード入力準備
	keyboard := bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	address := os.Getenv("GRPC_HOST")
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	fileRepository := file.NewFileRepository(grpc_server.NewFileRepositoryClient(conn))
	tableRepository := table.NewTableRepository(grpc_server.NewTableRepositoryClient(conn))
	cellRepository := cell.NewCellRepository(grpc_server.NewCellRepositoryClient(conn))
	storageRepository := storage.NewRepository(grpc_server.NewStorageRepositoryClient(conn))

	fileScanner := file.NewKeyInput(keyboard)
	tableScanner := table.NewKeyInput(keyboard)
	cellScanner := cell.NewKeyInput(keyboard)
	storageScanner := storage.NewKeyInput(keyboard)

	dbwriter := usecase.NewDBWriter(
		fileRepository,
		tableRepository,
		cellRepository,
		storageRepository,
		sheet.NewCSVSheet(),
	)

	tbl := []struct {
		name string
		call func() error
	}{
		{"", nil},
		{"CreateBucket", func() error {
			ent, _ := storageRepository.CreateBucket(ctx, storageScanner.CreateBucket())
			fmt.Println(ent)
			return nil
		}},
		{"ListBuckets", func() error {
			ent, _ := storageRepository.ListBuckets(ctx, storageScanner.ListBuckets())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"PutObject", func() error {
			ent, _ := storageRepository.PutObject(ctx, storageScanner.PutObject())
			fmt.Println(ent)
			return nil
		}},
		{"GetObject", func() error {
			ent, _ := storageRepository.GetObject(ctx, storageScanner.GetObject())
			fmt.Println(ent)
			return nil
		}},
		{"DeleteObject", func() error {
			ent, _ := storageRepository.DeleteObject(ctx, storageScanner.DeleteObject())
			fmt.Println(ent)
			return nil
		}},
		{"ListObjects", func() error {
			ent, _ := storageRepository.ListObjects(ctx, storageScanner.ListObjects())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"CreateFile", func() error {
			ent, _ := fileRepository.Create(ctx, fileScanner.Create())
			fmt.Println(ent)
			return nil
		}},
		{"ReadFile", func() error {
			ent, _ := fileRepository.Read(ctx, fileScanner.Read())
			fmt.Println(ent)
			for _, v := range ent.Tables {
				fmt.Println(v)
			}
			return nil
		}},
		{"UpdateFile", func() error {
			ent, _ := fileRepository.Update(ctx, fileScanner.Update())
			fmt.Println(ent)
			return nil
		}},
		{"DeleteFile", func() error {
			ent, _ := fileRepository.Delete(ctx, fileScanner.Delete())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"CreateTable", func() error {
			ent, _ := tableRepository.Create(ctx, tableScanner.Create())
			fmt.Println(ent)
			return nil
		}},
		{"ReadTable", func() error {
			ent, _ := tableRepository.Read(ctx, tableScanner.Read())
			fmt.Println(ent)
			for _, v := range ent.Cells {
				fmt.Println(v)
			}
			return nil
		}},
		{"UpdateTable", func() error {
			ent, _ := tableRepository.Update(ctx, tableScanner.Update())
			fmt.Println(ent)
			return nil
		}},
		{"DeleteTable", func() error {
			ent, _ := tableRepository.Delete(ctx, tableScanner.Delete())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"CreateCell", func() error {
			ent, _ := cellRepository.Create(ctx, cellScanner.Create())
			fmt.Println(ent)
			return nil
		}},
		{"ReadCell", func() error {
			ent, _ := cellRepository.Read(ctx, cellScanner.Read())
			fmt.Println(ent)
			return nil
		}},
		{"UpdateCell", func() error {
			ent, _ := cellRepository.Update(ctx, cellScanner.Update())
			fmt.Println(ent)
			return nil
		}},
		{"DeleteCell", func() error {
			ent, _ := cellRepository.Delete(ctx, cellScanner.Delete())
			fmt.Println(ent)
			return nil
		}},
		{"", nil},
		{"WriteCSV", func() error {
			if err := dbwriter.CreateFakeCSV(ctx, "sample.csv"); err != nil {
				return err
			}
			dbwriter.CreateAll(ctx, "sample.csv")
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
		keyboard.Scan()
		in := keyboard.Text()
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
