package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"sample/s3-grpc-server/grpc_client/service"
	"sample/s3-grpc-server/grpc_client/service/repository/cell"
	"sample/s3-grpc-server/grpc_client/service/repository/fileinfo"
	"sample/s3-grpc-server/grpc_client/service/repository/table"
	"sample/s3-grpc-server/grpc_client/service/storage"
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
	fileInfoClient := fileinfo.NewFileInfoClient(grpc_server.NewRepositoryClient(conn), fileinfo.NewKeyInput(keyboard))
	tableClient := table.NewTableClient(grpc_server.NewRepositoryClient(conn), table.NewKeyInput(keyboard))
	cellClient := cell.NewCellClient(grpc_server.NewRepositoryClient(conn), cell.NewKeyInput(keyboard))
	storageClient := storage.NewStorageClient(grpc_server.NewStorageClient(conn), storage.NewKeyInput(keyboard))
	client := service.NewClientService(storageClient, fileInfoClient, tableClient, cellClient)

	tbl := []struct {
		name string
		call func()
	}{
		{"CreateBucket", func() {
			ent, _ := client.CreateBucket(ctx)
			fmt.Println(ent)
		}},
		{"ListBuckets", func() {
			ent, _ := client.ListBuckets(ctx)
			fmt.Println(ent)
		}},
		{"PutObject", func() {
			ent, _ := client.PutObject(ctx)
			fmt.Println(ent)
		}},
		{"GetObject", func() {
			ent, _ := client.GetObject(ctx)
			fmt.Println(ent)
		}},
		{"DeleteObject", func() {
			ent, _ := client.DeleteObject(ctx)
			fmt.Println(ent)
		}},
		{"ListObjects", func() {
			ent, _ := client.ListObjects(ctx)
			fmt.Println(ent)
		}},
		{"CreateFileInfo", func() {
			ent, _ := client.CreateFileInfo(ctx)
			fmt.Println(ent)
		}},
		{"ReadFileInfo", func() {
			ent, _ := client.ReadFileInfo(ctx)
			fmt.Println(ent)
		}},
		{"UpdateFileInfo", func() {
			ent, _ := client.UpdateFileInfo(ctx)
			fmt.Println(ent)
		}},
		{"DeleteFileInfo", func() {
			ent, _ := client.DeleteFileInfo(ctx)
			fmt.Println(ent)
		}},
		{"CreateTable", func() {
			ent, _ := client.CreateTable(ctx)
			fmt.Println(ent)
		}},
		{"ReadTable", func() {
			ent, _ := client.ReadTable(ctx)
			fmt.Println(ent)
		}},
		{"UpdateTable", func() {
			ent, _ := client.UpdateTable(ctx)
			fmt.Println(ent)
		}},
		{"DeleteTable", func() {
			ent, _ := client.DeleteTable(ctx)
			fmt.Println(ent)
		}},
		{"CreateCell", func() {
			ent, _ := client.CreateCell(ctx)
			fmt.Println(ent)
		}},
		{"ReadCell", func() {
			ent, _ := client.ReadCell(ctx)
			fmt.Println(ent)
		}},
		{"UpdateCell", func() {
			ent, _ := client.UpdateCell(ctx)
			fmt.Println(ent)
		}},
		{"DeleteCell", func() {
			ent, _ := client.DeleteCell(ctx)
			fmt.Println(ent)
		}},
	}

	for {
		for i, v := range tbl {
			fmt.Printf("%d: %s\n", i+1, v.name)
		}
		fmt.Println("?: exit")
		fmt.Print("please enter >")
		keyboard.Scan()
		in := keyboard.Text()
		skip := false
		for i, v := range tbl {
			if fmt.Sprintf("%d", i+1) == in {
				v.call()
				skip = true
			}
		}
		if !skip {
			fmt.Println("bye.")
			os.Exit(0)
		}
	}
}
