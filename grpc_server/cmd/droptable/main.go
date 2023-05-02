package main

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"
)

func main() {
	db := repository.GormDB()
	db.Migrator().DropTable(&model.File{})
	db.Migrator().DropTable(&model.Cell{})
	db.Migrator().DropTable(&model.Table{})
}
