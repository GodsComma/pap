package main

import (
	"fmt"
	"io"
	"os"
)

const MOV byte = 0b10001000
const OP_MASK byte = 0b11111100

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
		a, _ := buffer[0], buffer[1]
		fmt.Printf("[%t][%b][%b]\n", (a&OP_MASK == MOV), (MOV & a), a)
	}
}
