package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
)

type GCSClient struct {
	ctx       context.Context
	bucket    string
	client    *storage.Client
	projectID string
}

func NewGCSClient(ctx context.Context, bucket string) (*GCSClient, error) {
	// FakeGCSを使用する場合はSTORAGE_EMULATOR_HOSTに値をセットする
	err := os.Setenv("STORAGE_EMULATOR_HOST", "localhost:4443")
	if err != nil {
		return nil, err
	}

	client, err := storage.NewClient(ctx) // option.WithEndpoint("http://localhost:8080/storage/v1/")

	if err != nil {
		return nil, err
	}

	return &GCSClient{
		ctx:       ctx,
		client:    client,
		bucket:    bucket,
		projectID: "kobo-2022",
	}, nil
}

func (x *GCSClient) ListBuckets() ([]Bucket, error) {
	var buckets []Bucket
	it := x.client.Buckets(x.ctx, x.projectID)
	for {
		oattrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, Bucket{
			Name: oattrs.Name,
		})
	}
	return buckets, nil
}

func (x *GCSClient) CreateBucket() error {
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

func (x *GCSClient) PutObject(key string, body io.Reader) error {
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

func (x *GCSClient) PutObjectWithString(key, content string) error {
	return x.PutObject(key, strings.NewReader(content))
}

func (x *GCSClient) GetObject(key string) (io.Reader, error) {
	return x.client.Bucket(x.bucket).Object(key).NewReader(x.ctx)
}

func (x *GCSClient) GetObjectWithString(key string) (string, error) {
	out, err := x.GetObject(key)
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

func (x *GCSClient) DeleteObject(key string) error {
	return x.client.Bucket(x.bucket).Object(key).Delete(x.ctx)
}

func (x *GCSClient) ListObjects(prefix string, nexttoken *string) ([]Object, error) {
	var objects []Object
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
		objects = append(objects, Object{
			Key: oattrs.Name,
		})
	}
	return objects, nil
}
