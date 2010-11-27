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

var MC_TABLE [256]func() (packet Packet)

func init() {
	// set default handler, which returns nil
	defH := func() Packet { return nil }
	for i := 0; i < 256; i++ {
		MC_TABLE[i] = defH
	}
	MC_TABLE[handshake.REQ_PID] = func() Packet { return new(handshake.Request) }
	MC_TABLE[login.REQ_PID] = func() Packet { return new(login.Request) }
}

func ReadPacket(reader io.Reader) (packet Packet, err os.Error) {
	// read packet id
	packetID, err := primitive.ReadByte(reader)
	if err != nil {
		return
	}

	// get the correct packet
	packet = MC_TABLE[packetID]()

	if packet == nil {
		return nil, os.NewError("Handler not found")
	}

	// now read the message
	err = packet.Read(reader)
	return
}
