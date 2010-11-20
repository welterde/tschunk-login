package main

import server "minecraft/server"
import log "log4go"

func main() {
	// load logging configuration
	log.LoadConfiguration("etc/logging.xml")
	
	// create server
	server, _ := server.Create("127.0.0.1:25565")

	// and run
	server.Serve()
}
