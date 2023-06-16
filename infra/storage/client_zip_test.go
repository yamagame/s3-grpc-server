package storage

import (
	"context"
	"encoding/json"
	"path/filepath"
	"sample/s3-grpc-server/entitiy/storage/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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

func TestZap(t *testing.T) {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger := zap.Must(cfg.Build())
	suger := logger.Sugar()

	const url = "http://example.com"

	// 構造化ロギング
	logger.Info("Failed to fetch URL.",
		// Structured context as strongly typed fields.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	// 構造化ロギング(suger版)
	suger.Infow("Failed to fetch URL.",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)

	// fmt.Sprintを使用、値を文字に変換して連結
	suger.Info("A", "B", "C")

	logger.Sync()
}
