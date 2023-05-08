package main

import (
	"context"
	"fmt"
	"log"

	"sample/s3-grpc-server/cmd/fxclient/app"
	"sample/s3-grpc-server/cmd/fxclient/constructor"
	"sample/s3-grpc-server/service/grpc/client/repository/cell"
	"sample/s3-grpc-server/service/grpc/client/repository/file"
	"sample/s3-grpc-server/service/grpc/client/repository/table"
	"sample/s3-grpc-server/service/grpc/client/storage"
	"sample/s3-grpc-server/share/sheet"

	grpc_server "sample/s3-grpc-server/proto/grpc_server"
)

func main() {
	ctx := context.Background()

	fmt.Println("start gRPC Client.")

	// 1. キーボード入力準備
	keyboard := constructor.NewKeyboardScanner()

	// 2. gRPCサーバーとのコネクションを確立
	// address := os.Getenv("GRPC_HOST")
	conn, err := constructor.NewGRPCConnection()
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	fileRepository := file.NewFileRepository(grpc_server.NewFileRepositoryClient(conn))
	tableRepository := table.NewTableRepository(grpc_server.NewTableRepositoryClient(conn))
	cellRepository := cell.NewCellRepository(grpc_server.NewCellRepositoryClient(conn))
	storageRepository := storage.NewStorageRepository(grpc_server.NewStorageRepositoryClient(conn))

	fileScanner := file.NewKeyInput(keyboard)
	tableScanner := table.NewKeyInput(keyboard)
	cellScanner := cell.NewKeyInput(keyboard)
	storageScanner := storage.NewKeyInput(keyboard)

	dbwriter := constructor.NewDBWriter(constructor.DBWriterIn{
		FileRepository:    fileRepository,
		TableRepository:   tableRepository,
		CellRepository:    cellRepository,
		StorageRepository: storageRepository,
		Sheet:             sheet.NewCSVSheet(),
	})

	app := app.NewApp(ctx, app.AppIn{
		FileRepository:    fileRepository,
		TableRepository:   tableRepository,
		CellRepository:    cellRepository,
		StorageRepository: storageRepository,
		FileScanner:       fileScanner,
		TableScanner:      tableScanner,
		CellScanner:       cellScanner,
		StorageScanner:    storageScanner,
		Keyboard:          keyboard,
		DBWriter:          dbwriter,
	})

	app.Start()
}
