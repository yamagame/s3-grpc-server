package storage

import (
	"bytes"
	"context"
	"io"
	"sample/s3-grpc-server/entitiy/storage/model"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3ClientConfig struct {
	Bucket   string
	Endpoint string
}

type S3Client struct {
	ctx      context.Context
	bucket   string
	s3client *s3.Client
}

func NewS3Client(ctx context.Context, options S3ClientConfig) (*S3Client, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, opt ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && options.Endpoint != "" {
			url := options.Endpoint
			if url != "" {
				return aws.Endpoint{URL: url, HostnameImmutable: true}, nil
			}
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(resolver),
	)
	if err != nil {
		return nil, err
	}

	s3client := s3.NewFromConfig(cfg)

	return &S3Client{
		ctx:      ctx,
		s3client: s3client,
		bucket:   options.Bucket,
	}, nil
}

// Close implements S3Client.Close
func (x *S3Client) Close() error {
	return nil
}

func (x *S3Client) CreateBucket(ctx context.Context, _ *model.CreateBucket) error {
	_, err := x.s3client.CreateBucket(x.ctx, &s3.CreateBucketInput{
		Bucket: aws.String(x.bucket),
	})
	return err
}

func (x *S3Client) ListBuckets(ctx context.Context, _ *model.ListBuckets) ([]model.Bucket, error) {
	var buckets []model.Bucket
	res, err := x.s3client.ListBuckets(x.ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	for _, bucket := range res.Buckets {
		buckets = append(buckets, model.Bucket{
			Name: *bucket.Name,
		})
	}
	return buckets, nil
}

func (x *S3Client) PutObject(ctx context.Context, req *model.PutObject, body io.Reader) error {
	_, err := x.s3client.PutObject(x.ctx, &s3.PutObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(req.Key),
		Body:   body,
	})
	return err
}

func (x *S3Client) PutObjectWithString(ctx context.Context, req *model.PutObject) error {
	return x.PutObject(ctx, req, strings.NewReader(req.Content))
}

func (x *S3Client) GetObject(ctx context.Context, req *model.GetObject) (io.Reader, error) {
	res, err := x.s3client.GetObject(x.ctx, &s3.GetObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(req.Key),
	})
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (x *S3Client) GetObjectWithString(ctx context.Context, req *model.GetObject) (string, error) {
	reader, err := x.GetObject(ctx, req)
	if err != nil {
		return "", err
	}

	streamToString := func(stream io.Reader) string {
		buf := new(bytes.Buffer)
		buf.ReadFrom(stream)
		return buf.String()
	}

	return streamToString(reader), nil
}

func (x *S3Client) DeleteObject(ctx context.Context, req *model.DeleteObject) error {
	_, err := x.s3client.DeleteObject(x.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(req.Key),
	})
	return err
}

func (x *S3Client) ListObjects(ctx context.Context, req *model.ListObjects) ([]model.Object, error) {
	res, err := x.s3client.ListObjectsV2(x.ctx, &s3.ListObjectsV2Input{
		Bucket:            aws.String(x.bucket),
		Prefix:            aws.String(req.Prefix),
		ContinuationToken: req.Next,
	})
	if err != nil {
		return nil, err
	}
	objects := make([]model.Object, len(res.Contents))
	for i, cont := range res.Contents {
		objects[i] = model.Object{
			Key: *cont.Key,
		}
	}
	return objects, nil
}
