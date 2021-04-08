package dryer

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func Print(clones [][]Token) {
	tableData := make([][]string, len(clones))
	srcBeg := clones[0][0]
	srcEnd := clones[0][1]
	patBeg := clones[1][0]
	patEnd := clones[1][1]
	tableData = append(tableData,
		[]string{srcBeg.Filename, fmt.Sprintf("%d:%d", srcBeg.Line, srcBeg.Column), fmt.Sprintf("%d:%d", srcEnd.Line, srcEnd.Column)},
		[]string{patBeg.Filename, fmt.Sprintf("%d:%d", patBeg.Line, patBeg.Column), fmt.Sprintf("%d:%d", patEnd.Line, patEnd.Column)},
	)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Filepath", "Start", "End"})
	table.SetBorder(false)
	table.SetAutoFormatHeaders(false)
	table.AppendBulk(tableData)
	table.Render() // Send output
}