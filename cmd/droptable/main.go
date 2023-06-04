package main

import (
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository"
)

func main() {
	db := repository.GormDB()
	db.Migrator().DropTable(&model.File{})
	db.Migrator().DropTable(&model.Cell{})
	db.Migrator().DropTable(&model.Table{})
	db.Migrator().DropTable(&model.Attr{})
}
