package repository

import (
	"context"
	"sample/s3-grpc-server/grpc_server/service/repository/cell"
	"sample/s3-grpc-server/grpc_server/service/repository/file"
	"sample/s3-grpc-server/grpc_server/service/repository/table"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type repositoryServer struct {
	file  *file.CRUDService
	table *table.CRUDService
	cell  *cell.CRUDService
	server.UnimplementedRepositoryServer
}

func NewRepositoryServer(
	file *file.CRUDService,
	table *table.CRUDService,
	cell *cell.CRUDService,
) *repositoryServer {
	return &repositoryServer{
		file:  file,
		table: table,
		cell:  cell,
	}
}

// file ----------------------------------------------------------------

// CreateFile implements repositoryServer.CreateFile
func (s *repositoryServer) CreateFile(ctx context.Context, in *server.CreateFileRequest) (*server.CreateFileResponse, error) {
	return s.file.Create(ctx, in)
}

// ReadFile implements repositoryServer.ReadFile
func (s *repositoryServer) ReadFile(ctx context.Context, in *server.ReadFileRequest) (*server.ReadFileResponse, error) {
	return s.file.Read(ctx, in)
}

// UpdateFile implements repositoryServer.UpdateFile
func (s *repositoryServer) UpdateFile(ctx context.Context, in *server.UpdateFileRequest) (*server.UpdateFileResponse, error) {
	return s.file.Update(ctx, in)
}

// DeleteFile implements repositoryServer.DeleteFile
func (s *repositoryServer) DeleteFile(ctx context.Context, in *server.DeleteFileRequest) (*server.DeleteFileResponse, error) {
	return s.file.Delete(ctx, in)
}

// table -------------------------------------------------------------------

// CreateTable implements repositoryServer.CreateTable
func (s *repositoryServer) CreateTable(ctx context.Context, in *server.CreateTableRequest) (*server.CreateTableResponse, error) {
	return s.table.Create(ctx, in)
}

// ReadTable implements repositoryServer.ReadTable
func (s *repositoryServer) ReadTable(ctx context.Context, in *server.ReadTableRequest) (*server.ReadTableResponse, error) {
	return s.table.Read(ctx, in)
}

// UpdateTable implements repositoryServer.UpdateTable
func (s *repositoryServer) UpdateTable(ctx context.Context, in *server.UpdateTableRequest) (*server.UpdateTableResponse, error) {
	return s.table.Update(ctx, in)
}

// DeleteTable implements repositoryServer.DeleteTable
func (s *repositoryServer) DeleteTable(ctx context.Context, in *server.DeleteTableRequest) (*server.DeleteTableResponse, error) {
	return s.table.Delete(ctx, in)
}

// cell --------------------------------------------------------------------

// CreateCell implements repositoryServer.CreateCell
func (s *repositoryServer) CreateCell(ctx context.Context, in *server.CreateCellRequest) (*server.CreateCellResponse, error) {
	return s.cell.Create(ctx, in)
}

// ReadCell implements repositoryServer.ReadCell
func (s *repositoryServer) ReadCell(ctx context.Context, in *server.ReadCellRequest) (*server.ReadCellResponse, error) {
	return s.cell.Read(ctx, in)
}

// UpdateCell implements repositoryServer.UpdateCell
func (s *repositoryServer) UpdateCell(ctx context.Context, in *server.UpdateCellRequest) (*server.UpdateCellResponse, error) {
	return s.cell.Update(ctx, in)
}

// DeleteCell implements repositoryServer.DeleteCell
func (s *repositoryServer) DeleteCell(ctx context.Context, in *server.DeleteCellRequest) (*server.DeleteCellResponse, error) {
	return s.cell.Delete(ctx, in)
}
