package packets

import "io"
import "os"

type Packet interface {
	PacketID() (id byte)
	Read(reader io.Reader) (err os.Error)
	Write(writer io.Writer) (err os.Error)
}
