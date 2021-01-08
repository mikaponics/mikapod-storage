package rpc_client

import (
	// "errors"
	"log"
	"net/rpc"
	"time"
)

type MikapodStorageService struct {
	Client *rpc.Client
}

// Function will attempt to connect to the RPC server and if the attempt fails
// then wait 10 seconds and retry - if 100 attempts fail then a fatal panic
// will be triggered.
func retryMultipleAndBlockingDailHTTP(addr string) (*rpc.Client) {
	i := 0
	for {
		client, err := rpc.DialHTTP("tcp", addr)
		if err != nil {
			i = i + 1
			if i >= 100 {
				log.Fatal("ERROR | MikapodStorageService | retryMultipleAndBlockingDailHTTP | Dialing TCP Error:", err)
			}
			time.Sleep(10 * time.Second)
		} else {
			return client
		}
	}
	return nil
}

func New(addr string) *MikapodStorageService {
	if addr == "" {
		log.Fatal("LOG | MikapodStorageService | New | No address set.")
	}

    // The following code will dial the server or attempt to redial the server
	// 100 times before it the client returns a fatal error.
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		client = retryMultipleAndBlockingDailHTTP(addr)
	}

	return &MikapodStorageService{
		Client: client,
	}
}
