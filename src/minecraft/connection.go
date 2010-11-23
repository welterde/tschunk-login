package connection


import "net"
import "bytes"
import log "log4go"

import packets "minecraft/packets"
import session "minecraft/session"


type Connection struct {
	conn    net.Conn
	Session *session.Session
}

func HandleConnection(daemon session.Daemon, conn net.Conn) {
	// create session instance
	sess := session.StartSession(daemon)

	// create connection instance
	con := &Connection{
		conn:    conn,
		Session: sess,
	}

	// start receive and transmit threads
	go con.receiveLoop()
	go con.transmitLoop()
}

func (con *Connection) receiveLoop() {
	for {
		log.Finest("Waiting for packet..")
		packet, err := packets.ReadPacket(con.conn)
		if err != nil {
			log.Error("Failure in receive loop: %v", err)
			return
		}
		log.Fine("got packet %v", packet.PacketID())

		// put into receive queue
		con.Session.RXQueue <- packet
	}
}

func (con *Connection) transmitLoop() {
	for {
		// get next packet from the queue
		packet := <-con.Session.TXQueue

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
		_, err := con.conn.Write(buf.Bytes())
		if err != nil {
			// TODO: fail
			return
		}
	}
}
