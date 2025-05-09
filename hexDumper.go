package main

import (
	"errors"
	"fmt"
	parser "hexDumper/argumentsParser"
	pf "hexDumper/formatPrinters"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	options, length, startOffset, path := parser.CommandLineParser()
	if slices.Contains(options, "help") { // When a help message is requested
		pf.HelpMessage()
		os.Exit(0)
	}
	if strings.Compare("", path) == 0 { // When there is no path to the source file is passed
		fmt.Println("Please enter a valid path to the source file")
		pf.HelpMessage()
		os.Exit(0)
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Unable to locate the file from disk. %s\n", err)
		os.Exit(0)
	}
	if fileInfo.IsDir() {
		fmt.Printf("The source path given is not a file but a directory. %s\n", err)
		os.Exit(0)
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("There was some error opening the file from disk. %s\n", err)
		os.Exit(0)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error while closing the file.", err)
		}
	}(file)
	var data = make([]byte, 16)
	for lenOffSet := startOffset; lenOffSet < length; {
		n, err := file.ReadAt(data, int64(lenOffSet))
		if n+lenOffSet > length { // Check if the total bits read from the file will go over the total length
			n = length - lenOffSet
		}
		if slices.Contains(options, "c") {
			pf.CanonicalHexValuePrinter(lenOffSet, data[:n])
			pf.AsciiPrint(data[:n])
		}
		if slices.Contains(options, "d") {
			pf.DecimalValuePrinter(lenOffSet, data[:n])
		}
		if len(options) == 0 {
			pf.HexValuePrinter(lenOffSet, data[:n])
		}
		if err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Println("\nError reading the data from the source file\n", err)
			}
			break
		}
		lenOffSet += n
	}

}
