package decoder

import (
	"fmt"
	"pap/constants"
	"strconv"
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

func Decode_IRD_0_8(a, b, c byte) string {
	// 8 bit displacement, 8 bit data
	reg := ND_REG[b&constants.RM_MASK]
	return fmt.Sprintf("mov %s, byte %d", reg, int(c))
}

func Decode_IRD_8_8(a, b, c, d byte) string {
	// 8 bit displacement, 8 bit data
	reg := DX_REG[b&constants.RM_MASK]
	return fmt.Sprintf("mov %s + %d], byte %d", reg, int(c), int(d))
}

func Decode_IRD_8_16(a, b, c, d, e byte) string {
	// 8 bit displacement, 16 bit data
	reg := DX_REG[b&constants.RM_MASK]
	displacement := int16(c)
	data := []byte{d, e}
	offset := strconv.Itoa(Convert_16bits_to_int(data))
	return fmt.Sprintf("mov, %s+ %d], %s", reg, displacement, offset)
}

func Decode_IRD_16_8(a, b, c, d, e byte) string {
	// 16 bit displacement, 8 bit data
	reg := DX_REG[b&constants.RM_MASK]
	d_data := []byte{c, d}
	displacement := strconv.Itoa(Convert_16bits_to_int(d_data))
	return fmt.Sprintf("mov, %s+ %s], word %d", reg, displacement, int(e))
}

func Decode_IRD_16_16(a, b, c, d, e, f byte) string {
	// 16 bit displacement, 16 bit data
	reg := DX_REG[b&constants.RM_MASK]
	d_data := []byte{c, d}
	displacement := strconv.Itoa(Convert_16bits_to_int(d_data))
	data := []byte{e, f}
	offset := strconv.Itoa(Convert_16bits_to_int(data))
	return fmt.Sprintf("mov, %s+ %s], word %s", reg, displacement, offset)
}
