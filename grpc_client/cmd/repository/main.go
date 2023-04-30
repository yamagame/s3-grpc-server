package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"sample/s3-grpc-server/grpc_client/cmd/repository/internal"
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
	client := internal.NewRepositoryClient(grpc_server.NewRepositoryClient(conn), internal.NewScanner(keyboard))

	for {
		fmt.Println("1: CreateFileInfo")
		fmt.Println("2: ReadFileInfo")
		fmt.Println("3: UpdateFileInfo")
		fmt.Println("4: DeleteFileInfo")
		fmt.Println("?: exit")
		fmt.Print("please enter >")

		keyboard.Scan()
		in := keyboard.Text()

		switch in {
		case "1":
			ent, _ := client.CreateFileInfo(ctx)
			fmt.Println(ent)
		case "2":
			ent, _ := client.ReadFileInfo(ctx)
			fmt.Println(ent)
		case "3":
			ent, _ := client.UpdateFileInfo(ctx)
			fmt.Println(ent)
		case "4":
			ent, _ := client.DeleteFileInfo(ctx)
			fmt.Println(ent)

		default:
			fmt.Println("bye.")
			os.Exit(0)
		}
	}
}
