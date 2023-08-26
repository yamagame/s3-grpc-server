package file

import (
	"context"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/interceptor"
)

func validators() []*interceptor.ValidateRule {
	return []*interceptor.ValidateRule{
		interceptor.Validate(&grpc_server.CreateFileRequest{}, func(ctx context.Context, req *grpc_server.CreateFileRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.ReadFileRequest{}, func(ctx context.Context, req *grpc_server.ReadFileRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.UpdateFileRequest{}, func(ctx context.Context, req *grpc_server.UpdateFileRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.DeleteFileRequest{}, func(ctx context.Context, req *grpc_server.DeleteFileRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.ListFileRequest{}, func(ctx context.Context, req *grpc_server.ListFileRequest) error {
			return nil
		}),
	}
}
