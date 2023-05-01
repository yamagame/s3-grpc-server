package storage

import (
	"fmt"
	"sample/s3-grpc-server/infra/storage"
	aws "sample/s3-grpc-server/proto/grpc_server"
)

type storageService struct {
	client storage.StorageInterface
}

func NewStorageService(client storage.StorageInterface) *storageService {
	return &storageService{
		client: client,
	}
}

// CreateBucket implements awsService.ListBuckets
func (s *storageService) CreateBucket(req *storage.CreateBucketEntity) (*storage.CreateBucketEntity, error) {
	err := s.client.CreateBucket()
	if err != nil {
		fmt.Println(err)
		req.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		req.Result = storage.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// ListBuckets implements awsService.ListBuckets
func (s *storageService) ListBuckets(req *storage.ListBucketsEntity) (*storage.ListBucketsEntity, error) {
	buckets, err := s.client.ListBuckets()
	if err != nil {
		fmt.Println(err)
		req.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		req.Result = storage.StorageResult(aws.Result_OK)
		req.Buckets = buckets
	}
	return req, nil
}

// PutObject implements awsService.PutObject
func (s *storageService) PutObject(req *storage.PutObjectEntity) (*storage.PutObjectEntity, error) {
	err := s.client.PutObjectWithString(req.Key, req.Content)
	if err != nil {
		fmt.Println(err)
		req.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		req.Result = storage.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// GetObject implements awsService.GetObject
func (s *storageService) GetObject(req *storage.GetObjectEntity) (*storage.GetObjectEntity, error) {
	ret, err := s.client.GetObjectWithString(req.Key)
	if err != nil {
		fmt.Println(err)
		req.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		req.Result = storage.StorageResult(aws.Result_OK)
		req.Content = ret
	}
	return req, nil
}

// DeleteObject implements awsService.DeleteObject
func (s *storageService) DeleteObject(req *storage.DeleteObjectEntity) (*storage.DeleteObjectEntity, error) {
	err := s.client.DeleteObject(req.Key)
	if err != nil {
		fmt.Println(err)
		req.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		req.Result = storage.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// ListObjects implements awsService.DeleteObject
func (s *storageService) ListObjects(req *storage.ListObjectsEntity) (*storage.ListObjectsEntity, error) {
	objects, err := s.client.ListObjects(req.Prefix, req.Next)
	if err != nil {
		fmt.Println(err)
		req.Result = storage.StorageResult(aws.Result_ERR)
	} else {
		req.Result = storage.StorageResult(aws.Result_OK)
		req.Keys = make([]string, len(objects))
		for i, object := range objects {
			req.Keys[i] = object.Key
		}
	}
	return req, nil
}
