package service

import (
	"sample/s3-grpc-server/infra"
	"sample/s3-grpc-server/service/grpc/interceptor"
	"sample/s3-grpc-server/service/grpc/server/repository/cell"
	"sample/s3-grpc-server/service/grpc/server/repository/file"
	"sample/s3-grpc-server/service/grpc/server/repository/table"
)

func NewCellServerRepository(repository infra.CellRepositoryInterface, validator *interceptor.Validator) *cell.CellServerRepository {
	return cell.NewCellServerRepository(repository, validator)
}

func NewFileServerRepository(repository infra.FileRepositoryInterface, validator *interceptor.Validator) *file.FileServerRepository {
	return file.NewFileServerRepository(repository, validator)
}

func NewTableServerRepository(repository infra.TableRepositoryInterface, validator *interceptor.Validator) *table.TableServerRepository {
	return table.NewTableServerRepository(repository, validator)
}
