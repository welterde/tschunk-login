package main

import server "minecraft/server"
import daemon "minecraft/daemon"
import log    "log4go"

func main() {
	// load logging configuration
	log.LoadConfiguration("etc/logging.xml")

	// create daemon object
	d := daemon.NewDaemon()

	// create server
	server, _ := server.Create(d, "127.0.0.1:25565")

	// and run
	server.Serve()
}
