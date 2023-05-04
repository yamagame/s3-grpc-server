package main

import (
	"context"
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
		GetClient(mode),
		file.NewFileRepository(gorm),
		table.NewTableRepository(gorm),
		cell.NewCellRepository(gorm),
	)
	// log.Printf("server listening at %v", lis.Addr())
	sugar.Infof("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		// log.Fatalf("failed to serve: %v", err)
		sugar.Infof("failed to serve: %v", err)
	}
}

func GetClient(mode string) storage.RepositoryClientInterface {
	getClient := func(mode string) (storage.RepositoryClientInterface, error) {
		if mode == "sftp" {
			fmt.Println("start sftp")
			return storage.NewSFTPClient(context.Background(), storage.SFTPClientConfig{
				Host:  os.Getenv("SFTP_HOST"),
				Port:  os.Getenv("SFTP_PORT"),
				User:  os.Getenv("SFTP_USERNAME"),
				Pass:  os.Getenv("SFTP_PASSWORD"),
				Share: os.Getenv("SFTP_SHAREDIR"),
			})
		}
		if mode == "gcs" {
			fmt.Println("start gcs")
			return storage.NewGCSClient(context.Background(), storage.GCSClientConfig{
				Bucket:    os.Getenv("BUCKET_NAME"),
				ProjectID: os.Getenv("GCS_PROJECT_ID"),
			})
		}
		fmt.Println("start s3")
		return storage.NewS3Client(context.Background(), storage.S3ClientConfig{
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
	storageRepository storage.RepositoryClientInterface,
	fileRepository file.RepositoryInterface,
	tableRepository table.RepositoryInterface,
	cellRepository cell.RepositoryInterface,
) *grpc.Server {
	server := grpc.NewServer()
	grpc_server.RegisterStorageRepositoryServer(server,
		storageService.NewStorageRepositoryServer(storage.NewStorageRepository(storageRepository)),
	)
	grpc_server.RegisterFileRepositoryServer(server,
		fileService.NewFileRepositoryServer(fileRepository),
	)
	grpc_server.RegisterTableRepositoryServer(server,
		tableService.NewTableRepositoryServer(tableRepository),
	)
	grpc_server.RegisterCellRepositoryServer(server,
		cellService.NewCellRepositoryServer(cellRepository),
	)
	reflection.Register(server)
	return server
}
