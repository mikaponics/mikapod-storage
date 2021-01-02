package rpc_server

import (
	"github.com/mikaponics/mikapod-storage/internal/storage"
)

type RPC struct{
	Store *storage.MikapodDB
}
