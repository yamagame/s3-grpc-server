package main

import (
	"fmt"
	"sample/s3-grpc-server/cmd/fxclient/app"
	"sample/s3-grpc-server/cmd/fxclient/constructor"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/client/repository/cell"
	"sample/s3-grpc-server/service/grpc/client/repository/file"
	"sample/s3-grpc-server/service/grpc/client/repository/table"
	"sample/s3-grpc-server/service/grpc/client/storage"
	"sample/s3-grpc-server/share/sheet"

	"go.uber.org/fx"
)

func main() {

	grpcProviders := []interface{}{
		grpc_server.NewFileRepositoryClient,
		grpc_server.NewTableRepositoryClient,
		grpc_server.NewCellRepositoryClient,
		grpc_server.NewStorageRepositoryClient,
	}

	repoRpviders := []interface{}{
		file.NewFileRepository,
		table.NewTableRepository,
		cell.NewCellRepository,
	}

	fx.New(
		fx.Provide(constructor.NewContext),
		fx.Provide(constructor.NewKeyboardScanner),
		fx.Provide(file.NewKeyInput),
		fx.Provide(table.NewKeyInput),
		fx.Provide(cell.NewKeyInput),
		fx.Provide(storage.NewKeyInput),
		fx.Provide(constructor.NewGRPCConnection),
		fx.Provide(constructor.GRPCClientConnection),
		fx.Provide(grpcProviders...),
		fx.Provide(repoRpviders...),
		fx.Provide(storage.NewStorageRepository),
		fx.Provide(sheet.NewCSVSheet),
		fx.Provide(constructor.NewDBWriter),
		fx.Provide(app.NewApp),
		fx.Invoke(func(app *app.App) {
			fmt.Println("start gRPC Client.")
			app.Start()
		}),
	)
}
