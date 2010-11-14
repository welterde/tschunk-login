package base

import "io"
import "os"
import "encoding/binary"

func ReadString(reader io.Reader) (s string, err os.Error) {
	// try to read string length(java type is short => int16)
	var length int16
	err = binary.Read(reader, binary.BigEndian, &length)
	if err != nil {
		return
	}

	// read that many bytes
	bs := make([]byte, uint16(length))
	_, err = io.ReadFull(reader, bs)
	return string(bs), err
}

func WriteString(writer io.Writer, s string) (err os.Error) {
	// convert string to an byte array
	bs := []byte(s)

	// try to write binary length
	err = binary.Write(writer, binary.BigEndian, int16(len(bs)))
	if err != nil {
		return
	}

	// now write the string
	_, err = writer.Write(bs)
	return
}
