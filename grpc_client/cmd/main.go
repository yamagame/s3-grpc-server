package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"sample/s3-grpc-server/grpc_client/cmd/service"
	repository_service "sample/s3-grpc-server/grpc_client/service/repository"
	storage_service "sample/s3-grpc-server/grpc_client/service/storage"
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
	repositoryClient := repository_service.NewRepositoryClient(grpc_server.NewRepositoryClient(conn), repository_service.NewScanner(keyboard))
	storageClient := storage_service.NewStorageClient(grpc_server.NewStorageClient(conn), storage_service.NewScanner(keyboard))
	client := service.NewClientService(storageClient, repositoryClient)

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
