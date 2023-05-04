package file

import (
	"sample/s3-grpc-server/infra/repository/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

func ToDomain(file *grpc_server.File) *model.File {
	return &model.File{
		ID:       file.Id,
		Filename: file.Filename,
	}
}

func ToInfra(file *model.File) *grpc_server.File {
	return &grpc_server.File{
		Id:       file.ID,
		Filename: file.Filename,
	}
}
