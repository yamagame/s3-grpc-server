package constructor

import (
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/server/repository/cell"
	"sample/s3-grpc-server/service/grpc/server/repository/file"
	"sample/s3-grpc-server/service/grpc/server/repository/table"
	"sample/s3-grpc-server/service/grpc/server/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RegisterFileServer(server *grpc.Server, fileService *file.FileServerRepository) {
	grpc_server.RegisterFileRepositoryServer(server, fileService)
}

func RegisterTableServer(server *grpc.Server, tableService *table.TableServerRepository) {
	grpc_server.RegisterTableRepositoryServer(server, tableService)
}

func RegisterCellServer(server *grpc.Server, cellService *cell.CellServerRepository) {
	grpc_server.RegisterCellRepositoryServer(server, cellService)
}

func RegisterStorageServer(server *grpc.Server, storageService *storage.StorageServerRepository) {
	grpc_server.RegisterStorageRepositoryServer(server, storageService)
}

func RegisterReflection(server *grpc.Server) {
	reflection.Register(server)
}
