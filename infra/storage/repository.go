package storage

import (
	"context"
	"fmt"
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type StorageRepository struct {
	client RepositoryInternalInterface
}

func NewStorageRepository(client RepositoryInternalInterface) RepositoryInterface {
	return &StorageRepository{
		client: client,
	}
}

// CreateBucket implements StorageRepository.CreateBucket
func (s *StorageRepository) CreateBucket(ctx context.Context, req *model.CreateBucket) (*model.CreateBucket, error) {
	err := s.client.CreateBucket(ctx, req)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(grpc_server.Result_ERR)
	} else {
		req.Result = model.StorageResult(grpc_server.Result_OK)
	}
	return req, nil
}

// ListBuckets implements StorageRepository.ListBuckets
func (s *StorageRepository) ListBuckets(ctx context.Context, req *model.ListBuckets) (*model.ListBuckets, error) {
	buckets, err := s.client.ListBuckets(ctx, req)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(grpc_server.Result_ERR)
	} else {
		req.Result = model.StorageResult(grpc_server.Result_OK)
		req.Buckets = buckets
	}
	return req, nil
}

// PutObject implements StorageRepository.PutObject
func (s *StorageRepository) PutObject(ctx context.Context, req *model.PutObject) (*model.PutObject, error) {
	err := s.client.PutObjectWithString(ctx, req)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(grpc_server.Result_ERR)
	} else {
		req.Result = model.StorageResult(grpc_server.Result_OK)
	}
	return req, nil
}

// GetObject implements StorageRepository.GetObject
func (s *StorageRepository) GetObject(ctx context.Context, req *model.GetObject) (*model.GetObject, error) {
	ret, err := s.client.GetObjectWithString(ctx, req)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(grpc_server.Result_ERR)
	} else {
		req.Result = model.StorageResult(grpc_server.Result_OK)
		req.Content = ret
	}
	return req, nil
}

// DeleteObject implements StorageRepository.DeleteObject
func (s *StorageRepository) DeleteObject(ctx context.Context, req *model.DeleteObject) (*model.DeleteObject, error) {
	err := s.client.DeleteObject(ctx, req)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(grpc_server.Result_ERR)
	} else {
		req.Result = model.StorageResult(grpc_server.Result_OK)
	}
	return req, nil
}

// ListObjects implements StorageRepository.ListObjects
func (s *StorageRepository) ListObjects(ctx context.Context, req *model.ListObjects) (*model.ListObjects, error) {
	objects, err := s.client.ListObjects(ctx, req)
	fmt.Println(objects)
	if err != nil {
		fmt.Println(err)
		req.Result = model.StorageResult(grpc_server.Result_ERR)
	} else {
		req.Result = model.StorageResult(grpc_server.Result_OK)
		req.Keys = make([]string, len(objects))
		for i, object := range objects {
			req.Keys[i] = object.Key
		}
	}
	return req, nil
}
