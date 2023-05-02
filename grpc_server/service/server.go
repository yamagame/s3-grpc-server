package service

import (
	repositoryService "sample/s3-grpc-server/grpc_server/service/repository"
	cellService "sample/s3-grpc-server/grpc_server/service/repository/cell"
	fileinfoService "sample/s3-grpc-server/grpc_server/service/repository/fileinfo"
	tableService "sample/s3-grpc-server/grpc_server/service/repository/table"
	storageService "sample/s3-grpc-server/grpc_server/service/storage"

	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/fileinfo"
	"sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/infra/storage"

	rpc "sample/s3-grpc-server/proto/grpc_server"

	"google.golang.org/grpc"
)

func NewServer(
	storageClient storage.StorageInterface,
	fileinfoRepository fileinfo.RepositoryInterface,
	tableRepository table.RepositoryInterface,
	cellRepository cell.RepositoryInterface,
) *grpc.Server {
	server := grpc.NewServer()
	rpc.RegisterStorageServer(server, storageService.NewStorageServer(storage.NewStorageService(storageClient)))
	rpc.RegisterRepositoryServer(server,
		repositoryService.NewRepositoryServer(
			fileinfoService.NewFileInfoServer(fileinfoRepository),
			tableService.NewTableServer(tableRepository),
			cellService.NewCellServer(cellRepository),
		),
	)
	return server
}
