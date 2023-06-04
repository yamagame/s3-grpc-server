package main

import (
	"database/sql"
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository"

	_ "github.com/go-sql-driver/mysql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db := repository.GormDB()
	db.AutoMigrate(&model.File{})
	db.AutoMigrate(&model.Cell{})
	db.AutoMigrate(&model.Table{})
	db.AutoMigrate(&model.Attr{})
}

func _main() {
	db, _ := sql.Open("mysql", repository.GormDSN())
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./cmd/migrate/migrations",
		"mysql", driver)
	if err != nil {
		panic(err)
	}

	// err = m.Steps(-1)
	// err = m.Down()
	// err = m.Up()
	// err = m.Drop()
	err = m.Force(20230522003)
	if err != nil {
		panic(err)
	}
}
