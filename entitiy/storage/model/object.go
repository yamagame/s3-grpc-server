package model

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
