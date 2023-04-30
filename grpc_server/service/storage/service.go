package storage

import (
	"fmt"
	"sample/s3-grpc-server/infra/storage"
	aws "sample/s3-grpc-server/proto/grpc_server"
)

type storageService struct {
	client storage.ClientInterface
}

func NewStorageService(client storage.ClientInterface) *storageService {
	return &storageService{
		client: client,
	}
}

// CreateBucket implements awsService.ListBuckets
func (s *storageService) CreateBucket(entity *storage.CreateBucketEntity) (*storage.CreateBucketEntity, error) {
	err := s.client.CreateBucket()
	if err != nil {
		fmt.Println(err)
		entity.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = storage.StorageResult(aws.Result_OK)
	}
	return entity, nil
}

// ListBuckets implements awsService.ListBuckets
func (s *storageService) ListBuckets(entity *storage.ListBucketsEntity) (*storage.ListBucketsEntity, error) {
	buckets, err := s.client.ListBuckets()
	if err != nil {
		fmt.Println(err)
		entity.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = storage.StorageResult(aws.Result_OK)
		entity.Buckets = buckets
	}
	return entity, nil
}

// PutObject implements awsService.PutObject
func (s *storageService) PutObject(entity *storage.PutObjectEntity) (*storage.PutObjectEntity, error) {
	err := s.client.PutObjectWithString(entity.Key, entity.Content)
	if err != nil {
		fmt.Println(err)
		entity.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = storage.StorageResult(aws.Result_OK)
	}
	return entity, nil
}

// GetObject implements awsService.GetObject
func (s *storageService) GetObject(entity *storage.GetObjectEntity) (*storage.GetObjectEntity, error) {
	ret, err := s.client.GetObjectWithString(entity.Key)
	if err != nil {
		fmt.Println(err)
		entity.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = storage.StorageResult(aws.Result_OK)
		entity.Content = ret
	}
	return entity, nil
}

// DeleteObject implements awsService.DeleteObject
func (s *storageService) DeleteObject(entity *storage.DeleteObjectEntity) (*storage.DeleteObjectEntity, error) {
	err := s.client.DeleteObject(entity.Key)
	if err != nil {
		fmt.Println(err)
		entity.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = storage.StorageResult(aws.Result_OK)
	}
	return entity, nil
}

// ListObjects implements awsService.DeleteObject
func (s *storageService) ListObjects(entity *storage.ListObjectsEntity) (*storage.ListObjectsEntity, error) {
	objects, err := s.client.ListObjects(entity.Prefix, entity.Next)
	if err != nil {
		fmt.Println(err)
		entity.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = storage.StorageResult(aws.Result_OK)
		entity.Keys = make([]string, len(objects))
		for i, object := range objects {
			entity.Keys[i] = object.Key
		}
	}
	return entity, nil
}
