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
	cloneData := make([][]string, 0)
	for _, k := range keys {
		c := getTokenClones(res[k])
		d := cloneTableData(c)
		cloneData = append(cloneData, d...)
	}

	if len(cloneData) > 1 {
		Print(cloneData)
	}

	return nil
}

func cloneTableData(clones [][]Token) [][]string {
	td := make([][]string, len(clones))
	srcB := clones[0][0]
	srcE := clones[0][1]
	patB := clones[1][0]
	patE := clones[1][1]
	td = append(td,
		[]string{srcB.Filename, fmt.Sprintf("%d:%d", srcB.Line, srcB.Column), fmt.Sprintf("%d:%d", srcE.Line, srcE.Column)},
		[]string{patB.Filename, fmt.Sprintf("%d:%d", patB.Line, patB.Column), fmt.Sprintf("%d:%d", patE.Line, patE.Column)},
		[]string{"\t"},
	)
	return td
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
