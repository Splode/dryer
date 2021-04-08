package dryer

import (
	"io"
	"text/scanner"
	"unicode"
)

func Tokenize(src io.Reader, name string) []Token {
	var s scanner.Scanner
	s.Init(src)
	s.Filename = name

	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '\'' || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0 || ch == '.' || ch == '='
	}

	tokens := make([]Token, 0)
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		str := s.TokenText()
		pos := Position{Filename: s.Filename, Offset: s.Offset, Line: s.Line, Column: s.Column}
		tokens = append(tokens, Token{tokenString: str, Position: pos})
	}
	return tokens
}
