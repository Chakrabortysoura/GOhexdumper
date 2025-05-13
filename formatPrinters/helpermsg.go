package formatPrinters

import (
	"fmt"
	lip "github.com/charmbracelet/lipgloss"
)

func HelpMessage() {
	message := fmt.Sprintf(
		"Command syntax: hexDumper  [OPTIONS] <SOURCE FILE PATH>\nUnavailability of any optional flags -> default output \nAVAILABLE OPTIONS: \n-c, --c            prints the output in canonical format along with ascii representation\n-d, --d            prints the output in decimal value\n-help, --help      Show this help message\n-n, --n <length>   How many bytes tto read from the source file\n-s, --s <offset>   Show this help message\n")
	var style = lip.NewStyle().Bold(true).Foreground(lip.Color("#FAFAFA")).Background(lip.Color("#4bbadb")).Border(lip.RoundedBorder(), true, true, true, true).Width(90)
	fmt.Println(style.Render(message))
}
