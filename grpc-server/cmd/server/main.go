package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"sample/s3-grpc-server/external/storage"
	service "sample/s3-grpc-server/grpc-server/service/storage"
	aws "sample/s3-grpc-server/proto/grpc-server"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
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
	bucket := os.Getenv("BUCKET_NAME")
	s := grpc.NewServer()
	client, err := storage.NewGCSClient(context.Background(), bucket)
	// client, err := domain.NewS3Client(context.Background(), bucket)
	if err != nil {
		// log.Fatalf("failed to listen: %v", err)
		sugar.Infof("failed to listen: %v", err)
	}
	aws.RegisterStorageServer(s, service.NewStorageServer(client))
	// log.Printf("server listening at %v", lis.Addr())
	sugar.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		sugar.Infof("failed to serve: %v", err)
	}
}
