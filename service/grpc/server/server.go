package grpc

import (
	"context"
	cellService "sample/s3-grpc-server/service/grpc/server/repository/cell"
	fileService "sample/s3-grpc-server/service/grpc/server/repository/file"
	tableService "sample/s3-grpc-server/service/grpc/server/repository/table"
	storageService "sample/s3-grpc-server/service/grpc/server/storage"

	"sample/s3-grpc-server/infra/repository/cell"
	"sample/s3-grpc-server/infra/repository/file"
	"sample/s3-grpc-server/infra/repository/table"
	"sample/s3-grpc-server/infra/storage"

	rpc "sample/s3-grpc-server/proto/grpc_server"

	"google.golang.org/grpc"
)

func NewServer(
	storageClient storage.StorageInterface,
	fileRepository file.RepositoryInterface,
	tableRepository table.RepositoryInterface,
	cellRepository cell.RepositoryInterface,
) *grpc.Server {
	server := grpc.NewServer()
	rpc.RegisterStorageServer(server, storageService.NewStorageServer(storage.NewStorageService(storageClient)))
	rpc.RegisterRepositoryServer(server,
		newCRUDRepositoryServer(
			fileService.NewFileRepository(fileRepository),
			tableService.NewTableRepository(tableRepository),
			cellService.NewCellRepository(cellRepository),
		),
	)
	return server
}

type CRUDRepositoryServer struct {
	file  *fileService.FileRepository
	table *tableService.TableRepository
	cell  *cellService.CellRepository
	rpc.UnimplementedRepositoryServer
}

func newCRUDRepositoryServer(
	file *fileService.FileRepository,
	table *tableService.TableRepository,
	cell *cellService.CellRepository,
) *CRUDRepositoryServer {
	return &CRUDRepositoryServer{
		file:  file,
		table: table,
		cell:  cell,
	}
}

// file ----------------------------------------------------------------

// CreateFile implements repositoryServer.CreateFile
func (s *CRUDRepositoryServer) CreateFile(ctx context.Context, in *rpc.CreateFileRequest) (*rpc.CreateFileResponse, error) {
	return s.file.Create(ctx, in)
}

// ReadFile implements repositoryServer.ReadFile
func (s *CRUDRepositoryServer) ReadFile(ctx context.Context, in *rpc.ReadFileRequest) (*rpc.ReadFileResponse, error) {
	return s.file.Read(ctx, in)
}

// UpdateFile implements repositoryServer.UpdateFile
func (s *CRUDRepositoryServer) UpdateFile(ctx context.Context, in *rpc.UpdateFileRequest) (*rpc.UpdateFileResponse, error) {
	return s.file.Update(ctx, in)
}

// DeleteFile implements repositoryServer.DeleteFile
func (s *CRUDRepositoryServer) DeleteFile(ctx context.Context, in *rpc.DeleteFileRequest) (*rpc.DeleteFileResponse, error) {
	return s.file.Delete(ctx, in)
}

// table -------------------------------------------------------------------

// CreateTable implements repositoryServer.CreateTable
func (s *CRUDRepositoryServer) CreateTable(ctx context.Context, in *rpc.CreateTableRequest) (*rpc.CreateTableResponse, error) {
	return s.table.Create(ctx, in)
}

// ReadTable implements repositoryServer.ReadTable
func (s *CRUDRepositoryServer) ReadTable(ctx context.Context, in *rpc.ReadTableRequest) (*rpc.ReadTableResponse, error) {
	return s.table.Read(ctx, in)
}

// UpdateTable implements repositoryServer.UpdateTable
func (s *CRUDRepositoryServer) UpdateTable(ctx context.Context, in *rpc.UpdateTableRequest) (*rpc.UpdateTableResponse, error) {
	return s.table.Update(ctx, in)
}

// DeleteTable implements repositoryServer.DeleteTable
func (s *CRUDRepositoryServer) DeleteTable(ctx context.Context, in *rpc.DeleteTableRequest) (*rpc.DeleteTableResponse, error) {
	return s.table.Delete(ctx, in)
}

// cell --------------------------------------------------------------------

// CreateCell implements repositoryServer.CreateCell
func (s *CRUDRepositoryServer) CreateCell(ctx context.Context, in *rpc.CreateCellRequest) (*rpc.CreateCellResponse, error) {
	return s.cell.Create(ctx, in)
}

// ReadCell implements repositoryServer.ReadCell
func (s *CRUDRepositoryServer) ReadCell(ctx context.Context, in *rpc.ReadCellRequest) (*rpc.ReadCellResponse, error) {
	return s.cell.Read(ctx, in)
}

// UpdateCell implements repositoryServer.UpdateCell
func (s *CRUDRepositoryServer) UpdateCell(ctx context.Context, in *rpc.UpdateCellRequest) (*rpc.UpdateCellResponse, error) {
	return s.cell.Update(ctx, in)
}

// DeleteCell implements repositoryServer.DeleteCell
func (s *CRUDRepositoryServer) DeleteCell(ctx context.Context, in *rpc.DeleteCellRequest) (*rpc.DeleteCellResponse, error) {
	return s.cell.Delete(ctx, in)
}
