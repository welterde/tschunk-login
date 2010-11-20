package server


import "os"
import "net"
import log "log4go"

import session "minecraft/session"


type Server struct {
	listener net.Listener
}

func Create(addr string) (server *Server, err os.Error) {
	// create structure
	server = new(Server)

	// try to listen on the socket
	log.Info("Trying to listen on %v", addr)
	server.listener, err = net.Listen("tcp", addr)
	return
}

func (server *Server) Serve() {
	for {
		// accept connection
		conn, err := server.listener.Accept()
		if err != nil {
			log.Warn("Error in accept() loop: %v", err)
			continue
		}

		// output this new connection
		log.Info("Got connection from %v on local socket %v", conn.RemoteAddr(), conn.LocalAddr())

		// start client
		session.StartSession(conn)
	}
}
