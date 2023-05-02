package table

import (
	"context"
	"sample/s3-grpc-server/grpc_server/service/repository/dto"
	"sample/s3-grpc-server/infra/repository/table"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type serverDTO struct {
	createTable dto.CreateTable
	readTable   dto.ReadTable
	updateTable dto.UpdateTable
	deleteTable dto.DeleteTable
}

type TableServer struct {
	serverDTO
	service table.RepositoryInterface
	server.UnimplementedRepositoryServer
}

func NewTableServer(service table.RepositoryInterface) *TableServer {
	return &TableServer{
		service: service,
	}
}

// Create implements RepositoryServer.Create
func (s *TableServer) Create(ctx context.Context, in *server.CreateTableRequest) (*server.CreateTableResponse, error) {
	fileInfo := s.createTable.Input(in)
	fileInfo, err := s.service.Create(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.createTable.Output(fileInfo), nil
}

// Read implements RepositoryServer.Read
func (s *TableServer) Read(ctx context.Context, in *server.ReadTableRequest) (*server.ReadTableResponse, error) {
	fileInfo := s.readTable.Input(in)
	fileInfo, err := s.service.Read(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.readTable.Output(fileInfo), nil
}

// Update implements RepositoryServer.Update
func (s *TableServer) Update(ctx context.Context, in *server.UpdateTableRequest) (*server.UpdateTableResponse, error) {
	fileInfo := s.updateTable.Input(in)
	fileInfo, err := s.service.Update(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.updateTable.Output(fileInfo), nil
}

// Delete implements RepositoryServer.Delete
func (s *TableServer) Delete(ctx context.Context, in *server.DeleteTableRequest) (*server.DeleteTableResponse, error) {
	fileInfo := s.deleteTable.Input(in)
	fileInfo, err := s.service.Delete(fileInfo)
	if err != nil {
		return nil, err
	}
	return s.deleteTable.Output(fileInfo), nil
}
