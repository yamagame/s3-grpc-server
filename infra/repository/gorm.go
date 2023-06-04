package repository

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormDB() *gorm.DB {
	dsn := GormDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// db.Logger = db.Logger.LogMode(logger.Info)
	return db.Debug()
}

func GormDSN() string {
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	dsn := "root:" + pass + "@tcp(" + host + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true"
	return dsn
}

func StartTestDatabase(cb func(db *gorm.DB)) {
	cb(GormDB())
}
