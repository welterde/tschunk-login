package base

import "io"
import "os"
import "encoding/binary"

func ReadLong(reader io.Reader) (i int64, err os.Error) {
	// try to read the int
	err = binary.Read(reader, binary.BigEndian, &i)

	return
}

func WriteLong(writer io.Writer, i int64) (err os.Error) {
	return binary.Write(writer, binary.BigEndian, i)
}
