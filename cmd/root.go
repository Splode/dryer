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
		src, pat := args[0], args[1]
		dryer.Parse(src, pat)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
