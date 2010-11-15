package server


import "os"
import "net"

import core   "core"
import client "minecraft/client"


type Server struct {
	listener net.Listener
	core     *core.Core
}

func Create(addr string, core *core.Core) (server *Server, err os.Error) {
	// create structure
	server = new(Server)

	// try to listen on the socket
	server.listener, err = net.Listen("tcp", addr)

	// don't forget to add core
	server.core = core
	return
}

func (server *Server) Serve() {
	for {
		// accept connection
		conn, err := server.listener.Accept()
		if err != nil {
			// TODO: print it
			continue
		}

		// start client
		client.StartClient(server.core, conn)
	}
}
