package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"sample/s3-grpc-server/proto/grpc_server"

	"sample/s3-grpc-server/infra/repository"

	cellInfra "sample/s3-grpc-server/infra/repository/cell"
	fileInfra "sample/s3-grpc-server/infra/repository/file"
	tableInfra "sample/s3-grpc-server/infra/repository/table"
	storageInfra "sample/s3-grpc-server/infra/storage"

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
	flag.Parse()
	mode := flag.Arg(0)
	logger, _ := zap.NewProduction()
	defer logger.Sync() //nolint

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
		storageInfra.NewStorageRepository(GetStorageInfraImpl(mode)),
		fileInfra.NewFileRepository(gorm),
		tableInfra.NewTableRepository(gorm),
		cellInfra.NewCellRepository(gorm),
	)
	// log.Printf("server listening at %v", lis.Addr())
	sugar.Infof("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		sugar.Infof("failed to serve: %v", err)
	}
}

func GetStorageInfraImpl(mode string) storageInfra.RepositoryInternalInterface {
	getClient := func(mode string) (storageInfra.RepositoryInternalInterface, error) {
		if mode == "sftp" {
			fmt.Println("start sftp")
			return storageInfra.NewSFTPClient(context.Background(), storageInfra.SFTPClientConfig{
				Host:  os.Getenv("SFTP_HOST"),
				Port:  os.Getenv("SFTP_PORT"),
				User:  os.Getenv("SFTP_USERNAME"),
				Pass:  os.Getenv("SFTP_PASSWORD"),
				Share: os.Getenv("SFTP_SHAREDIR"),
			})
		}
		if mode == "gcs" {
			fmt.Println("start gcs")
			return storageInfra.NewGCSClient(context.Background(), storageInfra.GCSClientConfig{
				Bucket:    os.Getenv("BUCKET_NAME"),
				ProjectID: os.Getenv("GCS_PROJECT_ID"),
			})
		}
		fmt.Println("start s3")
		return storageInfra.NewS3Client(context.Background(), storageInfra.S3ClientConfig{
			Bucket:   os.Getenv("BUCKET_NAME"),
			Endpoint: os.Getenv("S3_ENDPOINT"),
		})
	}
	client, err := getClient(mode)
	if err != nil {
		return nil
	}
	return client
}

func NewServer(
	storageRepository storageInfra.RepositoryInterface,
	fileRepository fileInfra.RepositoryInterface,
	tableRepository tableInfra.RepositoryInterface,
	cellRepository cellInfra.RepositoryInterface,
) *grpc.Server {
	server := grpc.NewServer()
	grpc_server.RegisterStorageRepositoryServer(server,
		storageService.NewStorageServerRepository(storageRepository),
	)
	grpc_server.RegisterFileRepositoryServer(server,
		fileService.NewFileServerRepository(fileRepository),
	)
	grpc_server.RegisterTableRepositoryServer(server,
		tableService.NewTableServerRepository(tableRepository),
	)
	grpc_server.RegisterCellRepositoryServer(server,
		cellService.NewCellServerRepository(cellRepository),
	)
	reflection.Register(server)
	return server
}
