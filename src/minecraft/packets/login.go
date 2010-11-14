package login


import "io"
import "os"


import . "minecraft/packets/base"


type Request struct {
	ProtocolVersion int32
	Username        string
	Password        string
	MapSeed         int64
	Dimension       byte
}

func (packet *Request) PacketID() (id byte) {
	return 0x01
}

func (packet *Request) Read(reader io.Reader) (err os.Error) {
	// read protocol version - should be 4 for now..
	packet.ProtocolVersion, err = ReadInt(reader)
	if err != nil {
		return
	}

	// username
	packet.Username, err = ReadString(reader)
	if err != nil {
		return
	}

	// password
	packet.Password, err = ReadString(reader)
	if err != nil {
		return
	}

	// ignored for now
	packet.MapSeed, err = ReadLong(reader)
	if err != nil {
		return
	}

	// ignored for now
	packet.Dimension, err = ReadByte(reader)
	return
}


type Response struct {
	EID       int32
	MapSeed   int64
	Dimension byte
}
