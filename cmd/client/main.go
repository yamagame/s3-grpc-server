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
		call func()
	}{
		{"", nil},
		{"CreateBucket", func() {
			ent, _ := storageRepository.CreateBucket(ctx, storageScanner.CreateBucket())
			fmt.Println(ent)
		}},
		{"ListBuckets", func() {
			ent, _ := storageRepository.ListBuckets(ctx, storageScanner.ListBuckets())
			fmt.Println(ent)
		}},
		{"", nil},
		{"PutObject", func() {
			ent, _ := storageRepository.PutObject(ctx, storageScanner.PutObject())
			fmt.Println(ent)
		}},
		{"GetObject", func() {
			ent, _ := storageRepository.GetObject(ctx, storageScanner.GetObject())
			fmt.Println(ent)
		}},
		{"DeleteObject", func() {
			ent, _ := storageRepository.DeleteObject(ctx, storageScanner.DeleteObject())
			fmt.Println(ent)
		}},
		{"ListObjects", func() {
			ent, _ := storageRepository.ListObjects(ctx, storageScanner.ListObjects())
			fmt.Println(ent)
		}},
		{"", nil},
		{"CreateFile", func() {
			ent, _ := fileRepository.Create(ctx, fileScanner.Create())
			fmt.Println(ent)
		}},
		{"ReadFile", func() {
			ent, _ := fileRepository.Read(ctx, fileScanner.Read())
			fmt.Println(ent)
			for _, v := range ent.Tables {
				fmt.Println(v)
			}
		}},
		{"UpdateFile", func() {
			ent, _ := fileRepository.Update(ctx, fileScanner.Update())
			fmt.Println(ent)
		}},
		{"DeleteFile", func() {
			ent, _ := fileRepository.Delete(ctx, fileScanner.Delete())
			fmt.Println(ent)
		}},
		{"", nil},
		{"CreateTable", func() {
			ent, _ := tableRepository.Create(ctx, tableScanner.Create())
			fmt.Println(ent)
		}},
		{"ReadTable", func() {
			ent, _ := tableRepository.Read(ctx, tableScanner.Read())
			fmt.Println(ent)
			for _, v := range ent.Cells {
				fmt.Println(v)
			}
		}},
		{"UpdateTable", func() {
			ent, _ := tableRepository.Update(ctx, tableScanner.Update())
			fmt.Println(ent)
		}},
		{"DeleteTable", func() {
			ent, _ := tableRepository.Delete(ctx, tableScanner.Delete())
			fmt.Println(ent)
		}},
		{"", nil},
		{"CreateCell", func() {
			ent, _ := cellRepository.Create(ctx, cellScanner.Create())
			fmt.Println(ent)
		}},
		{"ReadCell", func() {
			ent, _ := cellRepository.Read(ctx, cellScanner.Read())
			fmt.Println(ent)
		}},
		{"UpdateCell", func() {
			ent, _ := cellRepository.Update(ctx, cellScanner.Update())
			fmt.Println(ent)
		}},
		{"DeleteCell", func() {
			ent, _ := cellRepository.Delete(ctx, cellScanner.Delete())
			fmt.Println(ent)
		}},
		{"", nil},
		{"WriteCSV", func() {
			dbwriter.CreateFakeCSV(ctx, "sample.csv")
			dbwriter.CreateAll(ctx, "sample.csv")
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
					v.call()
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
