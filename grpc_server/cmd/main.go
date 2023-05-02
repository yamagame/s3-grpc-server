package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sample/s3-grpc-server/grpc_server/service"
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/fileinfo"
	"sample/s3-grpc-server/infra/storage"

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
	server := service.NewServer(storage.GetClient(mode), fileinfo.NewRepository(repository.GormDB()))
	// log.Printf("server listening at %v", lis.Addr())
	sugar.Infof("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		sugar.Infof("failed to serve: %v", err)
	}
}
