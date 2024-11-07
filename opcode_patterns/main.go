package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"pap/constants"
	"pap/decoder"
)

// fmt.Printf("[%08b] [%08b]\n", a, b)

func main() {
	file, err := os.Open("./input/add")

	if err != nil {
		panic("Error reading file: " + err.Error())
	}

	defer file.Close()

	buffer_one := make([]byte, 1)
	buffer_two := make([]byte, 2)

	var debug bool
	debug, _ = strconv.ParseBool(os.Args[1])

	for {
		extra, err := file.Read(buffer_two)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Done Reading File, EOF", extra)
				break
			}
		}
		// read the first 2 bytes
		a, b := buffer_two[0], buffer_two[1]
		if debug {
			fmt.Printf("[%b], [%b]\n", a, b)
		}

		if a&constants.OP_MASK_I2R == constants.MOV_I2R {
			// Check for the W bit
			const w_mask byte = 0b00001000
			if a&w_mask == w_mask {
				file.Read(buffer_one)
				fmt.Printf("%s", decoder.Decode_IR_16(a, b, buffer_one[0]))
			} else {
				fmt.Printf("%s", decoder.Decode_IR_8(a, b))
			}
		} else if a&constants.OP_MASK_R2R == constants.MOV_R2R {
			if b&constants.MOD_MASK == constants.MM_ND0 {
				if b&constants.RM_MASK == constants.D_ADDR {
					file.Read(buffer_two)
					fmt.Printf("%s\n", decoder.Decode_DA(a, b, buffer_two[0], buffer_two[1]))
				} else {
					fmt.Printf("%s", decoder.Decoder_MM_ND("mov", a, b))
				}
			} else if b&constants.MOD_MASK == constants.MM_08B {
				file.Read(buffer_one)
				fmt.Printf("%s", decoder.Decode_MM_08B("mov", a, b, buffer_one[0]))
			} else if b&constants.MOD_MASK == constants.MM_16B {
				file.Read(buffer_two)
				fmt.Printf("%s", decoder.Decode_MM_16B("mov", a, b, buffer_two[0], buffer_two[1]))
			} else {
				fmt.Printf("%s", decoder.Decode_RM_R2R("mov", a, b))
			}
		} else if a&constants.OP_MASK_DI2R == constants.MOV_DI2R {
			mode_mask := b & constants.MOD_MASK
			if mode_mask == constants.MM_08B {
				is8bit := a&constants.W_MASK == 0b0
				if is8bit {
					file.Read(buffer_two)
					fmt.Printf("%s\n", decoder.Decode_IRD_8_8(a, b, buffer_two[0], buffer_two[1]))
				} else {
					file.Read(buffer_one)
					file.Read(buffer_two)
					fmt.Printf("%s\n", decoder.Decode_IRD_8_16(a, b, buffer_one[0], buffer_two[0], buffer_two[1]))
				}
			} else if mode_mask == constants.MM_16B {
				is8bit := a&constants.W_MASK == 0b0
				if is8bit {
					file.Read(buffer_two)
					file.Read(buffer_one)
					fmt.Printf("%s\n", decoder.Decode_IRD_16_8(a, b, buffer_one[0], buffer_two[0], buffer_one[1]))
				} else {
					file.Read(buffer_two)
					c, d := buffer_two[0], buffer_two[1]
					file.Read(buffer_two)
					fmt.Printf("%s\n", decoder.Decode_IRD_16_16(a, b, c, d, buffer_two[0], buffer_two[1]))
				}
			} else if mode_mask == constants.MM_ND0 {
				is8bit := a&constants.W_MASK == 0b0
				if is8bit {
					file.Read(buffer_one)
					fmt.Printf("%s\n", decoder.Decode_IRD_0_8(a, b, buffer_one[0]))
				} else {
					file.Read(buffer_one)
					file.Read(buffer_two)
					fmt.Printf("%s\n",
						decoder.Decode_IRD_8_16(a, b, buffer_one[0], buffer_two[0], buffer_two[1]))
				}
			}
		} else if a&constants.OP_MASK_M2A == constants.MOV_M2A {
			file.Read(buffer_one)
			fmt.Printf("%s", decoder.Decode_M2A_16(a, b, buffer_one[0]))
		} else if a&constants.OP_MASK_M2A == constants.MOV_A2M {
			file.Read(buffer_one)
			fmt.Printf("%s", decoder.Decode_A2M_16(a, b, buffer_one[0]))
		} else if a&constants.OP_MASK_ADD_2B == constants.ADD_R2RM {
			mode := b & constants.MOD_MASK
			if mode == constants.MM_RM0 {
				is8bit := a&constants.W_MASK == 0b0
				if is8bit {
					fmt.Println("8 bit")
					fmt.Printf("[%08b] [%08b] [%08b] [%08b]\n", a, b, buffer_two[0], buffer_two[1])
					//fmt.Printf("%s", decoder.Decode_RM_R2R("mov", a, b))
				} else {
					fmt.Println("16 bit")
					file.Read(buffer_two)
					fmt.Printf("[%08b] [%08b] [%08b] [%08b]\n", a, b, buffer_two[0], buffer_two[1])
					fmt.Printf("%s\n", decoder.Decode_RM_R2RM("add", a, b, buffer_one[0]))
				}
			} else if mode == constants.MM_ND0 {
				fmt.Printf("%s", decoder.Decoder_MM_ND("mov", a, b))
			} else if mode == constants.MM_08B {
				file.Read(buffer_one)
				fmt.Printf("%s", decoder.Decode_MM_08B("mov", a, b, buffer_one[0]))
			} else if mode == constants.MM_16B {
				fmt.Println("16 bit displacement")
				// 16 bit displacement
			} else {
				panic("Add Register to Memory")
			}
		} else if a&constants.OP_MASK_ADD_2B == constants.ADD_I2RM {
			mode := b & constants.MOD_MASK
			if mode == constants.MM_RM0 {
				is8bit := a&constants.W_MASK == 0b0
				isSBit := a&constants.S_MASK == 0b0
				if isSBit && !is8bit {
					file.Read(buffer_two)
					fmt.Printf("%s\n", decoder.Decode_MM_16B("add", a, b, buffer_two[0], buffer_two[1]))
				} else {
					file.Read(buffer_one)
					fmt.Printf("%s\n", decoder.Decode_MM_08B("add", a, b, buffer_one[0]))
				}
			}
		} else {
			// fmt.Printf("[%08b] [%08b]\n", a, b)
			fmt.Printf("-> -> [%08b] [%08b]", a, b)
		}
	}
}
