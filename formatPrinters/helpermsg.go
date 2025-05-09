package formatPrinters

import "fmt"

func HelpMessage() {
	fmt.Println("Command syntax: hexDumper  [OPTIONS] <SOURCE FILE PATH>")
	fmt.Println("Unavailability of any optional flags -> default output ")
	fmt.Println("AVAILABLE OPTIONS: ")
	fmt.Println(" -c, --c            prints the output in canonical format along with ascii representation")
	fmt.Println(" -d, --d            prints the output in decimal value")
	fmt.Println(" -help, --help      Show this help message")
	fmt.Println(" -n, --n <length>   How many bytes tto read from the source file")
	fmt.Println(" -s, --s <offset>   Show this help message")
}
