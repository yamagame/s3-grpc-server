package storage

type StorageResult int32

const (
	Result_UNDEFINED StorageResult = 0
	Result_OK        StorageResult = 1
	Result_ERR       StorageResult = 2
)

type CreateBucket struct {
	Result StorageResult
}

type Bucket struct {
	Name string
}

type ListBuckets struct {
	Result  StorageResult
	Buckets []Bucket
}

type Object struct {
	Key     string
	Content string
}

type PutObject struct {
	Result  StorageResult
	Key     string
	Content string
}

type GetObject struct {
	Result  StorageResult
	Key     string
	Content string
}

type DeleteObject struct {
	Result StorageResult
	Key    string
}

type ListObjects struct {
	Result StorageResult
	Prefix string
	Keys   []string
	Next   *string
}
