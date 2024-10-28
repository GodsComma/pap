package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"pap/decoder"
)

func main() {
	file, err := os.Open("./input/listing_0039")

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
					fmt.Println("THIS IS THE DIRECT ADDRESS ON MOD 00")
					//TODO: Check for Direct Addressing
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
		} else {
			// fmt.Println("Not Immediate to Register")
		}
	}
}
