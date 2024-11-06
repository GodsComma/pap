package decoder

import (
	"fmt"
	"strconv"
)

func Decode_RM_R2R(a, b byte) string {
	is8bit := a&W_MASK == 0b0
	src := ""
	dest := ""
	if is8bit {
		src = B_REGS[b&RM_MASK]
		dest = B_REGS[b&REG_MASK>>3]
	} else {
		src = W_REGS[b&RM_MASK]
		dest = W_REGS[b&REG_MASK>>3]
	}

	return fmt.Sprintf("mov, %s, %s\n", src, dest)
}

func Decoder_MM_ND(a, b byte) string {
	is8bit := a&W_MASK == 0b0
	isDest := a&D_MASK == 0b0

	src := ""
	dest := ""

	if is8bit {
		src = B_REGS[b&REG_MASK>>3]
	} else {
		src = W_REGS[b&REG_MASK>>3]
	}

	dest = ND_REG[b&RM_MASK]
	if isDest {
		return fmt.Sprintf("mov %s, %s\n", dest, src)
	}
	return fmt.Sprintf("mov, %s, %s\n", src, dest)
}

func Decode_MM_08B(a, b, c byte) string {
	isDest := a&D_MASK == 0b0
	is8bit := a&W_MASK == 0b0
	src := ""
	if is8bit {
		src = B_REGS[b&REG_MASK>>3]
	} else {
		src = W_REGS[b&REG_MASK>>3]
	}
	dest := DX_REG[b&RM_MASK]
	offset := strconv.Itoa(int(int8(c)))

	if isDest {
		return fmt.Sprintf("mov %s %s], %s\n", dest, offset, src)
	}
	return fmt.Sprintf("mov, %s, %s %s]\n", src, dest, offset)

}

func Decode_MM_16B(a, b, c, d byte) string {
	is8bit := a&W_MASK == 0b0
	isDest := a&D_MASK == 0b0

	src := ""
	if is8bit {
		src = B_REGS[b&REG_MASK>>3]
	} else {
		src = W_REGS[b&REG_MASK>>3]
	}
	dest := DX_REG[b&RM_MASK]
	data := []byte{c, d}
	offset := strconv.Itoa(Convert_16bits_to_int(data))

	if isDest {
		return fmt.Sprintf("mov %s %s], %s\n", dest, offset, src)
	}
	return fmt.Sprintf("mov, %s, %s %s]\n", src, dest, offset)
}
