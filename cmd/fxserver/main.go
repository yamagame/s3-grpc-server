package main

import (
	"context"
	"flag"
	"sample/s3-grpc-server/cmd/fxserver/app"
	"sample/s3-grpc-server/cmd/fxserver/constructor"
	"sample/s3-grpc-server/infra"
	"sample/s3-grpc-server/service"
	"time"

	"sample/s3-grpc-server/infra/repository"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	mode := flag.Arg(0)

	serviceProviders := []interface{}{
		service.NewFileServerRepository,
		service.NewTableServerRepository,
		service.NewCellServerRepository,
		service.NewStorageServerRepository,
	}

	infraProviders := []interface{}{
		infra.NewFileRepository,
		infra.NewTableRepository,
		infra.NewCellRepository,
		infra.NewStorageRepository,
	}

	fx.New(
		fx.Provide(repository.GormDB),
		fx.Provide(constructor.GetStorageInfraImpl(mode)),
		fx.Provide(infraProviders...),
		fx.Provide(grpc.NewServer),
		fx.Provide(serviceProviders...),
		fx.Provide(constructor.NewLogger()),
		fx.Provide(app.NewApp),
		fx.Invoke(constructor.RegisterServer),
		fx.Invoke(func(app *app.App, lc fx.Lifecycle, log *zap.Logger) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					log.Info("start app", zap.String("key", "value"), zap.Time("now", time.Now()))
					go app.Start()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return app.Shutdown(ctx)
				},
			})
		}),
	).Run()
}
