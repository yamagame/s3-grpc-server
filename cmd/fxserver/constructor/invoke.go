package constructor

import (
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/server/repository/cell"
	"sample/s3-grpc-server/service/grpc/server/repository/file"
	"sample/s3-grpc-server/service/grpc/server/repository/table"
	"sample/s3-grpc-server/service/grpc/server/storage"

	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RegisterServerIn struct {
	dig.In
	Server         *grpc.Server
	FileService    *file.FileServerRepository
	TableService   *table.TableServerRepository
	CellService    *cell.CellServerRepository
	StorageService *storage.StorageServerRepository
}

func RegisterServer(in RegisterServerIn) {
	grpc_server.RegisterFileRepositoryServer(in.Server, in.FileService)
	grpc_server.RegisterTableRepositoryServer(in.Server, in.TableService)
	grpc_server.RegisterCellRepositoryServer(in.Server, in.CellService)
	grpc_server.RegisterStorageRepositoryServer(in.Server, in.StorageService)
	reflection.Register(in.Server)
}
