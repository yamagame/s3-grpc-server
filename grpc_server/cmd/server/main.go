package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	repositoryService "sample/s3-grpc-server/grpc_server/service/repository"
	storageService "sample/s3-grpc-server/grpc_server/service/storage"
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/storage"
	aws "sample/s3-grpc-server/proto/grpc_server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
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
	s := grpc.NewServer()
	aws.RegisterStorageServer(s, storageService.NewStorageServer(storage.NewStorageService(storage.GetClient(mode))))
	aws.RegisterRepositoryServer(s, repositoryService.NewRepositoryServer(repository.NewRepositoryService(repository.GormDB())))
	// log.Printf("server listening at %v", lis.Addr())
	sugar.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		sugar.Infof("failed to serve: %v", err)
	}
}
