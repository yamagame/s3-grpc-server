package storage

import (
	"fmt"
	"sample/s3-grpc-server/domain"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type storageService struct {
	client ClientInterface
}

func NewStorageService(client ClientInterface) *storageService {
	return &storageService{
		client: client,
	}
}

// CreateBucket implements awsService.ListBuckets
func (s *storageService) CreateBucket(entity *domain.CreateBucketEntity) (*domain.CreateBucketEntity, error) {
	_, err := s.client.CreateBucket()
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = domain.StorageResult(aws.Result_OK)
	}
	return entity, nil
}

// ListBuckets implements awsService.ListBuckets
func (s *storageService) ListBuckets(entity *domain.ListBucketsEntity) (*domain.ListBucketsEntity, error) {
	list, err := s.client.ListBuckets()
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = domain.StorageResult(aws.Result_OK)
		entity.Buckets = make([]domain.Bucket, len(list.Buckets))
		for i, bucket := range list.Buckets {
			entity.Buckets[i] = domain.Bucket{
				Name: *bucket.Name,
			}
		}
	}
	return entity, nil
}

// PutObject implements awsService.PutObject
func (s *storageService) PutObject(entity *domain.PutObjectEntity) (*domain.PutObjectEntity, error) {
	_, err := s.client.PutObjectWithString(entity.Key, entity.Content)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = domain.StorageResult(aws.Result_OK)
	}
	return entity, nil
}

// GetObject implements awsService.GetObject
func (s *storageService) GetObject(entity *domain.GetObjectEntity) (*domain.GetObjectEntity, error) {
	ret, err := s.client.GetObjectWithString(entity.Key)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = domain.StorageResult(aws.Result_OK)
		entity.Content = ret
	}
	return entity, nil
}

// DeleteObject implements awsService.DeleteObject
func (s *storageService) DeleteObject(entity *domain.DeleteObjectEntity) (*domain.DeleteObjectEntity, error) {
	_, err := s.client.DeleteObject(entity.Key)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = domain.StorageResult(aws.Result_OK)
	}
	return entity, nil
}

// ListObjects implements awsService.DeleteObject
func (s *storageService) ListObjects(entity *domain.ListObjectsEntity) (*domain.ListObjectsEntity, error) {
	res, err := s.client.ListObjects(entity.Prefix, entity.Next)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.StorageResult(aws.Result_ERR)
	} else {
		entity.Result = domain.StorageResult(aws.Result_OK)
		entity.Keys = make([]string, res.KeyCount)
		for i, cont := range res.Contents {
			entity.Keys[i] = *cont.Key
		}
	}
	return entity, nil
}
