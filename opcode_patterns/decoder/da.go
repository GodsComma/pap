package decoder

import (
	"fmt"
	"pap/constants"
)

func Decode_DA(a, b, c, d byte) string {
	data := []byte{c, d}
	dest := W_REGS[b&constants.REG_MASK>>3]
	return fmt.Sprintf("mov %s, [%d]", dest, Convert_16bits_to_int(data))
}
