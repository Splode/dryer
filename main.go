package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	src, err := os.Open("./src.js")
	if err != nil {
		log.Fatal(err)
	}
	pat, err := os.Open("./pat.js")
	if err != nil {
		log.Fatal(err)
	}
	srcTokens := tokenize(src)
	patTokens := tokenize(pat)
	res := search(tokenSliceToStringer(srcTokens), tokenSliceToStringer(patTokens), 10)
	for i, r := range res {
		fmt.Printf("%d: %v\n", i, r[0])
		fmt.Printf("%d: %v\n", i, r[1])
	}
}

func tokenSliceToStringer(tokens []token) []Stringer {
	s := make([]Stringer, len(tokens))
	for i, t := range tokens {
		s[i] = t
	}
	return s
}
