package infra

import (
	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/file"
	"sample/s3-grpc-server/infra/repository/table"

	"gorm.io/gorm"
)

type CellRepositoryInterface cell.RepositoryInterface
type FileRepositoryInterface file.RepositoryInterface
type TableRepositoryInterface table.RepositoryInterface

func NewCellRepository(db *gorm.DB) CellRepositoryInterface {
	return cell.NewCellRepository(db)
}

func NewFileRepository(db *gorm.DB) FileRepositoryInterface {
	return file.NewFileRepository(db)
}

func NewTableRepository(db *gorm.DB) TableRepositoryInterface {
	return table.NewTableRepository(db)
}
