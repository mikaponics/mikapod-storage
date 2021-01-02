package rpc_server

// Simple function to call by your client to confirm you are connected with the server.
func (rpc *RPC) Greet(name *string, reply *string) error {
	*reply = "Hello, " + *name
	return nil
}
