package dryer

// Config represents the configurable options for the dryer package.
type Config struct {
	TokenMin int      // TokenMin is the minimum number of tokens representing a match between 2 sources.
	Paths    []string // Paths represents the file paths to compare.
}
