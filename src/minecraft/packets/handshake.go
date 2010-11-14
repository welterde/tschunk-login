package handshake

import "io"
import "os"

import . "minecraft/packets/base"


type Request struct {
	username string
}

func (req *Request) PacketID() (id byte) {
	return 0x02
}

type Response struct {
	hash string
}

func (resp *Response) PacketID() (id byte) {
	return 0x02
}

func Read(reader io.Reader) (packet *Request, err os.Error) {
	// create structure to return
	packet = new(Request)

	// read string
	packet.username, err = ReadString(reader)

	return
}

func Write(writer io.Writer, packet *Response) (err os.Error) {
	// write hash
	err = WriteString(writer, packet.hash)

	return
}
