package gateway

import (
	"sample/s3-grpc-server/entitiy/storage/model"
	"sample/s3-grpc-server/proto/grpc_server"
)

type PutObject struct {
}

func (x *PutObject) Input(req *model.PutObject) *grpc_server.PutObjectRequest {
	return &grpc_server.PutObjectRequest{
		Key:     req.Key,
		Content: req.Content,
	}
}

func (x *PutObject) Output(res *grpc_server.PutObjectResponse) *model.PutObject {
	return &model.PutObject{
		Result: model.StorageResult(res.Result),
	}
}

type GetObject struct {
}

func (x *GetObject) Input(req *model.GetObject) *grpc_server.GetObjectRequest {
	return &grpc_server.GetObjectRequest{
		Key: req.Key,
	}
}

func (x *GetObject) Output(res *grpc_server.GetObjectResponse) *model.GetObject {
	return &model.GetObject{
		Result:  model.StorageResult(res.Result),
		Key:     res.Key,
		Content: res.Content,
	}
}

type DeleteObject struct {
}

func (x *DeleteObject) Input(req *model.DeleteObject) *grpc_server.DeleteObjectRequest {
	return &grpc_server.DeleteObjectRequest{
		Key: req.Key,
	}
}

func (x *DeleteObject) Output(res *grpc_server.DeleteObjectResponse) *model.DeleteObject {
	return &model.DeleteObject{
		Result: model.StorageResult(res.Result),
	}
}

type ListObjects struct {
}

func (x *ListObjects) Input(req *model.ListObjects) *grpc_server.ListObjectsRequest {
	return &grpc_server.ListObjectsRequest{
		Prefix: req.Prefix,
	}
}

func (x *ListObjects) Output(res *grpc_server.ListObjectsResponse) *model.ListObjects {
	return &model.ListObjects{
		Result: model.StorageResult(res.Result),
		Prefix: res.Prefix,
		Keys:   res.Keys,
		Next:   res.Next,
	}
}
