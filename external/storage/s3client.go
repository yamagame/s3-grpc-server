package storage

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	ctx      context.Context
	bucket   string
	s3client *s3.Client
}

func NewS3Client(ctx context.Context, bucket string) (*S3Client, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID {
			url := os.Getenv("S3_ENDPOINT")
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
		bucket:   bucket,
	}, nil
}

func (x *S3Client) ListBuckets() ([]Bucket, error) {
	var buckets []Bucket
	res, err := x.s3client.ListBuckets(x.ctx, &s3.ListBucketsInput{})
	if err != nil {
		return nil, err
	}
	for _, bucket := range res.Buckets {
		buckets = append(buckets, Bucket{
			Name: *bucket.Name,
		})
	}
	return buckets, nil
}

func (x *S3Client) CreateBucket() error {
	_, err := x.s3client.CreateBucket(x.ctx, &s3.CreateBucketInput{
		Bucket: aws.String(x.bucket),
	})
	return err
}

func (x *S3Client) PutObject(key string, body io.Reader) error {
	_, err := x.s3client.PutObject(x.ctx, &s3.PutObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	return err
}

func (x *S3Client) PutObjectWithString(key, content string) error {
	return x.PutObject(key, strings.NewReader(content))
}

func (x *S3Client) GetObject(key string) (io.Reader, error) {
	res, err := x.s3client.GetObject(x.ctx, &s3.GetObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (x *S3Client) GetObjectWithString(key string) (string, error) {
	reader, err := x.GetObject(key)
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

func (x *S3Client) DeleteObject(key string) error {
	_, err := x.s3client.DeleteObject(x.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(key),
	})
	return err
}

func (x *S3Client) ListObjects(prefix string, nexttoken *string) ([]Object, error) {
	res, err := x.s3client.ListObjectsV2(x.ctx, &s3.ListObjectsV2Input{
		Bucket:            aws.String(x.bucket),
		Prefix:            aws.String(prefix),
		ContinuationToken: nexttoken,
	})
	if err != nil {
		return nil, err
	}
	objects := make([]Object, len(res.Contents))
	for i, cont := range res.Contents {
		objects[i] = Object{
			Key: *cont.Key,
		}
	}
	return objects, nil
}
