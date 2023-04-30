package storage

import (
	"context"
	"fmt"
	"os"
)

func GetClient(mode string) ClientInterface {
	getClient := func(mode string) (ClientInterface, error) {
		if mode == "sftp" {
			fmt.Println("start sftp")
			return NewSFTPClient(context.Background(), SFTPClientConfig{
				Host:  os.Getenv("SFTP_HOST"),
				Port:  os.Getenv("SFTP_PORT"),
				User:  os.Getenv("SFTP_USERNAME"),
				Pass:  os.Getenv("SFTP_PASSWORD"),
				Share: os.Getenv("SFTP_SHAREDIR"),
			})
		}
		if mode == "gcs" {
			fmt.Println("start gcs")
			return NewGCSClient(context.Background(), GCSClientConfig{
				Bucket:    os.Getenv("BUCKET_NAME"),
				ProjectID: os.Getenv("GCS_PROJECT_ID"),
			})
		}
		fmt.Println("start s3")
		return NewS3Client(context.Background(), S3ClientConfig{
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
