package storage

import (
	"context"
	"fmt"
	"sample/s3-grpc-server/infra/storage/model"
	aws "sample/s3-grpc-server/proto/grpc_server"
)

type StorageServiceInterface interface {
	CreateBucket(ctx context.Context, req *model.CreateBucket) (*model.CreateBucket, error)
	ListBuckets(ctx context.Context, req *model.ListBuckets) (*model.ListBuckets, error)
	PutObject(ctx context.Context, req *model.PutObject) (*model.PutObject, error)
	GetObject(ctx context.Context, req *model.GetObject) (*model.GetObject, error)
	DeleteObject(ctx context.Context, req *model.DeleteObject) (*model.DeleteObject, error)
	ListObjects(ctx context.Context, req *model.ListObjects) (*model.ListObjects, error)
}

type StorageService struct {
	client StorageInterface
}

func NewStorageService(client StorageInterface) *StorageService {
	return &StorageService{
		client: client,
	}
}

// CreateBucket implements awsService.ListBuckets
func (s *StorageService) CreateBucket(ctx context.Context, req *model.CreateBucket) (*model.CreateBucket, error) {
	err := s.client.CreateBucket(ctx)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// ListBuckets implements awsService.ListBuckets
func (s *StorageService) ListBuckets(ctx context.Context, req *model.ListBuckets) (*model.ListBuckets, error) {
	buckets, err := s.client.ListBuckets(ctx)
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
func (s *StorageService) PutObject(ctx context.Context, req *model.PutObject) (*model.PutObject, error) {
	err := s.client.PutObjectWithString(ctx, req.Key, req.Content)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// GetObject implements awsService.GetObject
func (s *StorageService) GetObject(ctx context.Context, req *model.GetObject) (*model.GetObject, error) {
	ret, err := s.client.GetObjectWithString(ctx, req.Key)
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
func (s *StorageService) DeleteObject(ctx context.Context, req *model.DeleteObject) (*model.DeleteObject, error) {
	err := s.client.DeleteObject(ctx, req.Key)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(aws.Result_ERR)
	} else {
		req.Result = model.StorageResult(aws.Result_OK)
	}
	return req, nil
}

// ListObjects implements awsService.DeleteObject
func (s *StorageService) ListObjects(ctx context.Context, req *model.ListObjects) (*model.ListObjects, error) {
	objects, err := s.client.ListObjects(ctx, req.Prefix, req.Next)
	fmt.Println(objects)
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
