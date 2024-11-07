package decoder

import (
	"fmt"
	"pap/constants"
)

func Decode_Add_No_Displacement(a, b byte) string {
	is_8bit := a&constants.W_MASK == 0b0
	if is_8bit {
		return "8 Bit reat"
	}
	return "16 bit displacement"
}

func Decode_Add_Register(a, b byte) string {
	is8bit := a&constants.W_MASK == 0b0
	src := ""
	dest := ""
	if is8bit {
		src = B_REGS[b&constants.RM_MASK]
		dest = B_REGS[b&constants.REG_MASK>>3]
	} else {
		src = W_REGS[b&constants.RM_MASK]
		dest = W_REGS[b&constants.REG_MASK>>3]
	}

	return fmt.Sprintf("add, %s, %s\n", src, dest)
}
