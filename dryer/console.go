package dryer

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func print(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Filepath", "Start", "End", "Percentage"})
	table.AppendBulk(data)
	table.Render()
}
