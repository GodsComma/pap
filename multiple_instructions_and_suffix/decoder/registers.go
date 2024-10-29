package decoder

var B_REGS = map[byte]string{
	0b000: "al",
	0b001: "cl",
	0b010: "dl",
	0b011: "bl",
	0b100: "ah",
	0b101: "ch",
	0b110: "dh",
	0b111: "bh",
}

var W_REGS = map[byte]string{
	0b000: "ax",
	0b001: "cx",
	0b010: "dx",
	0b011: "bx",
	0b100: "sp",
	0b101: "bp",
	0b110: "si",
	0b111: "di",
}

var ND_REG = map[byte]string{
	0b000: "[bx + si]",
	0b001: "[bx + di]",
	0b010: "[bp + si]",
	0b011: "[bp + di]",
	0b100: "[si]",
	0b101: "[di]",
	0b111: "[bx]",
}

var DX_REG = map[byte]string{
	0b000: "[bx + si",
	0b001: "[bx + di",
	0b010: "[bp + si",
	0b011: "[bp + di",
	0b100: "[si  ",
	0b101: "[di  ",
	0b110: "[bp  ",
	0b111: "[bx  ",
}
