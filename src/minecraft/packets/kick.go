package kick

import "io"
import "os"

import . "minecraft/packets/base"


const PID = 0xFF


type Packet struct {
	Message string
}

func (msg *Packet) PacketID() (id byte) {
	return PID
}

func (msg *Packet) Read(reader io.Reader) (err os.Error) {
	return
}

func (msg *Packet) Write(writer io.Writer) (err os.Error) {
	// write reason(if there is any ;) for the kick
	err = WriteString(writer, msg.Message)
	
	return
}
