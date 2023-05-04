package gateway

import (
	"sample/s3-grpc-server/infra/storage/model"
	"sample/s3-grpc-server/libs/gateway"
	"sample/s3-grpc-server/proto/grpc_server"
)

type PutObject struct {
}

func (x *PutObject) ToDomain(req *grpc_server.PutObjectRequest, call func(table *model.PutObject) (*model.PutObject, error)) (*grpc_server.PutObjectResponse, error) {
	return gateway.ToDomain[model.PutObject, grpc_server.PutObjectRequest, grpc_server.PutObjectResponse](x, req, call)
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

func (x *GetObject) ToDomain(req *grpc_server.GetObjectRequest, call func(table *model.GetObject) (*model.GetObject, error)) (*grpc_server.GetObjectResponse, error) {
	return gateway.ToDomain[model.GetObject, grpc_server.GetObjectRequest, grpc_server.GetObjectResponse](x, req, call)
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

func (x *DeleteObject) ToDomain(req *grpc_server.DeleteObjectRequest, call func(table *model.DeleteObject) (*model.DeleteObject, error)) (*grpc_server.DeleteObjectResponse, error) {
	return gateway.ToDomain[model.DeleteObject, grpc_server.DeleteObjectRequest, grpc_server.DeleteObjectResponse](x, req, call)
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

func (x *ListObjects) ToDomain(req *grpc_server.ListObjectsRequest, call func(table *model.ListObjects) (*model.ListObjects, error)) (*grpc_server.ListObjectsResponse, error) {
	return gateway.ToDomain[model.ListObjects, grpc_server.ListObjectsRequest, grpc_server.ListObjectsResponse](x, req, call)
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
