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

	storageService "sample/s3-grpc-server/service/grpc/server/storage"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	mode := flag.Arg(0)

	grpcProviders := []interface{}{
		constructor.RegisterFileServer,
		constructor.RegisterTableServer,
		constructor.RegisterCellServer,
		constructor.RegisterStorageServer,
		constructor.RegisterReflection,
	}

	serviceProviders := []interface{}{
		constructor.NewFileService,
		constructor.NewTableService,
		constructor.NewCellService,
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
		fx.Invoke(grpcProviders...),
		fx.Invoke(func(app *app.App) {
			app.Start()
		}),
	)
}
