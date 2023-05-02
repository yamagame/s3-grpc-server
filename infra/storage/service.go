package storage

import (
	"fmt"
	"sample/s3-grpc-server/infra/storage/model"
	aws "sample/s3-grpc-server/proto/grpc_server"
)

type StorageService struct {
	client StorageInterface
}

func NewStorageService(client StorageInterface) *StorageService {
	return &StorageService{
		client: client,
	}
}

// CreateBucket implements awsService.ListBuckets
func (s *StorageService) CreateBucket(req *model.CreateBucket) (*model.CreateBucket, error) {
	err := s.client.CreateBucket()
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// ListBuckets implements awsService.ListBuckets
func (s *StorageService) ListBuckets(req *model.ListBuckets) (*model.ListBuckets, error) {
	buckets, err := s.client.ListBuckets()
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
		req.Buckets = buckets
	}
	return req, nil
}

// PutObject implements awsService.PutObject
func (s *StorageService) PutObject(req *model.PutObject) (*model.PutObject, error) {
	err := s.client.PutObjectWithString(req.Key, req.Content)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// GetObject implements awsService.GetObject
func (s *StorageService) GetObject(req *model.GetObject) (*model.GetObject, error) {
	ret, err := s.client.GetObjectWithString(req.Key)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
		req.Content = ret
	}
	return req, nil
}

// DeleteObject implements awsService.DeleteObject
func (s *StorageService) DeleteObject(req *model.DeleteObject) (*model.DeleteObject, error) {
	err := s.client.DeleteObject(req.Key)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// ListObjects implements awsService.DeleteObject
func (s *StorageService) ListObjects(req *model.ListObjects) (*model.ListObjects, error) {
	objects, err := s.client.ListObjects(req.Prefix, req.Next)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
		req.Keys = make([]string, len(objects))
		for i, object := range objects {
			req.Keys[i] = object.Key
		}
	}
	return req, nil
}
