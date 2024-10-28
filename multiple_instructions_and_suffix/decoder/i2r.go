package decoder

import (
	"fmt"
)

func Decode_IR_8(a, b byte) string {
	reg := B_REGS[a&0b0000111]
	return fmt.Sprintf("mov, %s, %d\n", reg, int8(b))
}
func Decode_IR_16(a, b, c byte) string {
	reg := W_REGS[a&0b0000111]
	data := []byte{b, c}
	return fmt.Sprintf("mov, %s, %d\n", reg, Convert_16bits_to_int(data))
}
