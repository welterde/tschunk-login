package server


import "os"
import "net"

import session "minecraft/session"


type Server struct {
	listener net.Listener
}

func Create(addr string) (server *Server, err os.Error) {
	// create structure
	server = new(Server)

	// try to listen on the socket
	server.listener, err = net.Listen("tcp", addr)
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
		session.StartSession(server.core, conn)
	}
}
