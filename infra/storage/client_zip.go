package storage

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"sample/s3-grpc-server/entitiy/storage/model"
	"strings"
)

type ZIPClient struct {
	ctx      context.Context
	filename string
	zip      *zip.ReadCloser
}

func NewZIPClient(ctx context.Context, filename string) (*ZIPClient, error) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	return &ZIPClient{
		ctx:      ctx,
		filename: filename,
		zip:      r,
	}, nil
}

// Close implements ZIPClient.Close
func (x *ZIPClient) Close() error {
	return x.zip.Close()
}

// CreateBucket implements ZIPClient.CreateBucket
func (x *ZIPClient) CreateBucket(ctx context.Context, _ *model.CreateBucket) error {
	fmt.Println("CreateBucket")
	return nil
}

// ListBuckets implements ZIPClient.ListBuckets
func (x *ZIPClient) ListBuckets(ctx context.Context, _ *model.ListBuckets) ([]model.Bucket, error) {
	fmt.Println("ListBuckets")
	var ret []model.Bucket
	ret = append(ret, model.Bucket{
		Name: x.filename,
	})
	return ret, nil
}

// PutObject implements ZIPClient.PutObject
func (x *ZIPClient) PutObject(ctx context.Context, req *model.PutObject, body io.Reader) error {
	return nil
}

// PutObjectWithString implements ZIPClient.PutObjectWithString
func (x *ZIPClient) PutObjectWithString(ctx context.Context, req *model.PutObject) error {
	return x.PutObject(ctx, req, strings.NewReader(req.Content))
}

// GetObject implements ZIPClient.GetObject
func (x *ZIPClient) GetObject(ctx context.Context, req *model.GetObject) (io.Reader, error) {
	fmt.Println("GetObject", req.Key)
	return x.zip.Open(req.Key)
}

// GetObjectWithString implements ZIPClient.GetObjectWithString
func (x *ZIPClient) GetObjectWithString(ctx context.Context, req *model.GetObject) (string, error) {
	out, err := x.GetObject(ctx, req)
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

// DeleteObject implements ZIPClient.DeleteObject
func (x *ZIPClient) DeleteObject(ctx context.Context, req *model.DeleteObject) error {
	fmt.Println("DeleteObject", req.Key)
	return nil
}

// ListObjects implements ZIPClient.ListObjects
func (x *ZIPClient) ListObjects(ctx context.Context, req *model.ListObjects) ([]model.Object, error) {
	fmt.Println("ListObjects", req.Prefix)
	var ret []model.Object
	for _, f := range x.zip.File {
		if !f.FileInfo().IsDir() {
			ret = append(ret, model.Object{
				Key: f.Name,
			})
		}
	}
	return ret, nil
}
