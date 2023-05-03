package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"sample/s3-grpc-server/infra/storage/model"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
)

type GCSClientConfig struct {
	Bucket    string
	ProjectID string
}

type GCSClient struct {
	ctx       context.Context
	bucket    string
	projectID string
	client    *storage.Client
}

func NewGCSClient(ctx context.Context, options GCSClientConfig) (*GCSClient, error) {
	// FakeGCSを使用する場合はSTORAGE_EMULATOR_HOSTに値をセットする
	// err := os.Setenv("STORAGE_EMULATOR_HOST", "localhost:4443")
	// if err != nil {
	// 	return nil, err
	// }

	client, err := storage.NewClient(ctx)

	if err != nil {
		return nil, err
	}

	return &GCSClient{
		ctx:       ctx,
		client:    client,
		bucket:    options.Bucket,
		projectID: options.ProjectID,
	}, nil
}

func (x *GCSClient) ListBuckets(ctx context.Context) ([]model.Bucket, error) {
	var buckets []model.Bucket
	it := x.client.Buckets(x.ctx, x.projectID)
	for {
		oattrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, model.Bucket{
			Name: oattrs.Name,
		})
	}
	return buckets, nil
}

func (x *GCSClient) CreateBucket(ctx context.Context) error {
	err := x.client.Bucket(x.bucket).Create(x.ctx, x.projectID, nil)
	if err != nil {
		fmt.Println(err)
	}
	if e, ok := err.(*googleapi.Error); ok {
		if e.Code == 409 {
			return nil
		}
	}
	return err
}

func (x *GCSClient) PutObject(ctx context.Context, key string, body io.Reader) error {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		return err
	}
	wc := x.client.Bucket(x.bucket).Object(key).NewWriter(x.ctx)
	_, err := wc.Write(buf.Bytes())
	if err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (x *GCSClient) PutObjectWithString(ctx context.Context, key, content string) error {
	return x.PutObject(ctx, key, strings.NewReader(content))
}

func (x *GCSClient) GetObject(ctx context.Context, key string) (io.Reader, error) {
	return x.client.Bucket(x.bucket).Object(key).NewReader(x.ctx)
}

func (x *GCSClient) GetObjectWithString(ctx context.Context, key string) (string, error) {
	out, err := x.GetObject(ctx, key)
	if err != nil {
		return "", err
	}

	streamToString := func(stream io.Reader) string {
		buf := new(bytes.Buffer)
		buf.ReadFrom(stream)
		return buf.String()
	}

	return streamToString(out), nil
}

func (x *GCSClient) DeleteObject(ctx context.Context, key string) error {
	return x.client.Bucket(x.bucket).Object(key).Delete(x.ctx)
}

func (x *GCSClient) ListObjects(ctx context.Context, prefix string, nexttoken *string) ([]model.Object, error) {
	var objects []model.Object
	it := x.client.Bucket(x.bucket).Objects(x.ctx, &storage.Query{
		Prefix: prefix,
	})
	for {
		oattrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objects = append(objects, model.Object{
			Key: oattrs.Name,
		})
	}
	return objects, nil
}
