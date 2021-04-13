package dryer

// Config represents the configurable options for the dryer package.
type Config struct {
	Paths    []string // Paths represents the file paths to compare.
	Dir      string   // Dir is the directory to search for files, if using a pattern.
	Pattern  string   // Pattern is a glob-like pattern used to match files.
	TokenMin int      // TokenMin is the minimum number of tokens representing a match between 2 sources.
	Absolute bool     // Absolute is true when displaying the absolute file path in reporting.
	Recurse  bool     // Recurse is true when recursively searching for files matching a pattern.
	Verbose  bool     // Verbose is used for detailed reporting.
}
