package dto

import (
	"sample/s3-grpc-server/infra/storage/model"
	"sample/s3-grpc-server/libs/dto"
	"sample/s3-grpc-server/proto/grpc_server"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type PutObject struct {
}

func (x *PutObject) Domain(req *grpc_server.PutObjectRequest, call func(table *model.PutObject) (*model.PutObject, error)) (*grpc_server.PutObjectResponse, error) {
	return dto.ToDomain[model.PutObject, grpc_server.PutObjectRequest, grpc_server.PutObjectResponse](x, req, call)
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

func (x *GetObject) Domain(req *grpc_server.GetObjectRequest, call func(table *model.GetObject) (*model.GetObject, error)) (*grpc_server.GetObjectResponse, error) {
	return dto.ToDomain[model.GetObject, grpc_server.GetObjectRequest, grpc_server.GetObjectResponse](x, req, call)
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

func (x *DeleteObject) Domain(req *grpc_server.DeleteObjectRequest, call func(table *model.DeleteObject) (*model.DeleteObject, error)) (*grpc_server.DeleteObjectResponse, error) {
	return dto.ToDomain[model.DeleteObject, grpc_server.DeleteObjectRequest, grpc_server.DeleteObjectResponse](x, req, call)
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

func (x *ListObjects) Domain(req *grpc_server.ListObjectsRequest, call func(table *model.ListObjects) (*model.ListObjects, error)) (*grpc_server.ListObjectsResponse, error) {
	return dto.ToDomain[model.ListObjects, grpc_server.ListObjectsRequest, grpc_server.ListObjectsResponse](x, req, call)
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
