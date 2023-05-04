package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"sample/s3-grpc-server/proto/grpc_server"

	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/file"
	"sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/infra/storage"

	cellRepo "sample/s3-grpc-server/infra/repository/cell"
	fileRepo "sample/s3-grpc-server/infra/repository/file"
	tableRepo "sample/s3-grpc-server/infra/repository/table"

	cellService "sample/s3-grpc-server/service/grpc/server/repository/cell"
	fileService "sample/s3-grpc-server/service/grpc/server/repository/file"
	tableService "sample/s3-grpc-server/service/grpc/server/repository/table"
	storageService "sample/s3-grpc-server/service/grpc/server/storage"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	mode := "s3"
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	zap.ReplaceGlobals(logger)

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		// log.Fatalf("failed to listen: %v", err)
		sugar.Infof("failed to listen: %v", err)
	}
	gorm := repository.GormDB()
	server := NewServer(
		storage.GetClient(mode),
		fileRepo.NewFileRepository(gorm),
		tableRepo.NewTableRepository(gorm),
		cellRepo.NewCellRepository(gorm),
	)
	// log.Printf("server listening at %v", lis.Addr())
	sugar.Infof("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		sugar.Infof("failed to serve: %v", err)
	}
}

func NewServer(
	storageClient storage.StorageInterface,
	fileRepository file.RepositoryInterface,
	tableRepository table.RepositoryInterface,
	cellRepository cell.RepositoryInterface,
) *grpc.Server {
	server := grpc.NewServer()
	grpc_server.RegisterStorageRepositoryServer(server,
		storageService.NewStorageServer(storage.NewStorageService(storageClient)),
	)
	grpc_server.RegisterFileRepositoryServer(server,
		fileService.NewFileRepository(fileRepository),
	)
	grpc_server.RegisterTableRepositoryServer(server,
		tableService.NewTableRepository(tableRepository),
	)
	grpc_server.RegisterCellRepositoryServer(server,
		cellService.NewCellRepository(cellRepository),
	)
	reflection.Register(server)
	return server
}
