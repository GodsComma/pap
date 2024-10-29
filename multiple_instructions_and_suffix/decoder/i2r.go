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

func Decode_IRD_8(a, b, c byte) string {
	memory_mode := b & MOD_MASK
	if memory_mode == MM_ND0 {
		reg := ND_REG[b&RM_MASK]
		return fmt.Sprintf("mov %s, %d\n", reg, int(c))
	}
	return ""
}

func Decode_IRD_16(a, b, c, d byte) string {
	fmt.Printf("[%08b] [%08b] [%08b] [%08b]\n", a, b, c, d)
	fmt.Println("16 bit decode")
	return ""
}
