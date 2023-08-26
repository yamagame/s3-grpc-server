package cell

import (
	"context"
	"sample/s3-grpc-server/proto/grpc_server"
	"sample/s3-grpc-server/service/grpc/interceptor"
)

func validators() []*interceptor.ValidateRule {
	return []*interceptor.ValidateRule{
		interceptor.Validate(&grpc_server.CreateCellRequest{}, func(ctx context.Context, req *grpc_server.CreateCellRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.ReadCellRequest{}, func(ctx context.Context, req *grpc_server.ReadCellRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.UpdateCellRequest{}, func(ctx context.Context, req *grpc_server.UpdateCellRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.DeleteCellRequest{}, func(ctx context.Context, req *grpc_server.DeleteCellRequest) error {
			return nil
		}),
		interceptor.Validate(&grpc_server.ListCellRequest{}, func(ctx context.Context, req *grpc_server.ListCellRequest) error {
			return nil
		}),
	}
}
