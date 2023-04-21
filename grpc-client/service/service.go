package service

import (
	"context"
	"sample/s3-grpc-server/domain"
	client "sample/s3-grpc-server/domain/client"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type AwsService struct {
	scanner AwsScannerInterface
	input   domain.AwsClientInputInterface
	output  domain.AwsClientOutputInterface
	client  aws.AwsClient
}

func NewAwsService(aws aws.AwsClient, scanner AwsScannerInterface) *AwsService {
	return &AwsService{
		scanner: scanner,
		input:   client.NewAwsClientInput(),
		output:  client.NewAwsClientOutput(),
		client:  aws,
	}
}

func (x *AwsService) CreateBucket(ctx context.Context) (*domain.CreateBucketEntity, error) {
	req := &aws.CreateBucketRequest{}
	res, err := x.client.CreateBucket(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.output.CreateBucket(res), nil
}

func (x *AwsService) ListBuckets(ctx context.Context) (*domain.ListBucketsEntity, error) {
	req := &aws.ListBucketsRequest{}
	res, err := x.client.ListBuckets(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.output.ListBuckets(res), nil
}

func (x *AwsService) PutObject(ctx context.Context) (*domain.PutObjectEntity, error) {
	req := x.input.PutObject(x.scanner.PutObject())
	res, err := x.client.PutObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.output.PutObject(res), nil
}

func (x *AwsService) GetObject(ctx context.Context) (*domain.GetObjectEntity, error) {
	req := x.input.GetObject(x.scanner.GetObject())
	res, err := x.client.GetObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.output.GetObject(res), nil
}

func (x *AwsService) DeleteObject(ctx context.Context) (*domain.DeleteObjectEntity, error) {
	req := x.input.DeleteObject(x.scanner.DeleteObject())
	res, err := x.client.DeleteObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.output.DeleteObject(res), nil
}
