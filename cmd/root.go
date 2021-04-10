package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/splode/dryer/dryer"
)

func rootCmd() *cobra.Command {
	var tokenMin int

	var rootCmd = &cobra.Command{
		Use:   "dryer",
		Short: "Dryer identifies duplicate code between files, allowing you to stay dry.",
		Long: `
'||''|.
 ||   ||  ... ..  .... ...   ....  ... ..
 ||    ||  ||' ''  '|.  |  .|...||  ||' ''
 ||    ||  ||       '|.|   ||       ||
.||...|'  .||.       '|     '|...' .||.
                  .. |
                   ''

Dryer identifies duplicate code between files, allowing you to stay dry.
`,
		Run: func(cmd *cobra.Command, args []string) {
			// src, pat := args[0], args[1]
			paths := []string{"mural-section.bak", "story.bak", "mural.bak"}
			dryer.Compare(tokenMin, paths...)
		},
	}

	rootCmd.Flags().IntVarP(&tokenMin, "token", "t", 25, "The minimum number of tokens considered for a clone.")

	return rootCmd
}

func Execute() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
