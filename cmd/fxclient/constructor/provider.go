package constructor

import (
	"bufio"
	"context"
	"log"
	"os"

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

func GRPCClientConnection(conn *grpc.ClientConn) grpc.ClientConnInterface {
	return conn
}
