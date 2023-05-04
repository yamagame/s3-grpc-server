package file

import (
	"sample/s3-grpc-server/entitiy/repository/model"
	"sample/s3-grpc-server/infra/repository"
)

type RepositoryInterface repository.CRUDRepositoryInterface[model.File]
