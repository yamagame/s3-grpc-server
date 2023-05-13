package storage

import (
	"context"
	"path/filepath"
	"sample/s3-grpc-server/entitiy/storage/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientZip(t *testing.T) {
	ctx := context.TODO()
	zip, err := NewZIPClient(ctx, filepath.Join("testdata", "sample.zip"))
	assert.NoError(t, err)
	ret1, err := zip.ListBuckets(ctx, &model.ListBuckets{})
	assert.NoError(t, err)
	assert.Equal(t, "testdata/sample.zip", ret1[0].Name)
	ret2, err := zip.ListObjects(ctx, &model.ListObjects{})
	assert.NoError(t, err)
	assert.Equal(t, "sample/hello2.txt", ret2[0].Key)
	assert.Equal(t, "sample/hello1.txt", ret2[1].Key)
	assert.Equal(t, "sample/dir/hello3.txt", ret2[2].Key)
	assert.Equal(t, "sample/dir/hello4.txt", ret2[3].Key)
	ret3, err := zip.GetObjectWithString(ctx, &model.GetObject{
		Key: "sample/hello1.txt",
	})
	assert.NoError(t, err)
	assert.Equal(t, "hello world 1\n", ret3)
	err = zip.Close()
	assert.NoError(t, err)
}
