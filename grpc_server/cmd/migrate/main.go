package main

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"
)

func main() {
	db := repository.GormDB()
	db.AutoMigrate(&model.File{})
	db.AutoMigrate(&model.Cell{})
	db.AutoMigrate(&model.Table{})
}
