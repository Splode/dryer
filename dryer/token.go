package dryer

import "fmt"

// token is a unit identified by the lexer from a source.
type token struct {
	tokenString string // tokenString representing the token contents
	position           // position of the token
}

// string returns the string representing the token.
func (t token) string() string {
	return t.tokenString
}

// position represents the position of a token in its source.
type position struct {
	filename string // filename, if any
	offset   int    // byte offset, starting at 0
	line     int    // line number, starting at 1
	column   int    // column number, starting at 1 (character count per line)
}

// lineColumn returns a string representing the line/column pair of a given token, e.g. "12:42".
func (p position) lineColumn() string {
	return fmt.Sprintf("%d:%d", p.line, p.column)
}

// tokenSliceToStringer returns a slice of stringer interface from an array of tokens.
func tokenSliceToStringer(tokens []token) []stringer {
	s := make([]stringer, len(tokens))
	for i, t := range tokens {
		s[i] = t
	}
	return s
}
