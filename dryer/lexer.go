package dryer

import (
	"io"
	"text/scanner"
	"unicode"
)

func tokenize(src io.Reader, name string) []token {
	var s scanner.Scanner
	s.Init(src)
	s.Filename = name
	s.Mode ^= scanner.ScanFloats // resolve errors with "px" values in CSS

	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '\'' || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0 || ch == '.' || ch == '='
	}

	tokens := make([]token, 0)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		str := s.TokenText()
		pos := position{filename: s.Filename, offset: s.Offset, line: s.Line, column: s.Column}
		tokens = append(tokens, token{tokenString: str, position: pos})
	}
	return tokens
}
