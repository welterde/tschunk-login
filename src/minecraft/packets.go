package packets


import "io"
import "os"

import primitive "minecraft/packets/base"
import handshake "minecraft/packets/handshake"
import login     "minecraft/packets/login"


type Packet interface {
	PacketID() (id byte)
	Read(reader io.Reader) os.Error
	Write(writer io.Writer) (err os.Error)
}


func ReadPacket(reader io.Reader) (packet Packet, err os.Error) {
	// read packet id
	packetID, err := primitive.ReadByte(reader)
	if err != nil {
		return
	}

	// get the correct packet
	switch packetID {
	case handshake.REQ_PID:
		packet = new(handshake.Request)
	case login.REQ_PID:
		packet = new(login.Request)
	}

	// now read the message
	err = packet.Read(reader)
	return
}
