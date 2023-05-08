package infra

import (
	"sample/s3-grpc-server/infra/storage"
)

type StorageRepository storage.StorageRepository
type StorageRepositoryInterface storage.RepositoryInterface
type StorageRepositoryClientInterface storage.RepositoryClientInterface

func NewStorageRepository(client storage.RepositoryClientInterface) StorageRepositoryInterface {
	return storage.NewStorageRepository(client)
}
