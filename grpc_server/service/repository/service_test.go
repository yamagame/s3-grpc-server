package repository

import (
	"fmt"
	"sample/s3-grpc-server/infra/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestRepositoryService(t *testing.T) {
	repository.StartTestDatabase(func(gormDB *gorm.DB) {
		sv := repository.NewRepositoryService(gormDB)
		gormDB.AutoMigrate(&repository.FileInfo{})
		file := &repository.FileInfo{
			Filename: "test.txt",
		}
		file, err := sv.CreateFileInfo(file)
		fmt.Println(file)
		assert.NoError(t, err)
		file, err = sv.ReadFileInfo(file)
		assert.NoError(t, err)
		file.Filename = "hello.txt"
		file, err = sv.UpdateFileInfo(file)
		assert.NoError(t, err)
		file, err = sv.DeleteFileInfo(file)
		assert.NoError(t, err)
		fmt.Println(file)
	})
}
