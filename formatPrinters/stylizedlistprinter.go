package formatPrinters

import (
	lip "github.com/charmbracelet/lipgloss"
	table "github.com/charmbracelet/lipgloss/table"
)

func StylizedListPrinter(columns []string, data [][]string, rowColor []string) *table.Table {
	//var headerColor = []lip.Color{lip.Color("#fcba03"), lip.Color("8eb6d1"), lip.Color("1090e6")} //color definitions for the 3 columns
	t := table.New().
		Border(lip.RoundedBorder()).
		BorderStyle(lip.NewStyle().Foreground(lip.Color("c2cacf"))).
		StyleFunc(func(row, col int) lip.Style {
			switch {
			case row == table.HeaderRow:
				return lip.NewStyle().Background(lip.Color("#fcba03")).Align(0.0).Foreground(lip.Color("#000000"))
			default:
				return lip.NewStyle().Background(lip.Color(rowColor[row])).Foreground(lip.Color("#000000")).Align(0.0)
			}
		}).
		Headers(columns...).
		Rows(data...)

	return t
}
