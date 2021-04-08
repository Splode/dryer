package dryer

import (
	"log"
	"os"
	"path"
)

func Parse() {
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

	srcTokens := Tokenize(src, srcFileName)
	patTokens := Tokenize(pat, patFileName)

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
