package decoder

import "fmt"

func Decode_M2A_16(a, b, c byte) string {
	data := []byte{b, c}
	return fmt.Sprintf("mov ax, [%d]\n", Convert_16bits_to_int(data))
}

func Decode_A2M_16(a, b, c byte) string {
	data := []byte{b, c}
	return fmt.Sprintf("mov [%d], ax\n", Convert_16bits_to_int(data))
}
