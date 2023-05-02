package repository

import (
	"context"
	"sample/s3-grpc-server/grpc_server/service/repository/cell"
	"sample/s3-grpc-server/grpc_server/service/repository/fileinfo"
	"sample/s3-grpc-server/grpc_server/service/repository/table"
	server "sample/s3-grpc-server/proto/grpc_server"
)

type repositoryServer struct {
	fileinfo *fileinfo.FileInfoServer
	table    *table.TableServer
	cell     *cell.CellServer
	server.UnimplementedRepositoryServer
}

func NewRepositoryServer(
	fileinfo *fileinfo.FileInfoServer,
	table *table.TableServer,
	cell *cell.CellServer,
) *repositoryServer {
	return &repositoryServer{
		fileinfo: fileinfo,
		table:    table,
		cell:     cell,
	}
}

// fileinfo ----------------------------------------------------------------

// CreateFileInfo implements repositoryServer.CreateFileInfo
func (s *repositoryServer) CreateFileInfo(ctx context.Context, in *server.CreateFileInfoRequest) (*server.CreateFileInfoResponse, error) {
	return s.fileinfo.Create(ctx, in)
}

// ReadFileInfo implements repositoryServer.ReadFileInfo
func (s *repositoryServer) ReadFileInfo(ctx context.Context, in *server.ReadFileInfoRequest) (*server.ReadFileInfoResponse, error) {
	return s.fileinfo.Read(ctx, in)
}

// UpdateFileInfo implements repositoryServer.UpdateFileInfo
func (s *repositoryServer) UpdateFileInfo(ctx context.Context, in *server.UpdateFileInfoRequest) (*server.UpdateFileInfoResponse, error) {
	return s.fileinfo.Update(ctx, in)
}

// DeleteFileInfo implements repositoryServer.DeleteFileInfo
func (s *repositoryServer) DeleteFileInfo(ctx context.Context, in *server.DeleteFileInfoRequest) (*server.DeleteFileInfoResponse, error) {
	return s.fileinfo.Delete(ctx, in)
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
