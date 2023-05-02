package file

import (
	"sample/s3-grpc-server/infra/repository/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

func ToDomain(file *server.File) *model.File {
	return &model.File{
		ID:       file.Id,
		Filename: file.Filename,
	}
}

func ToInfra(file *model.File) *server.File {
	return &server.File{
		Id:       file.ID,
		Filename: file.Filename,
	}
}
