package main

import (
	"fmt"

	"github.com/pterm/pterm"
)

func print(clones [][]token) {
	tableData := make([][]string, len(clones)+1)
	tableData = append(tableData, []string{"Filepath", "Starting Line", "Starting Column", "Ending Line", "Ending Column"})
	srcBeg := clones[0][0]
	srcEnd := clones[0][1]
	patBeg := clones[1][0]
	patEnd := clones[1][1]
	tableData = append(tableData,
		[]string{srcBeg.Filename, fmt.Sprint(srcBeg.Line), fmt.Sprint(srcBeg.Column), fmt.Sprint(srcEnd.Line), fmt.Sprint(srcEnd.Column)},
		[]string{patBeg.Filename, fmt.Sprint(patBeg.Line), fmt.Sprint(patBeg.Column), fmt.Sprint(patEnd.Line), fmt.Sprint(patEnd.Column)},
	)
	pterm.DefaultTable.WithHasHeader().WithData(tableData).Render()
}
