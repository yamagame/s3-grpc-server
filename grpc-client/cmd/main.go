package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	scanner "sample/s3-grpc-server/grpc-client/service/storage"
	storage "sample/s3-grpc-server/proto/grpc-server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	fmt.Println("start gRPC Client.")

	// 1. キーボード入力準備
	keyboard := bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	address := "localhost:50051"
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
	client := scanner.NewStorageClient(storage.NewStorageClient(conn), scanner.NewScanner(keyboard))

	for {
		fmt.Println("1: CreateBucket")
		fmt.Println("2: ListBuckets")
		fmt.Println("3: PutObject")
		fmt.Println("4: GetObject")
		fmt.Println("5: DeleteObject")
		fmt.Println("6: ListObjects")
		fmt.Println("9: exit")
		fmt.Print("please enter >")

		keyboard.Scan()
		in := keyboard.Text()

		switch in {
		case "1":
			ent, _ := client.CreateBucket(ctx)
			fmt.Println(ent)
		case "2":
			ent, _ := client.ListBuckets(ctx)
			fmt.Println(ent)
		case "3":
			ent, _ := client.PutObject(ctx)
			fmt.Println(ent)
		case "4":
			ent, _ := client.GetObject(ctx)
			fmt.Println(ent)
		case "5":
			ent, _ := client.DeleteObject(ctx)
			fmt.Println(ent)
		case "6":
			ent, _ := client.ListObjects(ctx)
			fmt.Println(ent)

		default:
			fmt.Println("bye.")
			os.Exit(0)
		}
	}
}
