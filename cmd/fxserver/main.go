package main

import (
	"flag"
	"sample/s3-grpc-server/cmd/fxserver/app"
	"sample/s3-grpc-server/cmd/fxserver/constructor"

	"sample/s3-grpc-server/infra/repository"

	cellInfra "sample/s3-grpc-server/infra/repository/cell"
	fileInfra "sample/s3-grpc-server/infra/repository/file"
	tableInfra "sample/s3-grpc-server/infra/repository/table"
	storageInfra "sample/s3-grpc-server/infra/storage"

	cellService "sample/s3-grpc-server/service/grpc/server/repository/cell"
	fileService "sample/s3-grpc-server/service/grpc/server/repository/file"
	tableService "sample/s3-grpc-server/service/grpc/server/repository/table"
	storageService "sample/s3-grpc-server/service/grpc/server/storage"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	mode := flag.Arg(0)

	serviceProviders := []interface{}{
		fileService.NewFileServerRepository,
		tableService.NewTableServerRepository,
		cellService.NewCellServerRepository,
		storageService.NewStorageServerRepository,
	}

	infraProviders := []interface{}{
		fileInfra.NewFileRepository,
		tableInfra.NewTableRepository,
		cellInfra.NewCellRepository,
		storageInfra.NewStorageRepository,
	}

	fx.New(
		fx.Provide(repository.GormDB),
		fx.Provide(constructor.GetStorageInfraImpl(mode)),
		fx.Provide(infraProviders...),
		fx.Provide(grpc.NewServer),
		fx.Provide(serviceProviders...),
		fx.Provide(app.NewApp),
		fx.Invoke(constructor.RegisterServer),
		fx.Invoke(func(app *app.App) {
			app.Start()
		}),
	)
}
