package service

import (
	"context"
	"sample/s3-grpc-server/domain"
	aws "sample/s3-grpc-server/grpc-server/proto/grpc-server"
)

type awsClientDomain struct {
	createBucket domain.AwsCreateBucketClientData
	listBuckets  domain.AwsListBucketsClientData
	putObject    domain.AwsPutObjectClientData
	getObject    domain.AwsGetObjectClientData
	deleteObject domain.AwsDeleteObjectClientData
	listobjects  domain.AwsListObjectsClientData
}

type AwsService struct {
	scanner AwsScannerInterface
	domain  awsClientDomain
	client  aws.AwsClient
}

func NewAwsService(aws aws.AwsClient, scanner AwsScannerInterface) *AwsService {
	return &AwsService{
		scanner: scanner,
		domain:  awsClientDomain{},
		client:  aws,
	}
}

func (x *AwsService) CreateBucket(ctx context.Context) (*domain.CreateBucketEntity, error) {
	req := &aws.CreateBucketRequest{}
	res, err := x.client.CreateBucket(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.createBucket.Output(res), nil
}

func (x *AwsService) ListBuckets(ctx context.Context) (*domain.ListBucketsEntity, error) {
	req := &aws.ListBucketsRequest{}
	res, err := x.client.ListBuckets(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.listBuckets.Output(res), nil
}

func (x *AwsService) PutObject(ctx context.Context) (*domain.PutObjectEntity, error) {
	req := x.domain.putObject.Input(x.scanner.PutObject())
	res, err := x.client.PutObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.putObject.Output(res), nil
}

func (x *AwsService) GetObject(ctx context.Context) (*domain.GetObjectEntity, error) {
	req := x.domain.getObject.Input(x.scanner.GetObject())
	res, err := x.client.GetObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.getObject.Output(res), nil
}

func (x *AwsService) DeleteObject(ctx context.Context) (*domain.DeleteObjectEntity, error) {
	req := x.domain.deleteObject.Input(x.scanner.DeleteObject())
	res, err := x.client.DeleteObject(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.deleteObject.Output(res), nil
}

func (x *AwsService) ListObjects(ctx context.Context) (*domain.ListObjectsEntity, error) {
	req := x.domain.listobjects.Input(x.scanner.ListObjects())
	res, err := x.client.ListObjects(ctx, req)
	if err != nil {
		return nil, err
	}
	return x.domain.listobjects.Output(res), nil
}
