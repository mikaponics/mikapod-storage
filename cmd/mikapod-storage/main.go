package main // github.com/mikaponics/mikapod-soil/cmd/storage

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/mikaponics/mikapod-storage/api"
	"github.com/mikaponics/mikapod-soil/configs"
	"github.com/mikaponics/mikapod-storage/internal/storage"
	"github.com/mikaponics/mikapod-storage/internal/rpc"
)



func main() {
	// Create our database on program load if it has not been created previously.
	storage.InitMikapodStorage()

    // Open a TCP server to the specified localhost and environment variable
	// specified port number.
	lis, err := net.Listen("tcp", configs.StorageServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Initialize our gRPC server using our TCP server.
	grpcServer := grpc.NewServer()
	pb.RegisterMikapodStorageServer(grpcServer, &rpc.MikapodStorageGRPC{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
