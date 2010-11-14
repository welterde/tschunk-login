package base

import "io"
import "os"
import "encoding/binary"

func ReadInt(reader io.Reader) (i int32, err os.Error) {
	// try to read the int
	err = binary.Read(reader, binary.BigEndian, &i)

	return
}

func WriteInt(writer io.Writer, i int32) (err os.Error) {
	return binary.Write(writer, binary.BigEndian, i)
}
