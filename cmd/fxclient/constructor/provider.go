package constructor

import (
	"bufio"
	"context"
	"log"
	"os"
	"sample/s3-grpc-server/service/grpc/client/repository/cell"
	"sample/s3-grpc-server/service/grpc/client/repository/file"
	"sample/s3-grpc-server/service/grpc/client/repository/table"
	"sample/s3-grpc-server/service/grpc/client/storage"
	"sample/s3-grpc-server/share/sheet"
	"sample/s3-grpc-server/usecase"

	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGRPCConnection() (*grpc.ClientConn, error) {
	address := os.Getenv("GRPC_HOST")
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return nil, err
	}
	return conn, nil
}

func NewContext() context.Context {
	return context.Background()
}

func NewKeyboardScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

type DBWriterIn struct {
	dig.In
	FileRepository    *file.FileClientRepository
	TableRepository   *table.TableClientRepository
	CellRepository    *cell.CellClientRepository
	StorageRepository *storage.StorageClientRepository
	Sheet             *sheet.CSVSheet
}

func NewDBWriter(in DBWriterIn) *usecase.DBWriter {
	return usecase.NewDBWriter(
		in.FileRepository,
		in.TableRepository,
		in.CellRepository,
		in.StorageRepository,
		in.Sheet,
	)
}

func GRPCClientConnection(conn *grpc.ClientConn) grpc.ClientConnInterface {
	return conn
}
