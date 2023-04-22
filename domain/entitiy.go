package domain

type StorageResult int32

const (
	Result_UNDEFINED StorageResult = 0
	Result_OK        StorageResult = 1
	Result_ERR       StorageResult = 2
)

type CreateBucketEntity struct {
	Result StorageResult
}

type Bucket struct {
	Name string
}

type ListBucketsEntity struct {
	Result  StorageResult
	Buckets []Bucket
}

type PutObjectEntity struct {
	Result  StorageResult
	Key     string
	Content string
}

type GetObjectEntity struct {
	Result  StorageResult
	Key     string
	Content string
}

type DeleteObjectEntity struct {
	Result StorageResult
	Key    string
}

type ListObjectsEntity struct {
	Result StorageResult
	Prefix string
	Keys   []string
	Next   *string
}
