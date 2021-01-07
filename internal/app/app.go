package app // github.com/mikaponics/mikapod-storage/internal

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/mikaponics/mikapod-storage/configs"
	"github.com/mikaponics/mikapod-storage/internal/rpc_server"
	"github.com/mikaponics/mikapod-storage/internal/storage"
)

type MikapodStorage struct {
	tcpAddr   *net.TCPAddr
	listener  *net.TCPListener
	rpcServer *rpc_server.RPC
}

func InitMikapodStorage() *MikapodStorage {
	tcpAddr, err := net.ResolveTCPAddr("tcp", configs.MikapodStorageServiceAddress)
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.InitMikapodDB()

	r := &rpc_server.RPC{
		Store: storage,
	}

	log.Println("RPC API was initialized.")
	return &MikapodStorage{
		tcpAddr:   tcpAddr,
		listener:  nil,
		rpcServer: r,
	}
}

// Function will consume the main runtime loop and run the business logic
// of the Mikapod Logger application.
func (app *MikapodStorage) RunMainRuntimeLoop() {
	rpc.Register(app.rpcServer)
	rpc.HandleHTTP()
	log.Println("RPC was initialized.")
	l, e := net.ListenTCP("tcp", app.tcpAddr)
	app.listener = l // Track the `listener` so we can gracefully shutdown later.
	if e != nil {
		log.Fatal("listen error:", e.Error())
	}
	log.Println("Started storage service.")
	http.Serve(l, nil)
}

// Function will tell the application to stop the main runtime loop when
// the process has been finished.
func (app *MikapodStorage) StopMainRuntimeLoop() {
	log.Printf("Starting graceful shutdown now...")

	// Finish any RPC communication taking place at the moment before
	// shutting down the RPC server.
	app.listener.Close()
}
