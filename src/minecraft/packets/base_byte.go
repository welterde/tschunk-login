package base

import "io"
import "os"
import "encoding/binary"

func ReadByte(reader io.Reader) (i byte, err os.Error) {
	// try to read the int
	err = binary.Read(reader, binary.BigEndian, &i)

	return
}

func WriteByte(writer io.Writer, i byte) (err os.Error) {
	return binary.Write(writer, binary.BigEndian, i)
}
