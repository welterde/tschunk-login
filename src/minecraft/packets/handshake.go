package handshake

import "io"
import "os"

import . "minecraft/packets/base"


const REQ_PID = 0x02
const RES_PID = 0x02


type Request struct {
	username string
}

func (req *Request) PacketID() (id byte) {
	return 0x02
}

func (packet *Request) Read(reader io.Reader) (err os.Error) {
	// read string
	packet.username, err = ReadString(reader)

	return
}

func (packet *Request) Write(writer io.Writer) (err os.Error) {
	return
}


type Response struct {
	hash string
}

func (packet *Response) PacketID() (id byte) {
	return 0x02
}

func (packet *Response) Read(reader io.Reader) (err os.Error) {
	return
}

func (packet *Response) Write(writer io.Writer) (err os.Error) {
	// write hash
	err = WriteString(writer, packet.hash)

	return
}
