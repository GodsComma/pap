package main

import (
	"fmt"
	"io"
	"os"
)

const MOV byte = 0b10001000
const OP_MASK byte = 0b11111100
const D_MASK byte =  0b00000010
const W_MASK byte =  0b00000001

const MOD_MASK byte = 0b11000000
const REG_MASK byte = 0b00111000
const RM_MASK  byte = 0b00000111

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
func main() {
	file, err := os.Open("./input/listing_0038")
	defer file.Close()

	if err != nil {
		panic("Error reading file: " + err.Error())
	}

	// Read in 2 bytes per read
	buffer := make([]byte, 2)

	for {
		_, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Done Reading File, EOF")
				break
			}
		}
		a, b := buffer[0], buffer[1]
    fmt.Printf("[%b], [%b]\n", a,b)
		op_code := ""
		if a & OP_MASK == MOV {
		 	op_code = "mov" 
		}
		// We take shortcuts here
		// We know MOD is always 11
		// We know D is always 0
		//fmt.Printf("[%b][%b]\n", a, b)
		is8bit := a & W_MASK == 0b0
		src := ""
		dest := ""
		if is8bit {
		   //fmt.Printf("8 Bit REG: %08b RM: %08b\n", b & REG_MASK, b & RM_MASK)
		   src  = B_REGS[b & RM_MASK]
		   dest = B_REGS[b & REG_MASK >> 3] 
		} else {
		   //fmt.Printf("16 Bit REG: %08b RM: %08b\n", b & REG_MASK >> 3, b & RM_MASK)
		   src = W_REGS[b & RM_MASK]
		   dest = W_REGS[b & REG_MASK >> 3]
		}
		fmt.Printf("%s, %s, %s\n", op_code, src, dest) 
	}
}
