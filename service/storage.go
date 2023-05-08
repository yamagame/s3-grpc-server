package service

import (
	"sample/s3-grpc-server/infra"
	"sample/s3-grpc-server/service/grpc/server/storage"
)

func NewStorageServerRepository(repository infra.StorageRepositoryInterface) *storage.StorageServerRepository {
	return storage.NewStorageServerRepository(repository)
}
