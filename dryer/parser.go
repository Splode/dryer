package dryer

import (
	"fmt"
	"sort"
)

func Parse(s, p string, min int) error {
	srcFile, err := open(s)
	if err != nil {
		return err
	}
	srcTokens := Tokenize(srcFile.src, srcFile.absolutePath)

	patFile, err := open(p)
	if err != nil {
		return err
	}
	patTokens := Tokenize(patFile.src, patFile.absolutePath)

	res := Search(tokenSliceToStringer(srcTokens), tokenSliceToStringer(patTokens), min)

	keys := sortedKeys(res)
	for _, k := range keys {
		c := getTokenClones(res[k])
		Print(c)
		fmt.Println()
	}

	return nil
}

func getTokenClones(strs [][]Stringer) [][]Token {
	srcBeg := strs[0][0].(Token)
	srcEnd := strs[0][len(strs[0])-1].(Token)

	patBeg := strs[1][0].(Token)
	patEnd := strs[1][len(strs[0])-1].(Token)

	clones := make([][]Token, 0)
	clones = append(clones, []Token{srcBeg, srcEnd}, []Token{patBeg, patEnd})
	return clones
}

func sortedKeys(m map[int][][]Stringer) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Slice(keys, func(a, b int) bool {
		return keys[a] > keys[b]
	})
	return keys
}
