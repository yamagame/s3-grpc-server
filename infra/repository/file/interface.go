package file

import (
	"sample/s3-grpc-server/infra/repository"
	"sample/s3-grpc-server/infra/repository/model"
)

type RepositoryInterface repository.CRUDRepositoryInterface[model.File]
