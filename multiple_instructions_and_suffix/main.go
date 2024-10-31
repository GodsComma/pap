package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"pap/decoder"
)

func main() {
	file, err := os.Open("./input/listing_0040")

	if err != nil {
		panic("Error reading file: " + err.Error())
	}

	defer file.Close()

	buffer_one := make([]byte, 1)
	buffer_two := make([]byte, 2)

	var debug bool
	debug, _ = strconv.ParseBool(os.Args[1])

	for {
		_, err := file.Read(buffer_two)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Done Reading File, EOF")
				break
			}
		}
		a, b := buffer_two[0], buffer_two[1]
		if debug {
			fmt.Printf("[%b], [%b]\n", a, b)
		}

		if a&decoder.OP_MASK_I2R == decoder.MOV_I2R {
			// Check for the W bit
			const w_mask byte = 0b00001000
			if a&w_mask == w_mask {
				file.Read(buffer_one)
				fmt.Printf("%s", decoder.Decode_IR_16(a, b, buffer_one[0]))
			} else {
				fmt.Printf("%s", decoder.Decode_IR_8(a, b))
			}
		} else if a&decoder.OP_MASK_R2R == decoder.MOV_R2R {
			if b&decoder.MOD_MASK == decoder.MM_ND0 {
				if b&decoder.RM_MASK == decoder.D_ADDR {
					file.Read(buffer_two)
					fmt.Printf("%s\n", decoder.Decode_DA(a, b, buffer_two[0], buffer_two[1]))
				} else {
					fmt.Printf("%s", decoder.Decoder_MM_ND(a, b))
				}
			} else if b&decoder.MOD_MASK == decoder.MM_08B {
				file.Read(buffer_one)
				fmt.Printf("%s", decoder.Decode_MM_08B(a, b, buffer_one[0]))

			} else if b&decoder.MOD_MASK == decoder.MM_16B {
				file.Read(buffer_two)
				fmt.Printf("%s", decoder.Decode_MM_16B(a, b, buffer_two[0], buffer_two[1]))
			} else {
				fmt.Printf("%s", decoder.Decode_RM_R2R(a, b))
			}
		} else if a&decoder.OP_MASK_DI2R == decoder.MOV_DI2R {
			mode_mask := b & decoder.MOD_MASK
			if mode_mask == decoder.MM_08B {
				is8bit := a&decoder.W_MASK == 0b0
				if is8bit {
					file.Read(buffer_two)
					fmt.Printf("%s\n", decoder.Decode_IRD_8_8(a, b, buffer_two[0], buffer_two[1]))
				} else {
					file.Read(buffer_one)
					file.Read(buffer_two)
					fmt.Printf("%s\n", decoder.Decode_IRD_8_16(a, b, buffer_one[0], buffer_two[0], buffer_two[1]))
				}
			} else if mode_mask == decoder.MM_16B {
				is8bit := a&decoder.W_MASK == 0b0
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

			} else if mode_mask == decoder.MM_ND0 {
				is8bit := a&decoder.W_MASK == 0b0
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
		} else if a&decoder.OP_MASK_M2A == decoder.MOV_M2A {
			file.Read(buffer_one)
			fmt.Printf("%s", decoder.Decode_M2A_16(a, b, buffer_one[0]))
		} else if a&decoder.OP_MASK_M2A == decoder.MOV_A2M {
			file.Read(buffer_one)
			fmt.Printf("%s", decoder.Decode_A2M_16(a, b, buffer_one[0]))
		} else {
			fmt.Printf("-> -> [%08b] [%08b]", a, b)
			// fmt.Println("Not Immediate to Register")
		}
	}
}
