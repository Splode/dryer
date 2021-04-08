package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/splode/dryer/dryer"
)

var rootCmd = &cobra.Command{
	Use:   "dryer",
	Short: "Dryer helps",
	Long:  `something`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("root cmd")
		dryer.Parse()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
