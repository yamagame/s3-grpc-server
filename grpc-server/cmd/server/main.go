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
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func db() *gorm.DB {
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	dsn := "root:" + pass + "@tcp(" + host + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	return db
}

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
	getClient := func(mode string) (storage.ClientInterface, error) {
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
