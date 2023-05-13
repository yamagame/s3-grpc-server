package infra

import (
	"sample/s3-grpc-server/infra/storage"
)

type StorageRepository storage.StorageRepository
type StorageRepositoryInterface storage.RepositoryInterface
type StorageRepositoryClientInterface storage.RepositoryInternalInterface

func NewStorageRepository(client storage.RepositoryInternalInterface) StorageRepositoryInterface {
	return storage.NewStorageRepository(client)
}
