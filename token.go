package main

type token struct {
	tokenString string
	Position
}

type Position struct {
	Filename string // filename, if any
	Offset   int    // byte offset, starting at 0
	Line     int    // line number, starting at 1
	Column   int    // column number, starting at 1 (character count per line)
}

func (t token) String() string {
	return t.tokenString
}
