package main

import server "minecraft/server"

func main() {
	// create server
	server, _ := server.Create("127.0.0.1:25565")

	// and run
	server.Serve()
}
