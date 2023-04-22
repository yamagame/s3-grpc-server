package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"sample/s3-grpc-server/grpc-client/service"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	fmt.Println("start gRPC Client.")

	// 1. 標準入力から文字列を受け取るスキャナを用意
	scanner := bufio.NewScanner(os.Stdin)

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
	client := aws.NewAwsClient(conn)

	// 5. gRPCサービスを生成
	storage := service.NewAwsService(client, service.NewAwsScanner(scanner))

	for {
		fmt.Println("1: CreateBucket")
		fmt.Println("2: ListBuckets")
		fmt.Println("3: PutObject")
		fmt.Println("4: GetObject")
		fmt.Println("5: DeleteObject")
		fmt.Println("6: ListObjects")
		fmt.Println("9: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			ent, _ := storage.CreateBucket(ctx)
			fmt.Println(ent)
		case "2":
			ent, _ := storage.ListBuckets(ctx)
			fmt.Println(ent)
		case "3":
			ent, _ := storage.PutObject(ctx)
			fmt.Println(ent)
		case "4":
			ent, _ := storage.GetObject(ctx)
			fmt.Println(ent)
		case "5":
			ent, _ := storage.DeleteObject(ctx)
			fmt.Println(ent)
		case "6":
			ent, _ := storage.ListObjects(ctx)
			fmt.Println(ent)

		default:
			fmt.Println("bye.")
			os.Exit(0)
		}
	}
}
