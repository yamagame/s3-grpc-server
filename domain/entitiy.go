package domain

type AwsResult int32

const (
	Result_UNDEFINED AwsResult = 0
	Result_OK        AwsResult = 1
	Result_ERR       AwsResult = 2
)

type CreateBucketEntity struct {
	Result AwsResult
}

type Bucket struct {
	Name string
}

type ListBucketsEntity struct {
	Result  AwsResult
	Buckets []Bucket
}

type PutObjectEntity struct {
	Result  AwsResult
	Key     string
	Content string
}

type GetObjectEntity struct {
	Result  AwsResult
	Key     string
	Content string
}

type DeleteObjectEntity struct {
	Result AwsResult
	Key    string
}

type ListObjectsEntity struct {
	Result AwsResult
	Prefix string
	Keys   []string
	Next   *string
}
