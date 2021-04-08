package dryer

type Token struct {
	tokenString string
	Position
}

type Position struct {
	Filename string // filename, if any
	Offset   int    // byte offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (character count per line)
}

func (t Token) String() string {
	return t.tokenString
}

func tokenSliceToStringer(tokens []Token) []Stringer {
	s := make([]Stringer, len(tokens))
	for i, t := range tokens {
		s[i] = t
	}
	return s
}
