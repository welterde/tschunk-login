package base

import "io"
import "os"
import "encoding/binary"

func ReadShort(reader io.Reader) (i int16, err os.Error) {
	// try to read the int
	err = binary.Read(reader, binary.BigEndian, &i)

	return
}

func WriteShort(writer io.Writer, i int16) (err os.Error) {
	return binary.Write(writer, binary.BigEndian, i)
}
