package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	ctx           context.Context
	bucket        string
	s3client      *s3.Client
	cognitoClient *cognitoidentityprovider.Client
	userPoolId    string
}

func NewClient(ctx context.Context, bucket string) (*Client, error) {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		fmt.Println("service", service)
		if service == s3.ServiceID {
			url := os.Getenv("S3_ENDPOINT")
			if url != "" {
				fmt.Println("S3_ENDPOINT:", url)
				return aws.Endpoint{URL: url, HostnameImmutable: true}, nil
			}
		}
		if service == cognitoidentityprovider.ServiceID {
			url := os.Getenv("COGNITO_ENDPOINT")
			if url != "" {
				fmt.Println("COGNITO_ENDPOINT:", url)
				// Authorization: AWS4-HMAC-SHA256 Credential=mock_access_key/20220524/us-east-1/cognito-idp/aws4_request, SignedHeaders=content-length;content-type;host;x-amz-date, Signature=asdf
				return aws.Endpoint{URL: url, HostnameImmutable: true}, nil
			}
		}
		return aws.Endpoint{}, &aws.EndpointNotFoundError{}
	})

	cred := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN"))
	if cred == nil {
		panic("failed to fetch credentials provider")
	}

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithEndpointResolverWithOptions(resolver),
		config.WithCredentialsProvider(cred),
	)
	if err != nil {
		return nil, err
	}

	s3client := s3.NewFromConfig(cfg)

	var userPoolId = "ap-northeast-1_XXXXXXXXX"
	cognitoClient := cognitoidentityprovider.NewFromConfig(cfg)

	return &Client{
		ctx:           ctx,
		s3client:      s3client,
		bucket:        bucket,
		cognitoClient: cognitoClient,
		userPoolId:    userPoolId,
	}, nil
}

// s3 -----------------------------------------------------------------------------------

func (x *Client) ListBuckets() (*s3.ListBucketsOutput, error) {
	return x.s3client.ListBuckets(x.ctx, &s3.ListBucketsInput{})
}

func (x *Client) CreateBucket(name string) (*s3.CreateBucketOutput, error) {
	return x.s3client.CreateBucket(x.ctx, &s3.CreateBucketInput{
		Bucket: aws.String(name),
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

// cognito -----------------------------------------------------------------------------------

func (x *Client) ListUsers() (*cognitoidentityprovider.ListUsersOutput, error) {
	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: &x.userPoolId,
	}
	output, err := x.cognitoClient.ListUsers(x.ctx, input)
	if err != nil {
		return nil, err
	}
	return output, nil
}
