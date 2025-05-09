package formatPrinters

import (
	"fmt"
	"slices"
)

func HexValuePrinter(lineNumber int, data []byte) {
	//Printing the hex values in every two bytes in reverse order
	if len(data)%2 != 0 {
		data = append(data, slices.Repeat([]byte{0}, 1)...)
	}
	fmt.Printf("%06x  ", lineNumber)
	for i := 0; i < len(data); i++ {
		if i > 0 && i%2 != 0 {
			fmt.Printf("%02x%02x ", data[i], data[i-1])
		}
	}
	fmt.Println()
}
func DecimalValuePrinter(lineNumber int, data []byte) {
	//Printing every two bit data in decimal format
	if len(data)%2 != 0 {
		data = append(data, slices.Repeat([]byte{0}, 1)...)
	}
	fmt.Printf("%06x  ", lineNumber)
	for i := 0; i < len(data); i += 2 {
		if i > 0 {
			fmt.Printf("%d%d  ", data[i-2], data[i-1])
		}
	}
	fmt.Println()
}
func CanonicalHexValuePrinter(lineNumber int, data []byte) {
	//Printing the contents in hex format with their ascii representation
	if len(data)%2 != 0 {
		data = append(data, slices.Repeat([]byte{0}, 1)...)
	}
	fmt.Printf("%06x  ", lineNumber)
	for i := 0; i < len(data); i++ {
		if i > 0 && i%2 != 0 {
			fmt.Printf("%02x%02x ", data[i-1], data[i])
		}
	}
}
func AsciiPrint(data []byte) {
	//Printing every 16 byte data in ascii format
	fmt.Print(" |")
	for _, v := range data {
		if 32 <= int(v) && int(v) <= 126 {
			fmt.Printf("%c", v)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println("| ")
}
