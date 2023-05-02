package model

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
