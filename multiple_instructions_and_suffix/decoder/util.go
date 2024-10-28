package decoder

import (
	"bytes"
	"encoding/binary"
)

func Convert_16bits_to_int(a []byte) int16 {
	var ret int16
	buffer := bytes.NewReader(a)
	err := binary.Read(buffer, binary.NativeEndian, &ret)
	if err != nil {
		panic("Failed read 2bytes into a 16bit binary")
	}
	return ret
}
