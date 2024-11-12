package decoder

import (
	"fmt"
	"pap/constants"
	"strconv"
)

func Decode_RM_R2R(op string, a, b byte) string {
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

	return fmt.Sprintf("%s, %s, %s\n", op, src, dest)
}

func Decode_RM_R2RM(op string, a, b, c byte) string {
	src := ""
	data := []byte{b, c}
	offset := strconv.Itoa(Convert_16bits_to_int(data))
	src = W_REGS[b&constants.RM_MASK]
	return fmt.Sprintf("%s, %s, %s\n", op, src, offset)
}

func Decoder_MM_ND(op string, a, b byte) string {
	is8bit := a&constants.W_MASK == 0b0
	isDest := a&constants.D_MASK == 0b0

	src := ""
	dest := ""

	if is8bit {
		src = B_REGS[b&constants.REG_MASK>>3]
	} else {
		src = W_REGS[b&constants.REG_MASK>>3]
	}

	dest = ND_REG[b&constants.RM_MASK]
	if isDest {
		return fmt.Sprintf("%s %s, %s\n", op, dest, src)
	}
	return fmt.Sprintf("%s, %s, %s\n", op, src, dest)
}

func Decode_MM_08B(op string, a, b, c byte) string {
	isDest := a&constants.D_MASK == 0b0
	is8bit := a&constants.W_MASK == 0b0
	src := ""
	if is8bit {
		src = B_REGS[b&constants.REG_MASK>>3]
	} else {
		src = W_REGS[b&constants.REG_MASK>>3]
	}
	dest := DX_REG[b&constants.RM_MASK]
	offset := strconv.Itoa(int(int8(c)))

	if isDest {
		return fmt.Sprintf("%s %s+ %s], %s\n", op, dest, offset, src)
	}
	return fmt.Sprintf("%s, %s, %s+ %s]\n", op, src, dest, offset)
}

func Decode_MM_16B(op string, a, b, c, d byte) string {
	is8bit := a&constants.W_MASK == 0b0
	isDest := a&constants.D_MASK == 0b0

	src := ""
	if is8bit {
		src = B_REGS[b&constants.REG_MASK>>3]
	} else {
		src = W_REGS[b&constants.REG_MASK>>3]
	}
	dest := DX_REG[b&constants.RM_MASK]
	data := []byte{c, d}
	offset := strconv.Itoa(Convert_16bits_to_int(data))

	if isDest {
		return fmt.Sprintf("%s %s %s], %s\n", op, dest, offset, src)
	}
	return fmt.Sprintf("%s, %s, %s %s]\n", op, src, dest, offset)
}

func Decode_RR_D8(op string, a, b, c byte) string {
	dest := W_REGS[b&constants.RM_MASK]
	offset := strconv.Itoa(int(int8(c)))
	return fmt.Sprintf("%s, %s, %s\n", op, dest, offset)
}

func Decode_D08_D8(op string, a, b, c byte) string {
	dest := ND_REG[b&constants.RM_MASK]
	offset := strconv.Itoa(int(int8(c)))
	return fmt.Sprintf("%s, %s, %s\n", op, dest, offset)
}
