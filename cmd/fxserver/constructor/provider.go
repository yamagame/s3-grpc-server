package constructor

import (
	"context"
	"fmt"
	"os"

	cellInfra "sample/s3-grpc-server/infra/repository/cell"
	fileInfra "sample/s3-grpc-server/infra/repository/file"
	tableInfra "sample/s3-grpc-server/infra/repository/table"
	storageInfra "sample/s3-grpc-server/infra/storage"

	cellService "sample/s3-grpc-server/service/grpc/server/repository/cell"
	fileService "sample/s3-grpc-server/service/grpc/server/repository/file"
	tableService "sample/s3-grpc-server/service/grpc/server/repository/table"
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

func NewFileService(repo *fileInfra.FileRepository) *fileService.FileServerRepository {
	return fileService.NewFileServerRepository(repo)
}

func NewTableService(repo *tableInfra.TableRepository) *tableService.TableServerRepository {
	return tableService.NewTableServerRepository(repo)
}

func NewCellService(repo *cellInfra.CellRepository) *cellService.CellServerRepository {
	return cellService.NewCellServerRepository(repo)
}
