package dryer

import (
	"log"
	"os"
	"path/filepath"
)

func Parse() {
	srcFilePath := "./src.js"
	src, err := os.Open(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}

	patFilePath := "./pat.js"
	pat, err := os.Open(patFilePath)
	if err != nil {
		log.Fatal(err)
	}

	srcPath, err := filepath.Abs(srcFilePath)
	if err != nil {
		log.Fatal(err)
	}
	srcTokens := Tokenize(src, srcPath)

	patPath, err := filepath.Abs(patFilePath)
	if err != nil {
		log.Fatal(err)
	}
	patTokens := Tokenize(pat, patPath)

	res := Search(tokenSliceToStringer(srcTokens), tokenSliceToStringer(patTokens), 10)

	for _, r := range res {

		srcBeg := r[0][0].(Token)
		srcEnd := r[0][len(r[0])-1].(Token)

		patBeg := r[1][0].(Token)
		patEnd := r[1][len(r[0])-1].(Token)

		clones := make([][]Token, 0)
		clones = append(clones, []Token{srcBeg, srcEnd}, []Token{patBeg, patEnd})
		Print(clones)
	}
}
