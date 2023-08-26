package table

import (
	"context"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/interceptor"
)

func validators() []*interceptor.ValidateRule {
	return []*interceptor.ValidateRule{
		interceptor.Validate(&grpc_server.CreateTableRequest{}, func(ctx context.Context, req *grpc_server.CreateTableRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.ReadTableRequest{}, func(ctx context.Context, req *grpc_server.ReadTableRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.UpdateTableRequest{}, func(ctx context.Context, req *grpc_server.UpdateTableRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.DeleteTableRequest{}, func(ctx context.Context, req *grpc_server.DeleteTableRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.ListTableRequest{}, func(ctx context.Context, req *grpc_server.ListTableRequest) error {
			return nil
		}),
	}
}
