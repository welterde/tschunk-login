package session


import log "log4go"

import packets "minecraft/packets"
import entity  "minecraft/entity"


type Handler func(sess *Session, packet packets.Packet)

type Session struct {
	RXQueue       chan packets.Packet
	TXQueue       chan packets.Packet
	Daemon        Daemon
	handlers      map[byte]Handler
	EntityManager *entity.EntityManager
}

func StartSession(daemon Daemon) (sess *Session) {
	// create session instance
	sess = &Session{
		RXQueue:       make(chan packets.Packet, 50),
		TXQueue:       make(chan packets.Packet, 1024),
		Daemon:        daemon,
		handlers:      make(map[byte]Handler),
		EntityManager: entity.NewEntityManager(),
	}

	// start handler thread
	go sess.handlerLoop()

	return
}

func (sess *Session) Transmit(packet packets.Packet) {
	sess.TXQueue <- packet
}

func (sess *Session) handlerLoop() {
	for {
		// wait for packet
		packet := <-sess.RXQueue

		// get the handler
		handler := sess.handlers[packet.PacketID()]

		// try to run handler
		if handler != nil {
			handler(sess, packet)
		} else {
			log.Error("Handler for %v not found!", packet.PacketID())
		}
	}
}
