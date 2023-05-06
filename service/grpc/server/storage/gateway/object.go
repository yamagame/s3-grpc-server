package gateway

import (
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type PutObject struct {
}

func (x *PutObject) Input(req *grpc_server.PutObjectRequest) *model.PutObject {
	return &model.PutObject{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *PutObject) Output(res *model.PutObject) *grpc_server.PutObjectResponse {
	return &grpc_server.PutObjectResponse{
		Result: grpc_server.Result(res.Result),
		Key:    res.Key,
	}
}

type GetObject struct {
}

func (x *GetObject) Input(req *grpc_server.GetObjectRequest) *model.GetObject {
	return &model.GetObject{
		Key: req.Key,
	}
}

func (x *GetObject) Output(res *model.GetObject) *grpc_server.GetObjectResponse {
	return &grpc_server.GetObjectResponse{
		Result:  grpc_server.Result(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type DeleteObject struct {
}

func (x *DeleteObject) Input(req *grpc_server.DeleteObjectRequest) *model.DeleteObject {
	return &model.DeleteObject{
		Key: req.Key,
	}
}

func (x *DeleteObject) Output(res *model.DeleteObject) *grpc_server.DeleteObjectResponse {
	return &grpc_server.DeleteObjectResponse{
		Result: grpc_server.Result(res.Result),
		Key:    res.Key,
	}
}

type ListObjects struct {
}

func (x *ListObjects) Input(req *grpc_server.ListObjectsRequest) *model.ListObjects {
	return &model.ListObjects{
		Prefix: req.Prefix,
	}
}

func (x *ListObjects) Output(res *model.ListObjects) *grpc_server.ListObjectsResponse {
	return &grpc_server.ListObjectsResponse{
		Result: grpc_server.Result(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
