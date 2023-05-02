package dto

import (
	"sample/s3-grpc-server/infra/storage/model"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type PutObject struct {
}

func (x *PutObject) Input(req *server.PutObjectRequest) *model.PutObject {
	return &model.PutObject{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *PutObject) Output(res *model.PutObject) *server.PutObjectResponse {
	return &server.PutObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type GetObject struct {
}

func (x *GetObject) Input(req *server.GetObjectRequest) *model.GetObject {
	return &model.GetObject{
		Key: req.Key,
	}
}

func (x *GetObject) Output(res *model.GetObject) *server.GetObjectResponse {
	return &server.GetObjectResponse{
		Result:  server.Result(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type DeleteObject struct {
}

func (x *DeleteObject) Input(req *server.DeleteObjectRequest) *model.DeleteObject {
	return &model.DeleteObject{
		Key: req.Key,
	}
}

func (x *DeleteObject) Output(res *model.DeleteObject) *server.DeleteObjectResponse {
	return &server.DeleteObjectResponse{
		Result: server.Result(res.Result),
		Key:    res.Key,
	}
}

type ListObjects struct {
}

func (x *ListObjects) Input(req *server.ListObjectsRequest) *model.ListObjects {
	return &model.ListObjects{
		Prefix: req.Prefix,
	}
}

func (x *ListObjects) Output(res *model.ListObjects) *server.ListObjectsResponse {
	return &server.ListObjectsResponse{
		Result: server.Result(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
