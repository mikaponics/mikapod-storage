package app // github.com/mikaponics/mikapod-storage/internal

import (
    "log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/mikaponics/mikapod-storage/api"
	"github.com/mikaponics/mikapod-storage/configs"
	"github.com/mikaponics/mikapod-storage/internal/storage"
)

type MikapodStorage struct {
    db *storage.MikapodDB
    grpcServer *grpc.Server
}

func InitMikapodStorage() (*MikapodStorage) {
    return &MikapodStorage{
        db: storage.InitMikapodDB(),
        grpcServer: nil,
    }
}


// Function will consume the main runtime loop and run the business logic
// of the Mikapod Logger application.
func (app *MikapodStorage) RunMainRuntimeLoop() {
    // Open a TCP server to the specified localhost and environment variable
    // specified port number.
    lis, err := net.Listen("tcp", configs.MikapodStorageServiceAddress)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // Initialize our gRPC server using our TCP server.
    grpcServer := grpc.NewServer()

    // Save reference to our application state.
    app.grpcServer = grpcServer

    // For debugging purposes only.
    log.Printf("gRPC server is running.")

    // Block the main runtime loop for accepting and processing gRPC requests.
    pb.RegisterMikapodStorageServer(grpcServer, &MikapodStorageGRPC{
        // DEVELOPERS NOTE:
        // We want to attach to every gRPC call the following variables...
        db: app.db,
    })
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (app *MikapodStorage) StopMainRuntimeLoop() {
    log.Printf("Starting graceful shutdown now...")

    // Finish any RPC communication taking place at the moment before
    // shutting down the gRPC server.
    app.grpcServer.GracefulStop()
}
