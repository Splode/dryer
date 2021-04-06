package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	srcFilePath := "./src.js"
	_, srcFileName := path.Split(srcFilePath)
	src, err := os.Open("./src.js")
	if err != nil {
		log.Fatal(err)
	}

	patFilePath := "./pat.js"
	_, patFileName := path.Split(patFilePath)
	pat, err := os.Open("./pat.js")
	if err != nil {
		log.Fatal(err)
	}

	srcTokens := tokenize(src, srcFileName)
	patTokens := tokenize(pat, patFileName)

	res := search(tokenSliceToStringer(srcTokens), tokenSliceToStringer(patTokens), 10)

	for i, r := range res {
		fmt.Printf("%d: %v\n", i, r[0])
		fmt.Printf("%d: %v\n", i, r[1][0].String())

		srcBeg := r[0][0].(token)
		srcEnd := r[0][len(r[0])-1].(token)
		fmt.Printf("file: %s, line: %v, column: %v\t", srcBeg.Filename, srcBeg.Line, srcBeg.Column)
		fmt.Printf("file: %s, line: %v, column: %v\n", srcEnd.Filename, srcEnd.Line, srcEnd.Column)

		patBeg := r[1][0].(token)
		patEnd := r[1][len(r[0])-1].(token)
		fmt.Printf("file: %s, line: %v, column: %v\t", patBeg.Filename, patBeg.Line, patBeg.Column)
		fmt.Printf("file: %s, line: %v, column: %v\n", patEnd.Filename, patEnd.Line, patEnd.Column)
	}
}

func tokenSliceToStringer(tokens []token) []Stringer {
	s := make([]Stringer, len(tokens))
	for i, t := range tokens {
		s[i] = t
	}
	return s
}
