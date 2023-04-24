package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path"
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

func (x *SFTPClient) ListBuckets() ([]Bucket, error) {
	var buckets []Bucket
	fmt.Println("ListBuckets")
	buckets = append(buckets, Bucket{
		Name: x.share,
	})
	return buckets, nil
}

func (x *SFTPClient) CreateBucket() error {
	fmt.Println("CreateBucket", x.share)
	return nil
}

func (x *SFTPClient) PutObject(key string, body io.Reader) error {
	fmt.Println("leave your mark")
	wc, err := x.client.Create(path.Join(x.share, key))
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, body); err != nil {
		return err
	}
	if _, err := wc.Write(buf.Bytes()); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (x *SFTPClient) PutObjectWithString(key, content string) error {
	return x.PutObject(key, strings.NewReader(content))
}

func (x *SFTPClient) GetObject(key string) (io.Reader, error) {
	fmt.Println("GetObject", key)
	p := path.Join(x.share, key)
	return x.client.Open(p)
}

func (x *SFTPClient) GetObjectWithString(key string) (string, error) {
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

func (x *SFTPClient) DeleteObject(key string) error {
	fmt.Println("DeleteObject", key)
	return x.client.Remove(path.Join(x.share, key))
}

func (x *SFTPClient) ListObjects(prefix string, nexttoken *string) ([]Object, error) {
	var objects []Object
	fmt.Println("ListObjects", prefix)
	w := x.client.Walk(path.Join(x.share, prefix))
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		if !w.Stat().IsDir() {
			objects = append(objects, Object{
				Key: strings.TrimPrefix(w.Path(), x.share),
			})
		}
	}
	return objects, nil
}
