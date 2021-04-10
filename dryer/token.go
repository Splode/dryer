package dryer

type token struct {
	tokenString string
	position
}

type position struct {
	Filename string // filename, if any
	Offset   int    // byte offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (character count per line)
}

func (t token) string() string {
	return t.tokenString
}

func tokenSliceToStringer(tokens []token) []stringer {
	s := make([]stringer, len(tokens))
	for i, t := range tokens {
		s[i] = t
	}
	return s
}
