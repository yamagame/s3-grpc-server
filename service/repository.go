package service

import (
	"sample/s3-grpc-server/infra"
	"sample/s3-grpc-server/service/grpc/server/repository/cell"
	"sample/s3-grpc-server/service/grpc/server/repository/file"
	"sample/s3-grpc-server/service/grpc/server/repository/table"
)

func NewCellServerRepository(repository infra.CellRepositoryInterface) *cell.CellServerRepository {
	return cell.NewCellServerRepository(repository)
}

func NewFileServerRepository(repository infra.FileRepositoryInterface) *file.FileServerRepository {
	return file.NewFileServerRepository(repository)
}

func NewTableServerRepository(repository infra.TableRepositoryInterface) *table.TableServerRepository {
	return table.NewTableServerRepository(repository)
}
