package session


import "net"
import "bytes"

import packets "minecraft/packets"


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
		core:     core,
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
		packet, err := packets.ReadPacket(sess.conn)
		if err != nil {
			// TODO: do something useful here
			return
		}

		// get the handler
		handler := sess.handlers[packet.PacketID()]

		// try to run handler
		if handler != nil {
			handler(sess, packet)
		} else {
			// TODO: log this
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
