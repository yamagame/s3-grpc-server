package main

import (
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository"
)

func main() {
	db := repository.GormDB()
	db.AutoMigrate(&model.File{})
	db.AutoMigrate(&model.Cell{})
	db.AutoMigrate(&model.Table{})
}
