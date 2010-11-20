package session


import "net"
import "bytes"
import log "log4go"

import packets "minecraft/packets"
import entity  "minecraft/entity"


type Handler func(sess *Session, packet packets.Packet)

type Session struct {
	conn     net.Conn
	txQueue  chan packets.Packet
	handlers map[byte]Handler
}

func StartSession(conn net.Conn) {
	// create session instance
	sess := &Session{
		conn:     conn,
		txQueue:  make(chan packets.Packet, 1024),
		handlers: make(map[byte]Handler),
	}

	// start receive and transmit threads
	go sess.receiveLoop()
	go sess.transmitLoop()

	// TODO: register new client with core
}

func (sess *Session) Transmit(packet packets.Packet) {
	sess.txQueue <- packet
}

func (sess *Session) receiveLoop() {
	for {
		log.Finest("Waiting for packet..")
		packet, err := packets.ReadPacket(sess.conn)
		if err != nil {
			log.Error("Failure in receive loop: %v", err)
			return
		}
		log.Fine("got packet %v", packet.PacketID())

		// get the handler
		handler := sess.handlers[packet.PacketID()]

		// try to run handler
		if handler != nil {
			handler(sess, packet)
		} else {
			log.Info("Missing handler for packet-id: %v", packet.PacketID())
		}
	}
}

func (sess *Session) transmitLoop() {
	for {
		// get next packet from the queue
		packet := <-sess.txQueue

		// check if the queue still exists
		if packet == nil {
			return
		}

		// TODO: log packet id and some other stuff before sending
		log.Fine("sending packet %v", packet.PacketID())

		// convert to bytes
		buf := &bytes.Buffer{}
		packet.Write(buf)

		// now try to send it
		_, err := sess.conn.Write(buf.Bytes())
		if err != nil {
			// TODO: fail
			return
		}
	}
}
