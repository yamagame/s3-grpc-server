package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sample/s3-grpc-server/infra/repository"
	cellRepo "sample/s3-grpc-server/infra/repository/cell"
	fileRepo "sample/s3-grpc-server/infra/repository/file"
	tableRepo "sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/infra/storage"
	grpc "sample/s3-grpc-server/service/grpc/server"

	"go.uber.org/zap"
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
	server := grpc.NewServer(
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
