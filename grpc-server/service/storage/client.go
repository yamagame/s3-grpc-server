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

type Client struct {
	ctx      context.Context
	bucket   string
	s3client *s3.Client
}

func NewClient(ctx context.Context, bucket string) (*Client, error) {
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

	return &Client{
		ctx:      ctx,
		s3client: s3client,
		bucket:   bucket,
	}, nil
}

func (x *Client) ListBuckets() (*s3.ListBucketsOutput, error) {
	return x.s3client.ListBuckets(x.ctx, &s3.ListBucketsInput{})
}

func (x *Client) CreateBucket() (*s3.CreateBucketOutput, error) {
	return x.s3client.CreateBucket(x.ctx, &s3.CreateBucketInput{
		Bucket: aws.String(x.bucket),
	})
}

func (x *Client) PutObject(key string, body io.Reader) (*s3.PutObjectOutput, error) {
	return x.s3client.PutObject(x.ctx, &s3.PutObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(key),
		Body:   body,
	})
}

func (x *Client) PutObjectWithString(key, content string) (*s3.PutObjectOutput, error) {
	return x.PutObject(key, strings.NewReader(content))
}

func (x *Client) GetObject(key string) (*s3.GetObjectOutput, error) {
	return x.s3client.GetObject(x.ctx, &s3.GetObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(key),
	})
}

func (x *Client) GetObjectWithString(key string) (string, error) {
	out, err := x.GetObject(key)
	if err != nil {
		return "", err
	}

	streamToString := func(stream io.Reader) string {
		buf := new(bytes.Buffer)
		buf.ReadFrom(stream)
		return buf.String()
	}

	return streamToString(out.Body), nil
}

func (x *Client) DeleteObject(key string) (*s3.DeleteObjectOutput, error) {
	return x.s3client.DeleteObject(x.ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(x.bucket),
		Key:    aws.String(key),
	})
}

func (x *Client) ListObjects(prefix string, nexttoken *string) (*s3.ListObjectsV2Output, error) {
	return x.s3client.ListObjectsV2(x.ctx, &s3.ListObjectsV2Input{
		Bucket:            aws.String(x.bucket),
		Prefix:            aws.String(prefix),
		ContinuationToken: nexttoken,
	})
}
