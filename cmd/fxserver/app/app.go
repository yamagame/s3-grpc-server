package app

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sample/s3-grpc-server/service/grpc/server/repository/cell"
	"sample/s3-grpc-server/service/grpc/server/repository/file"
	"sample/s3-grpc-server/service/grpc/server/repository/table"
	"sample/s3-grpc-server/service/grpc/server/storage"

	"go.uber.org/dig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AppIn struct {
	dig.In
	Server         *grpc.Server
	FileService    *file.FileServerRepository
	TableService   *table.TableServerRepository
	CellService    *cell.CellServerRepository
	StorageService *storage.StorageServerRepository
	Logger         *zap.SugaredLogger
}

type App struct {
	server         *grpc.Server
	fileService    *file.FileServerRepository
	tableService   *table.TableServerRepository
	cellService    *cell.CellServerRepository
	storageService *storage.StorageServerRepository
	logger         *zap.SugaredLogger
}

func NewApp(in AppIn) *App {
	return &App{
		server:         in.Server,
		fileService:    in.FileService,
		tableService:   in.TableService,
		cellService:    in.CellService,
		storageService: in.StorageService,
		logger:         in.Logger,
	}
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func (x *App) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		// log.Fatalf("failed to listen: %v", err)
		x.logger.Infof("failed to listen: %v", err)
	}

	// log.Printf("server listening at %v", lis.Addr())
	x.logger.Infof("server listening at %v", lis.Addr())
	if err := x.server.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		x.logger.Infof("failed to serve: %v", err)
	}
}

func (x *App) Shutdown(ctx context.Context) error {
	x.logger.Sync()
	return nil
}
