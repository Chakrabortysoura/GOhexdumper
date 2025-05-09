package argumentsParser

import (
	"fmt"
	pf "hexDumper/formatPrinters"
	"math"
	"os"
	"strconv"
	"strings"
)

func convertToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("There was an error optional arguments format")
		pf.HelpMessage()
		os.Exit(1)
	}
	return result
}
func CommandLineParser() ([]string, int, int, string) {
	var path string //Path for the source file
	var options []string
	length := math.MaxInt
	startOffset := 0
	for i := 1; i < len(os.Args); i++ {
		if strings.Contains(os.Args[i], "-") && len(os.Args[i][1:]) >= 1 { //To parse optional flags passed to the programme
			temp := os.Args[i][strings.LastIndex(os.Args[i], "-")+1:]
			if strings.Compare("n", temp) == 0 {
				length = convertToInt(os.Args[i+1])
				i++
			} else if strings.Compare("s", temp) == 0 {
				startOffset = convertToInt(os.Args[i+1])
				i++
			} else {
				options = append(options, os.Args[i][strings.LastIndex(os.Args[i], "-")+1:])
			}
		} else {
			path = os.Args[i]
		}
	}
	return options, length, startOffset, path
}
