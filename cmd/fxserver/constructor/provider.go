package constructor

import (
	"context"
	"fmt"
	"os"

	storageInfra "sample/s3-grpc-server/infra/storage"
)

func GetStorageInfraImpl(mode string) func() (storageInfra.RepositoryClientInterface, error) {
	return func() (storageInfra.RepositoryClientInterface, error) {
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
}
