package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/splode/dryer/dryer"
)

func rootCmd() *cobra.Command {
	var dir string
	var pattern string
	var tokenMin int
	var absolute bool
	var recurse bool
	var verbose bool

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
			paths := args
			cfg := &dryer.Config{
				Paths: paths, Dir: dir, Pattern: pattern, TokenMin: tokenMin,
				Absolute: absolute, Recurse: recurse, Verbose: verbose,
			}
			if err := dryer.Compare(cfg); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 && pattern == "" {
				return errors.New("requires a minimum of 2 sources when a pattern isn't specified")
			}
			return nil
		},
	}

	rootCmd.SetUsageTemplate(getUsage())

	rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "Directory to look for files.")
	rootCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "A glob-like pattern to match files.")
	rootCmd.Flags().IntVarP(&tokenMin, "token", "t", 25, "The minimum number of tokens considered for a clone.")
	rootCmd.Flags().BoolVarP(&absolute, "absolute", "a", false, "Display a file's absolute path.")
	rootCmd.Flags().BoolVarP(&recurse, "recurse", "r", false, "Recursively match files when given a pattern.")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Display verbose output.")

	return rootCmd
}

func Execute() {
	if err := rootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getUsage() string {
	return `
Usage:
  dryer [flags] [source]...

Examples:
  dryer <fileOne> <fileTwo>
  dryer --token 30 <fileOne> <fileTwo> <fileThree>
  dryer --pattern *.js --dir ./src

Flags:
  -a, --absolute        Display a file's absolute path.
  -r, --recurse          Recursively match files when given a pattern.
  -d, --dir string       Directory to look for files. (default ".")
  -h, --help             Display help for dryer.
  -p, --pattern string   A glob-like pattern to match files.
  -t, --token int        The minimum number of tokens considered for a clone. (default 25)
  -v, --verbose          Display verbose output.
`
}
