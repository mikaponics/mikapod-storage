package rpc_client

import (
	// "errors"
	"log"
	"net/rpc"
)

type MikapodStorageService struct {
	Client *rpc.Client
}

func New(addr string) *MikapodStorageService {
	if addr == "" {
		log.Fatal("LOG | MikapodStorageService | New | No address set.")
	}

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal("ERROR | MikapodStorageService | New | Dialing TCP Error:", err)
	}

	return &MikapodStorageService{
		Client: client,
	}
}
