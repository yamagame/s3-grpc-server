package storage

import (
	"fmt"
	"sample/s3-grpc-server/domain"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type awsService struct {
	aws *Client
}

func NewAWSService(client *Client) *awsService {
	return &awsService{
		aws: client,
	}
}

// CreateBucket implements awsService.ListBuckets
func (s *awsService) CreateBucket(entity *domain.CreateBucketEntity) (*domain.CreateBucketEntity, error) {
	_, err := s.aws.CreateBucket()
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
	}
	return entity, nil
}

// ListBuckets implements awsService.ListBuckets
func (s *awsService) ListBuckets(entity *domain.ListBucketsEntity) (*domain.ListBucketsEntity, error) {
	list, err := s.aws.ListBuckets()
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
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
func (s *awsService) PutObject(entity *domain.PutObjectEntity) (*domain.PutObjectEntity, error) {
	_, err := s.aws.PutObjectWithString(entity.Key, entity.Content)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
	}
	return entity, nil
}

// GetObject implements awsService.GetObject
func (s *awsService) GetObject(entity *domain.GetObjectEntity) (*domain.GetObjectEntity, error) {
	ret, err := s.aws.GetObjectWithString(entity.Key)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
		entity.Content = ret
	}
	return entity, nil
}

// DeleteObject implements awsService.DeleteObject
func (s *awsService) DeleteObject(entity *domain.DeleteObjectEntity) (*domain.DeleteObjectEntity, error) {
	_, err := s.aws.DeleteObject(entity.Key)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
	}
	return entity, nil
}

// ListObjects implements awsService.DeleteObject
func (s *awsService) ListObjects(entity *domain.ListObjectsEntity) (*domain.ListObjectsEntity, error) {
	res, err := s.aws.ListObjects(entity.Prefix, entity.Next)
	if err != nil {
		fmt.Println(err)
		entity.Result = domain.AwsResult(aws.Result_ERR)
	} else {
		entity.Result = domain.AwsResult(aws.Result_OK)
		entity.Keys = make([]string, res.KeyCount)
		for i, cont := range res.Contents {
			entity.Keys[i] = *cont.Key
		}
	}
	return entity, nil
}
