package dryer

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func Print(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetBorder(false)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Filepath", "Start", "End"})
	table.AppendBulk(data)
	table.Render()
}
