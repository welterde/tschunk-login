package client


import "net"
import "bytes"

import packets "minecraft/packets"


type Client struct {
	conn    net.Conn
	txQueue chan packets.Packet
}

func StartClient(conn net.Conn) {
	// create client instance
	client := &Client{
		conn:    conn,
		txQueue: make(chan packets.Packet, 1024),
	}

	// start receive and transmit threads
	go client.receiveLoop()
	go client.transmitLoop()

	// TODO: add handler for handshake to the map
}

func (client *Client) receiveLoop() {
	for {
		_, err := packets.ReadPacket(client.conn)
		if err != nil {
			// TODO: do something useful here
			return
		}

		// TODO: call handler
	}
}

func (client *Client) transmitLoop() {
	for {
		// get next packet from the queue
		packet := <-client.txQueue

		// check if the queue still exists
		if packet == nil {
			return
		}

		// TODO: log packet id and some other stuff before sending

		// convert to bytes
		buf := &bytes.Buffer{}
		packet.Write(buf)

		// now try to send it
		_, err := client.conn.Write(buf.Bytes())
		if err != nil {
			// TODO: fail
			return
		}
	}
}
