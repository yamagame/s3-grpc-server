package service

import (
	repositoryService "sample/s3-grpc-server/grpc_server/service/repository"
	fileinfoService "sample/s3-grpc-server/grpc_server/service/repository/fileinfo"
	storageService "sample/s3-grpc-server/grpc_server/service/storage"

	"sample/s3-grpc-server/infra/repository/fileinfo"
	"sample/s3-grpc-server/infra/storage"

	rpc "sample/s3-grpc-server/proto/grpc_server"

	"google.golang.org/grpc"
)

func NewServer(storageClient storage.StorageInterface, fileinfoRepository fileinfo.RepositoryInterface) *grpc.Server {
	server := grpc.NewServer()
	rpc.RegisterStorageServer(server, storageService.NewStorageServer(storage.NewStorageService(storageClient)))
	rpc.RegisterRepositoryServer(server, repositoryService.NewRepositoryServer(
		fileinfoService.NewFileInfoServer(fileinfoRepository),
	))
	return server
}
