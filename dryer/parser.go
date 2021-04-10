package dryer

import (
	"fmt"
	"sort"
	"sync"

	"github.com/splode/dryer/pkg/strings"
)

func Compare(tokenMin int, paths ...string) {
	pathMatrix := strings.UniqueMatrix(paths...)

	var wg sync.WaitGroup
	wg.Add(len(pathMatrix))
	res := make([][][]string, 0)
	for _, matrix := range pathMatrix {
		m := matrix
		go func(wg *sync.WaitGroup) {
			clones, err := parse(m[0], m[1], tokenMin)
			if err != nil {
				fmt.Println(err)
			}
			res = append(res, clones)
			wg.Done()
		}(&wg)
	}
	wg.Wait()

	for _, r := range res {
		if len(r) > 1 {
			Print(r)
		}
	}
}

func parse(s, p string, min int) ([][]string, error) {
	srcFile, err := open(s)
	if err != nil {
		return nil, err
	}
	srcTokens := Tokenize(srcFile.src, srcFile.absolutePath)

	patFile, err := open(p)
	if err != nil {
		return nil, err
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

	return cloneData, nil
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
