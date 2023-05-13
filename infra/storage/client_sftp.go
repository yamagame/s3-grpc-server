package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path"
	"path/filepath"
	"sample/s3-grpc-server/entitiy/storage/model"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SFTPClientConfig struct {
	Host  string
	Port  string
	User  string
	Pass  string
	Share string
}

type SFTPClient struct {
	ctx    context.Context
	share  string
	client *sftp.Client
}

func NewSFTPClient(ctx context.Context, options SFTPClientConfig) (*SFTPClient, error) {
	fmt.Println("Create sshClientConfig")
	sshConfig := &ssh.ClientConfig{
		User: options.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(options.Pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	fmt.Println("SSH connect")
	addr := fmt.Sprintf("%s:%s", options.Host, options.Port)
	fmt.Println(addr)

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	fmt.Println("open an SFTP session over an existing ssh connection")
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}

	return &SFTPClient{
		ctx:    ctx,
		share:  options.Share,
		client: client,
	}, nil
}

// Close implements SFTPClient.Close
func (x *SFTPClient) Close() error {
	return x.client.Close()
}

// CreateBucket implements SFTPClient.CreateBucket
func (x *SFTPClient) CreateBucket(ctx context.Context, _ *model.CreateBucket) error {
	fmt.Println("CreateBucket", x.share)
	return nil
}

// ListBuckets implements SFTPClient.ListBuckets
func (x *SFTPClient) ListBuckets(ctx context.Context, _ *model.ListBuckets) ([]model.Bucket, error) {
	var buckets []model.Bucket
	fmt.Println("ListBuckets")
	buckets = append(buckets, model.Bucket{
		Name: x.share,
	})
	return buckets, nil
}

// PutObject implements SFTPClient.PutObject
func (x *SFTPClient) PutObject(ctx context.Context, req *model.PutObject, body io.Reader) error {
	fmt.Println("leave your mark")
	wc, err := x.client.Create(path.Join(x.share, req.Key))
	if err != nil {
		return err
	}
	if _, err := io.Copy(wc, body); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

// PutObjectWithString implements SFTPClient.PutObjectWithString
func (x *SFTPClient) PutObjectWithString(ctx context.Context, req *model.PutObject) error {
	return x.PutObject(ctx, req, strings.NewReader(req.Content))
}

// GetObject implements SFTPClient.GetObject
func (x *SFTPClient) GetObject(ctx context.Context, req *model.GetObject) (io.Reader, error) {
	fmt.Println("GetObject", req.Key)
	p := path.Join(x.share, req.Key)
	return x.client.Open(p)
}

// GetObjectWithString implements SFTPClient.GetObjectWithString
func (x *SFTPClient) GetObjectWithString(ctx context.Context, req *model.GetObject) (string, error) {
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

// DeleteObject implements SFTPClient.DeleteObject
func (x *SFTPClient) DeleteObject(ctx context.Context, req *model.DeleteObject) error {
	fmt.Println("DeleteObject", req.Key)
	return x.client.Remove(path.Join(x.share, req.Key))
}

// ListObjects implements SFTPClient.ListObjects
func (x *SFTPClient) ListObjects(ctx context.Context, req *model.ListObjects) ([]model.Object, error) {
	var objects []model.Object
	fmt.Println("ListObjects", req.Prefix)
	w := x.client.Walk(path.Join(x.share, req.Prefix))
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		if !w.Stat().IsDir() {
			name := filepath.Base(w.Path())
			if name[0] != '.' {
				objects = append(objects, model.Object{
					Key: strings.TrimPrefix(w.Path(), x.share),
				})
			}
		}
	}
	return objects, nil
}
