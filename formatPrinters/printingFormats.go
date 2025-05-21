package formatPrinters

import (
	"fmt"
	"slices"
)

func OctalValuePrinter(data []byte) string {
	//Printing the hex values in every two bytes in reverse order
	if len(data)%2 != 0 {
		data = append(data, slices.Repeat([]byte{0}, 1)...)
	}
	var result string
	for i := 0; i < len(data); i++ {
		if i > 0 && i%2 != 0 {
			result += fmt.Sprintf("%06o ", uint16(data[i])<<8|uint16(data[i-1]))
		}
	}
	return result
}
func HexValuePrinter(data []byte) string {
	//Printing the hex values in every two bytes in reverse order
	if len(data)%2 != 0 {
		data = append(data, slices.Repeat([]byte{0}, 1)...)
	}
	var result string
	for i := 0; i < len(data); i++ {
		if i > 0 && i%2 != 0 {
			result += fmt.Sprintf("%02x%02x ", data[i], data[i-1])
		}
	}
	return result
}
func DecimalValuePrinter(data []byte) string {
	//Printing every two bit data in decimal format
	if len(data)%2 != 0 {
		data = append(data, slices.Repeat([]byte{0}, 1)...)
	}
	var result string
	for i := 0; i < len(data); i += 2 {
		if i > 0 {
			result += fmt.Sprintf("%05d ", uint16(data[i-2])<<8|uint16(data[i-1]))
		}
	}
	return result
}
func CanonicalHexValuePrinter(data []byte) string {
	//Printing the contents in hex format with their ascii representation
	if len(data)%2 != 0 {
		data = append(data, slices.Repeat([]byte{0}, 1)...)
	}
	var result string
	for i := 0; i < len(data); i++ {
		if i > 0 && i%2 != 0 {
			result += fmt.Sprintf("%02x%02x ", data[i-1], data[i])
		}
	}
	return result
}
func AsciiPrint(data []byte) string {
	//Printing every 16 byte data in ascii format
	var result string
	for _, v := range data {
		if 32 <= int(v) && int(v) <= 126 {
			result += fmt.Sprintf("%c", v)
		} else {
			result += fmt.Sprint(".")
		}
	}
	return result
}
