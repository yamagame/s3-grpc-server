package constructor

import (
	"context"
	"os"

	storageInfra "sample/s3-grpc-server/infra/storage"

	"go.uber.org/dig"
	"go.uber.org/zap"
)

func GetStorageInfraImpl(mode string) func(*zap.SugaredLogger) (storageInfra.RepositoryInternalInterface, error) {
	return func(log *zap.SugaredLogger) (storageInfra.RepositoryInternalInterface, error) {
		if mode == "sftp" {
			log.Infof("start sftp")
			return storageInfra.NewSFTPClient(context.Background(), storageInfra.SFTPClientConfig{
				Host:  os.Getenv("SFTP_HOST"),
				Port:  os.Getenv("SFTP_PORT"),
				User:  os.Getenv("SFTP_USERNAME"),
				Pass:  os.Getenv("SFTP_PASSWORD"),
				Share: os.Getenv("SFTP_SHAREDIR"),
			})
		}
		if mode == "gcs" {
			log.Infof("start gcs")
			return storageInfra.NewGCSClient(context.Background(), storageInfra.GCSClientConfig{
				Bucket:    os.Getenv("BUCKET_NAME"),
				ProjectID: os.Getenv("GCS_PROJECT_ID"),
			})
		}
		log.Infof("start s3")
		return storageInfra.NewS3Client(context.Background(), storageInfra.S3ClientConfig{
			Bucket:   os.Getenv("BUCKET_NAME"),
			Endpoint: os.Getenv("S3_ENDPOINT"),
		})
	}
}

type LoggerOut struct {
	dig.Out
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
}

func NewLogger() func() LoggerOut {
	return func() LoggerOut {
		logger := zap.NewExample()
		return LoggerOut{
			Logger: logger,
			Sugar:  logger.Sugar(),
		}
	}
}
